package mtproto

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"time"
)

func (session *MSession) sendPacket(msg TL, resp chan response) error {
	obj := msg.encode()

	x := NewEncodeBuf(256)

	// padding for tcpsize
	x.Int(0)

	if session.encrypted {
		needAck := true
		switch msg.(type) {
		case TL_ping, TL_msgs_ack:
			needAck = false
		}
		z := NewEncodeBuf(256)
		newMsgId := GenerateMessageId()
		z.Bytes(session.serverSalt)
		z.Long(session.sessionId)
		z.Long(newMsgId)
		if needAck {
			z.Int(session.lastSeqNo | 1)
		} else {
			z.Int(session.lastSeqNo)
		}
		z.Int(int32(len(obj)))
		z.Bytes(obj)

		msgKey := sha1(z.buf)[4:20]
		aesKey, aesIV := generateAES(msgKey, session.authKey, false)

		y := make([]byte, len(z.buf)+((16-(len(obj)%16))&15))
		copy(y, z.buf)
		encryptedData, err := doAES256IGEencrypt(y, aesKey, aesIV)
		if err != nil {
			return err
		}

		session.lastSeqNo += 2
		if needAck {
			session.mutex.Lock()
			session.msgsIdToAck[newMsgId] = packetToSend{msg, resp}
			session.mutex.Unlock()
		}

		x.Bytes(session.authKeyHash)
		x.Bytes(msgKey)
		x.Bytes(encryptedData)

		if resp != nil {
			session.mutex.Lock()
			session.msgsIdToResp[newMsgId] = resp
			session.mutex.Unlock()
		}

	} else {
		x.Long(0)
		x.Long(GenerateMessageId())
		x.Int(int32(len(obj)))
		x.Bytes(obj)

	}

	// minus padding
	size := len(x.buf)/4 - 1

	if size < 127 {
		x.buf[3] = byte(size)
		x.buf = x.buf[3:]
	} else {
		binary.LittleEndian.PutUint32(x.buf, uint32(size<<8|127))
	}
	_, err := session.tcpconn.Write(x.buf)
	if err != nil {
		return err
	}

	return nil
}

func (session *MSession) read() (interface{}, error) {
	var err error
	var n int
	var size int
	var data interface{}
	tcpconn := session.tcpconn

	err = tcpconn.SetReadDeadline(time.Now().Add(300 * time.Second))
	if err != nil {
		return nil, err
	}

	// Read packet size
	b := make([]byte, 1)
	n, err = tcpconn.Read(b)	// Wait for an incoming byte
	if err != nil {
		return nil, err
	}

	if b[0] < 127 {
		size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = tcpconn.Read(b)
		if err != nil {
			return nil, err
		}
		size = (int(b[0]) | int(b[1])<<8 | int(b[2])<<16) << 2
	}

	// Read packet
	left := size
	buf := make([]byte, size)
	for left > 0 {
		n, err = tcpconn.Read(buf[size-left:])
		if err != nil {
			return nil, err
		}
		left -= n
	}

	if size == 4 {
		return nil, fmt.Errorf("Server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
	}

	// Deserialize incoming packet
	data, session.msgId, session.seqNo, err = deserialize(buf, session.authKey)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (session *MSession) makeAuthKey() error {
	var x []byte
	var err error
	var data interface{}

	// (send) req_pq
	nonceFirst := GenerateNonce(16)
	err = session.sendPacket(TL_req_pq{nonceFirst}, nil)
	if err != nil {
		return err
	}

	// (parse) resPQ
	data, err = session.read()
	if err != nil {
		return err
	}
	res, ok := data.(TL_resPQ)
	if !ok {
		return errors.New("Handshake: Need resPQ")
	}
	if !bytes.Equal(nonceFirst, res.Nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	found := false
	for _, b := range res.Fingerprints {
		if uint64(b) == telegramPublicKey_FP {
			found = true
			break
		}
	}
	if !found {
		return errors.New("Handshake: No fingerprint")
	}

	// (encoding) p_q_inner_data
	p, q := splitPQ(res.Pq)
	nonceSecond := GenerateNonce(32)
	nonceServer := res.Server_nonce
	innerData1 := (TL_p_q_inner_data{res.Pq, p, q, nonceFirst, nonceServer, nonceSecond}).encode()

	x = make([]byte, 255)
	copy(x[0:], sha1(innerData1))
	copy(x[20:], innerData1)
	encryptedData1 := doRSAencrypt(x)
	// (send) req_DH_params
	err = session.sendPacket(TL_req_DH_params{nonceFirst, nonceServer, p, q, telegramPublicKey_FP, encryptedData1}, nil)
	if err != nil {
		return err
	}

	// (parse) server_DH_params_{ok, fail}
	data, err = session.read()
	if err != nil {
		return err
	}
	dh, ok := data.(TL_server_DH_params_ok)
	if !ok {
		return errors.New("Handshake: Need server_DH_params_ok")
	}
	if !bytes.Equal(nonceFirst, dh.Nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dh.Server_nonce) {
		return errors.New("Handshake: Wrong Server_nonce")
	}
	t1 := make([]byte, 48)
	copy(t1[0:], nonceSecond)
	copy(t1[32:], nonceServer)
	hash1 := sha1(t1)

	t2 := make([]byte, 48)
	copy(t2[0:], nonceServer)
	copy(t2[16:], nonceSecond)
	hash2 := sha1(t2)

	t3 := make([]byte, 64)
	copy(t3[0:], nonceSecond)
	copy(t3[32:], nonceSecond)
	hash3 := sha1(t3)

	tmpAESKey := make([]byte, 32)
	tmpAESIV := make([]byte, 32)

	copy(tmpAESKey[0:], hash1)
	copy(tmpAESKey[20:], hash2[0:12])

	copy(tmpAESIV[0:], hash2[12:20])
	copy(tmpAESIV[8:], hash3)
	copy(tmpAESIV[28:], nonceSecond[0:4])

	// (parse-thru) server_DH_inner_data
	decodedData, err := doAES256IGEdecrypt(dh.Encrypted_answer, tmpAESKey, tmpAESIV)
	if err != nil {
		return err
	}
	innerbuf := NewDecodeBuf(decodedData[20:])
	data = innerbuf.Object()
	if innerbuf.err != nil {
		return innerbuf.err
	}
	dhi, ok := data.(TL_server_DH_inner_data)
	if !ok {
		return errors.New("Handshake: Need server_DH_inner_data")
	}
	if !bytes.Equal(nonceFirst, dhi.Nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhi.Server_nonce) {
		return errors.New("Handshake: Wrong Server_nonce")
	}

	_, g_b, g_ab := makeGAB(dhi.G, dhi.G_a, dhi.Dh_prime)
	session.authKey = g_ab.Bytes()
	if session.authKey[0] == 0 {
		session.authKey = session.authKey[1:]
	}
	session.authKeyHash = sha1(session.authKey)[12:20]
	t4 := make([]byte, 32+1+8)
	copy(t4[0:], nonceSecond)
	t4[32] = 1
	copy(t4[33:], sha1(session.authKey)[0:8])
	nonceHash1 := sha1(t4)[4:20]
	session.serverSalt = make([]byte, 8)
	copy(session.serverSalt, nonceSecond[:8])
	xor(session.serverSalt, nonceServer[:8])

	// (encoding) client_DH_inner_data
	innerData2 := (TL_client_DH_inner_data{nonceFirst, nonceServer, 0, g_b}).encode()
	x = make([]byte, 20+len(innerData2)+(16-((20+len(innerData2))%16))&15)
	copy(x[0:], sha1(innerData2))
	copy(x[20:], innerData2)
	encryptedData2, err := doAES256IGEencrypt(x, tmpAESKey, tmpAESIV)

	// (send) set_client_DH_params
	err = session.sendPacket(TL_set_client_DH_params{nonceFirst, nonceServer, encryptedData2}, nil)
	if err != nil {
		return err
	}

	// (parse) dh_gen_{ok, Retry, fail}
	data, err = session.read()
	if err != nil {
		return err
	}
	dhg, ok := data.(TL_dh_gen_ok)
	if !ok {
		return errors.New("Handshake: Need dh_gen_ok")
	}
	if !bytes.Equal(nonceFirst, dhg.Nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhg.Server_nonce) {
		return errors.New("Handshake: Wrong Server_nonce")
	}
	if !bytes.Equal(nonceHash1, dhg.New_nonce_hash1) {
		return errors.New("Handshake: Wrong New_nonce_hash1")
	}

	// (all ok)
	err = session.saveSession()
	if err != nil {
		return err
	}

	return nil
}
