package mtp

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

func serialize(encrypt bool) {}

func Deserialize(buf []byte, authKey []byte) (data interface{}, msgId int64, seqNo int32, err error) {
	dbuf := NewDecodeBuf(buf)

	authKeyHash := dbuf.Bytes(8)
	if binary.LittleEndian.Uint64(authKeyHash) == 0 {
		msgId = dbuf.Long()
		messageLen := dbuf.Int()
		if int(messageLen) != dbuf.size-20 {
			//TODO: check if 0 works for seqNo
			return nil, msgId, 0, fmt.Errorf("Message len: %d (need %d)", messageLen, dbuf.size-20)
		}
		//m.seqNo = 0
		seqNo = 0

		data = dbuf.Object()
		if dbuf.err != nil {
			return nil, msgId, seqNo, dbuf.err
		}

	} else {
		msgKey := dbuf.Bytes(16)
		encryptedData := dbuf.Bytes(dbuf.size - 24)
		//aesKey, aesIV := generateAES(msgKey, m.authKey, true)
		aesKey, aesIV := generateAES(msgKey, authKey, true)
		x, err := doAES256IGEdecrypt(encryptedData, aesKey, aesIV)
		if err != nil {
			//TODO: check if 0 works for msgId and seqNo
			return nil, 0, 0, err
		}
		dbuf = NewDecodeBuf(x)
		_ = dbuf.Long() // salt
		_ = dbuf.Long() // session_id
		//m.msgId = dbuf.Long()
		//m.seqNo = dbuf.Int()
		msgId = dbuf.Long()
		seqNo = dbuf.Int()
		messageLen := dbuf.Int()
		if int(messageLen) > dbuf.size-32 {
			return nil, msgId, seqNo, fmt.Errorf("Message len: %d (need less than %d)", messageLen, dbuf.size-32)
		}
		if !bytes.Equal(sha1(dbuf.buf[0 : 32+messageLen])[4:20], msgKey) {
			return nil, msgId, seqNo, errors.New("Wrong msg_key")
		}

		data = dbuf.Object()
		if dbuf.err != nil {
			return nil, msgId, seqNo, dbuf.err
		}

	}
	mod := msgId & 3
	if mod != 1 && mod != 3 {
		return nil, msgId, seqNo, fmt.Errorf("Wrong bits of message_id: %d", mod)
	}

	return data, msgId, seqNo, nil
}
