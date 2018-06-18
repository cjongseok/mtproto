package mtproto

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/cjongseok/slog"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	ENV_AUTHKEY     = "MTPROTO_AUTHKEY"
	ENV_AUTHHASH    = "MTPROTO_AUTHHASH"
	ENV_SERVER_SALT = "MTPROTO_SALT"
	ENV_ADDR        = "MTPROTO_ADDR"
	ENV_USE_IPV6    = "MTPROTO_USE_IPV6"

	// API Errors
	errorSeeOther     = 303
	errorBadRequest   = 400
	errorUnauthorized = 401
	errorForbidden    = 403
	errorNotFound     = 404
	errorFlood        = 420
	errorInternal     = 500
)

type handshakingFailure struct {
	msg string
}

func (h handshakingFailure) Error() string {
	return h.msg
}

type Session struct {
	connId    int32
	sessionId int64
	//phonenumber string
	//addr        string
	//useIPv6     bool
	listeners []chan Event
	tcpconn   *net.TCPConn
	f         *os.File
	queueSend chan packetToSend

	readInterrupter chan struct{}
	sendInterrupter chan struct{}
	pingInterrupter chan struct{}

	isReading bool
	isSending bool
	isPing    bool

	readWaitGroup sync.WaitGroup
	sendWaitGroup sync.WaitGroup
	pingWaitGroup sync.WaitGroup

	c *Credentials
	//authKey     []byte
	//authKeyHash []byte
	//serverSalt  []byte
	//encrypted   bool

	mutex        *sync.Mutex
	lastSeqNo    int32
	msgsIdToAck  map[int64]packetToSend
	msgsIdToResp map[int64]chan response
	seqNo        int32
	msgId        int64

	appConfig Configuration
	//user         *TL_user
	//updatesState *TL_updates_state
	user         *PredUser
	updatesState *PredUpdatesState

	// dcOptions contains Telegram server list. Its 1st key is a IP version, such as 'ipv4' and 'ipv6',
	// and 2nd key is server's id.
	dcOptions map[string]map[int32]PredDcOption
}

// IP versions for dcOptions 1st key
const (
	ipv4 = "ipv4"
	ipv6 = "ipv6"
)

type packetToSend struct {
	msg  TL
	resp chan response
}

type response struct {
	//data TL
	data interface{}
	err  error
}

// CHECK: Only mmanager call this method?
func (session *Session) close() {
	session.stopPing()
	session.stopSend()
	session.pingWaitGroup.Wait()
	session.sendWaitGroup.Wait()

	session.tcpconn.Close()

	session.stopRead()
	session.readWaitGroup.Wait()

	// notify that the connection is gracefully closed
	if session.updatesState == nil {
		session.notify(SessionDiscarded{session.connId, session.sessionId, &PredUpdatesState{}})
	} else {
		session.notify(SessionDiscarded{session.connId, session.sessionId, session.updatesState})
	}
	session.listeners = nil
}

//func sessionFilePath(sessionFileHome string, phonenumber string) string {
//	return sessionFileHome + "/.telegram_" + phonenumber
//}

func isIPv6(addr string) bool {
	tcpAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return false
	}
	if tcpAddr.IP.To4() != nil {
		return false
	} else if tcpAddr.IP.To16() != nil {
		return true
	}
	return false
}

func newSession(phone string, apiID int32, apiHash, ip string, port int, appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event) (*Session, error) {
	var err error

	session := new(Session)
	// create empty key file
	session.f, err = os.OpenFile(appConfig.KeyPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	// open session
	credentials := Credentials{
		Phone:   phone,
		ApiID:   apiID,
		ApiHash: apiHash,
		IP:      ip,
		Port:    port,
	}
	//if err == nil {
	//	session.addr = addr
	//	session.useIPv6 = useIPv6
	//	session.encrypted = false
	useIPv6 := isIPv6(credentials.IP)
	err = session.open(useIPv6, credentials, appConfig /*sendQueue,*/, sessionListener, false)
	if err != nil {
		return nil, err
	}
	return session, nil
	//}
	//return nil, err
}

// byte array string is bracketed space separated numbers in a string
func byteArrayString2byteArray(str string) []byte {
	runes := []rune(str)
	if len(runes) < 3 {
		return nil
	}
	byteStrings := strings.Split(string(runes[1:len(runes)-1]), " ")
	byteArr := make([]byte, len(byteStrings))
	for i, s := range byteStrings {
		s2int, _ := strconv.Atoi(s)
		byteArr[i] = byte(s2int)
	}
	return byteArr
}

// loadSession loads session from the credentials file of appConfig
// it does not guarantee its session id matches with the previous one.
func loadSession(appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event) (*Session, error) {
	// load session info from either session file
	session := new(Session)
	var err error
	if appConfig.KeyPath == "" {
		return nil, fmt.Errorf("no credentials file")
	}
	session.f, err = os.OpenFile(appConfig.KeyPath, os.O_RDONLY, 0600)
	if err != nil {
		return nil, fmt.Errorf("no credentials file; %v", err)
	}
	credentials, err := NewCredentialsFromFile(session.f)
	if err != nil {
		return nil, fmt.Errorf("new credentials from file failure; %v", err)
	}
	useIPv6 := isIPv6(credentials.IP)
	err = session.open(useIPv6, *credentials, appConfig /*sendQueue,*/, sessionListener, true)
	if err != nil {
		return session, handshakingFailure{fmt.Sprintf("Handshaking Failure: %v", err)}
	}
	return session, nil
}

func (session *Session) open(useIPv6 bool, c Credentials, appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event, getUpdateStates bool) error {
	var err error
	var tcpAddr *net.TCPAddr

	// set up rest of session
	session.c = &c
	session.appConfig = appConfig
	rand.Seed(time.Now().UnixNano())
	session.sessionId = rand.Int63()
	session.AddSessionListener(sessionListener)

	// connect
	serverAddr := fmt.Sprintf("%s:%d", c.IP, c.Port)
	tcpAddr, err = net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		return err
	}
	slog.Logf(session, "dial TCP to %s\n", serverAddr)
	session.tcpconn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}

	// packet length is encoded by a single byte (see: https://core.telegram.org/mtproto)
	_, err = session.tcpconn.Write([]byte{0xef})
	if err != nil {
		return err
	}
	// get new authKey if need
	if c.AuthKey == nil {
		authKey, authKeyHash, salt, err := session.makeAuthKey()
		if err != nil {
			return err
		}
		c.AuthKey = authKey
		c.AuthKeyHash = authKeyHash
		c.Salt = salt
	}

	// start goroutines
	session.queueSend = make(chan packetToSend, 64)
	//session.queueSend = sendQueue
	session.sendInterrupter = make(chan struct{})
	session.readInterrupter = make(chan struct{})
	session.pingInterrupter = make(chan struct{})

	session.sendWaitGroup = sync.WaitGroup{}
	session.readWaitGroup = sync.WaitGroup{}
	session.pingWaitGroup = sync.WaitGroup{}

	session.msgsIdToAck = make(map[int64]packetToSend)
	session.msgsIdToResp = make(map[int64]chan response)
	session.mutex = &sync.Mutex{}
	session.sendWaitGroup.Add(1)
	session.readWaitGroup.Add(1)
	session.isSending = true
	session.isReading = true
	go session.sendRoutine(session.appConfig.SendInterval)
	go session.readRoutine()

	// (help_getConfig)
	var x response
	resp := make(chan response, 1)
	session.queueSend <- packetToSend{
		msg: &ReqInvokeWithLayer{
			Layer: int32(layer),
			Query: Pack(&ReqInitConnection{
				ApiId:          session.c.ApiID,
				DeviceModel:    session.appConfig.DeviceModel,
				SystemVersion:  session.appConfig.SystemVersion,
				AppVersion:     session.appConfig.Version,
				SystemLangCode: session.appConfig.Language,
				LangCode:       session.appConfig.Language,
				Query:          Pack(&ReqHelpGetConfig{}),
			}),
		},
		resp: resp,
	}
	select {
	case x = <-resp:
		if x.err != nil {
			return fmt.Errorf("TL_invokeWithLayer Failure;", x.err)
		}
	case <-time.After(TIMEOUT_INVOKE_WITH_LAYER):
		return fmt.Errorf("TL_invokeWithLayer Timeout(%f s)", TIMEOUT_INVOKE_WITH_LAYER.Seconds())
	}

	switch x.data.(type) {
	case *PredConfig:
		session.dcOptions = make(map[string]map[int32]PredDcOption)
		session.dcOptions[ipv4] = make(map[int32]PredDcOption)
		session.dcOptions[ipv6] = make(map[int32]PredDcOption)
		for _, v := range x.data.(*PredConfig).DcOptions {
			isIPv6 := true
			dcOption := v.GetValue()
			tcpAddr, err := net.ResolveIPAddr("ip", dcOption.GetIpAddress())
			if err == nil {
				ip := tcpAddr.IP.To4()
				if ip != nil {
					isIPv6 = false
				}
				//if useIPv6 {
				if isIPv6 {
					session.dcOptions[ipv6][dcOption.Id] = *dcOption
				} else {
					session.dcOptions[ipv4][dcOption.Id] = *dcOption
				}
			}
		}
		marshaled, err := json.Marshal(x.data)
		if err == nil {
			slog.Logf(session, "help config: %s\n", marshaled)
		}
	default:
		return fmt.Errorf("no help config; got %T, instead", x)
	}

	// get updates state
	session.updatesState = new(PredUpdatesState)
	if getUpdateStates {
		//TODO: From second session, query getUpdatesState with invokeWithLayer and initConnection
		resp = make(chan response, 1)
		session.queueSend <- packetToSend{
			msg:  &ReqUpdatesGetState{},
			resp: resp,
		}
		select {
		case x = <-resp:
			if x.err != nil {
				return fmt.Errorf("updatesGetState failure; %v", x.err)
			}
		case <-time.After(TIMEOUT_UPDATES_GETSTATE):
			//session.close()
			return fmt.Errorf("updatesGetState timeout(%f s)", TIMEOUT_UPDATES_GETSTATE.Seconds())
		}
	}

	// start keep alive ping
	session.isPing = true
	session.pingWaitGroup.Add(1)
	go session.pingRoutine()

	// notify the connection established
	session.notify(SessionEstablished{session})

	return nil
}

func (session *Session) AddSessionListener(listener chan Event) {
	session.listeners = append(session.listeners, listener)
}

func (session *Session) RemoveSessionListener(toremove chan Event) error {
	// find the listener index in the list
	for index, registered := range session.listeners {
		if registered == toremove {
			// remove the element by the index
			copy(session.listeners[index:], session.listeners[index+1:])
			session.listeners[len(session.listeners)-1] = nil
			session.listeners = session.listeners[:len(session.listeners)-1]
			return nil
		}
	}
	return fmt.Errorf("no listener, %x", toremove)
}

func (session *Session) notify(e Event) {
	slog.Logf(session, "notify Event, %s, to %v\n", slog.Stringify(e), session.listeners)
	for _, listener := range session.listeners {
		// TODO: it doesn't work. think of another solution to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (session *Session) process(msgId int64, seqNo int32, data interface{}) interface{} {
	returned := func() interface{} {
		switch data.(type) {
		case TL_msg_container:
			data := data.(TL_msg_container).Items
			for _, v := range data {
				session.process(v.Msg_id, v.Seq_no, v.Data)
			}

		case TL_bad_server_salt:
			data := data.(TL_bad_server_salt)
			session.c.Salt = data.new_server_salt
			// save the session on sign in
			_ = session.c.Save(session.f)
			session.mutex.Lock()
			defer session.mutex.Unlock()
			for k, v := range session.msgsIdToAck {
				delete(session.msgsIdToAck, k)
				session.queueSend <- v
			}

		case TL_new_session_created:
			data := data.(TL_new_session_created)
			session.c.Salt = data.server_salt
			// save the session on sign in
			_ = session.c.Save(session.f)

		case TL_ping:
			data := data.(TL_ping)
			session.queueSend <- packetToSend{TL_pong{msgId, data.ping_id}, nil}

		case TL_pong:
			// ignore

		case TL_msgs_ack:
			data := data.(TL_msgs_ack)
			session.mutex.Lock()
			defer session.mutex.Unlock()
			for _, v := range data.msgIds {
				delete(session.msgsIdToAck, v)
			}

		case TL_rpc_result:
			//slog.Logf(session, "rpc_result before casting: %v\n", data)
			data := data.(TL_rpc_result)
			//slog.Logln(session, "stringify(rpc_result.Obj):", slog.Stringify(data.Obj))
			//slog.Logf(session, "rpc_result.Obj: %T: %v\n", data.Obj, data.Obj)
			//if rpcerr, ok := data.Obj.(TL_rpc_error); ok {
			//slog.Logln(session, "ok stringify(rpcerror):", rpcerr)
			//slog.Logf(session, "ok rpcerror: %v\n", rpcerr)
			//slog.Logf(session, "ok rpcerror.code: %d, rpcerror.msg: %s\n", rpcerr.error_code, rpcerr.error_message)
			//}
			x := session.process(msgId, seqNo, data.Obj)
			session.mutex.Lock()
			defer session.mutex.Unlock()
			v, ok := session.msgsIdToResp[data.req_msg_id]
			if ok {
				go func() {
					var resp response
					rpcError, ok := x.(TL_rpc_error)
					if ok {
						//resp.err = session.handleRPCError(rpcError)
						resp.err = rpcError
					} else {
						resp.data = x
					}
					//resp.data = x.(TL)
					//resp.data = x
					v <- resp
					close(v)
				}()
			}
			delete(session.msgsIdToAck, data.req_msg_id)

		case TL_rpc_error:
			data := data.(TL_rpc_error)
			return data

			//TODO: Update classifier should be auto-generated from scheme.tl
			//TODO: how to handle seq?
			// Date, Pts, Qts updates
		case *PredUpdatesState:
			data := data.(*PredUpdatesState)
			session.updatesState.Pts = data.Pts
			session.updatesState.Qts = data.Qts
			session.updatesState.Date = data.Date
			session.updatesState.Seq = data.Seq
			marshaled, err := json.Marshal(data)
			if err == nil {
				slog.Logf(session, "updatesState: %s\n", marshaled)
			} else {
				slog.Logf(session, "updatesState: %v\n", data)
			}
			return data

			// Date updates
		case *PredUpdates:
			data := data.(*PredUpdates)
			session.updatesState.Date = data.Date
			session.updatesState.Seq = data.Seq
			session.notify(updateReceived{data})
			return data
		case *PredUpdateShort:
			data := data.(*PredUpdateShort)
			//session.updatesState.Pts ++	//TODO: need to comment in it?
			session.updatesState.Date = data.Date
			session.notify(updateReceived{data})
			return data

			// Pts updates
		case *PredUpdateNewMessage:
			data := data.(*PredUpdateNewMessage)
			session.updatesState.Pts = data.Pts
			session.notify(updateReceived{data})
			return data
		case *PredUpdateReadMessagesContents:
			data := data.(*PredUpdateReadMessagesContents)
			session.updatesState.Pts = data.Pts
			session.notify(updateReceived{data})
			return data
		case *PredUpdateDeleteMessages:
			data := data.(*PredUpdateDeleteMessages)
			session.updatesState.Pts = data.Pts
			session.notify(updateReceived{data})
			return data

			// Pts and Date updates
		case *PredUpdateShortMessage:
			data := data.(*PredUpdateShortMessage)
			session.updatesState.Pts = data.Pts
			session.updatesState.Date = data.Date
			session.notify(updateReceived{data})
			return data
		case *PredUpdateShortChatMessage:
			data := data.(*PredUpdateShortChatMessage)
			session.updatesState.Pts = data.Pts
			session.updatesState.Date = data.Date
			session.notify(updateReceived{data})
			return data
		case *PredUpdateShortSentMessage:
			data := data.(*PredUpdateShortSentMessage)
			session.updatesState.Pts = data.Pts
			session.updatesState.Date = data.Date
			session.notify(updateReceived{data})
			return data

			// Qts updates
		case *PredUpdateNewEncryptedMessage:
			data := data.(*PredUpdateNewEncryptedMessage)
			session.updatesState.Qts = data.Qts
			session.notify(updateReceived{data})
			return data

			// Channel updates
		case *PredUpdateChannel:
			data := data.(*PredUpdateChannel)
			session.notify(updateReceived{data})
			return data
		case *PredUpdateChannelMessageViews:
			data := data.(*PredUpdateChannelMessageViews)
			session.notify(updateReceived{data})
			return data
		case *PredUpdateChannelTooLong:
			data := data.(*PredUpdateChannelTooLong)
			session.updatesState.Pts = data.Pts
			session.notify(updateReceived{data})
			return data
		case *PredUpdateReadChannelInbox:
			data := data.(*PredUpdateReadChannelInbox)
			session.notify(updateReceived{data})
			return data
		case *PredUpdateReadChannelOutbox:
			data := data.(*PredUpdateReadChannelOutbox)
			session.notify(updateReceived{data})
			return data
		case *PredUpdateNewChannelMessage:
			data := data.(*PredUpdateNewChannelMessage)
			session.updatesState.Pts = data.Pts
			session.notify(updateReceived{data})
			return data

		default:
			marshaled, err := json.Marshal(data)
			if err == nil {
				slog.Logf(session, "process: unknown data type %T {%s}\n", data, marshaled)
			} else {
				slog.Logf(session, "process: unknown data type %T {%v}\n", data, data)
			}
			return data
		}
		return nil
	}()

	if returned != nil {
		return returned
	}

	// TODO: Check why I should do this
	if (seqNo & 1) == 1 {
		session.queueSend <- packetToSend{TL_msgs_ack{[]int64{msgId}}, nil}
	}

	return nil
}

func (session *Session) stopRead() {
	if session.isReading {
		close(session.readInterrupter)
	}
}

func (session *Session) stopSend() {
	if session.isSending {
		close(session.sendInterrupter)
		close(session.queueSend)
	}
}

func (session *Session) stopPing() {
	if session.isPing {
		close(session.pingInterrupter)
	}
}

func (session *Session) pingRoutine() {
	slog.Logln(session, "ping: start")
	defer func() {
		session.isPing = false
		session.pingWaitGroup.Done()
	}()
	for {
		select {
		case <-session.pingInterrupter:
			session.isPing = false
			return
		case <-time.After(session.appConfig.PingInterval):
			session.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
		}
	}
}

func (session *Session) sendRoutine(interval time.Duration) {
	slog.Logln(session, "send: start")
	defer func() {
		session.isSending = false
		session.sendWaitGroup.Done()
	}()
	wg := sync.WaitGroup{}
	t := time.NewTimer(interval)
	timerInterrupter := make(chan struct{})
	go func() {
		for {
			select {
			case <-timerInterrupter:
				return
			case <-t.C:
				wg.Done()
			}
		}
	}()
	for {
		select {
		case <-session.sendInterrupter:
			slog.Logln(session, "send: stop")
			session.isSending = false
			close(timerInterrupter)
			return
		case x := <-session.queueSend:
			if _, ok := x.msg.(TL_ping); !ok {
				slog.Logf(session, "send %s\n", slog.Stringify(x.msg))
			}
			if x.msg != nil {
				//TODO: alternate interval based scheduler with frequency scheduler
				wg.Wait()
				err := session.send(x.msg, x.resp)
				wg.Add(1)
				t.Reset(interval)
				if err != nil {
					slog.Logln(session, "send err:", err)
				}
			}
		}
	}
}

func (session *Session) readRoutine() {
	slog.Logln(session, "read: start")
	defer func() {
		session.isReading = false
		session.readWaitGroup.Done()
	}()

	innerRoutineWG := sync.WaitGroup{}

	for {
		// Run async wait for data from server
		ch := make(chan interface{}, 1)
		go func(ch chan<- interface{}) {
			refreshUntilSuccess := func(session *Session) {
				//respChan := make(chan sessionResponse)
				//for {
				session.notify(refreshSession{
					session.sessionId,
					session.c.Phone,
					untilSuccess,
					nil,
				})
				//resp := <-respChan
				//if resp.err == nil {
				//	break
				//}
				//}
			}
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := session.read()
			if _, ok := data.(TL_pong); !ok {
				slog.Logf(session, "read: type: %v, data: %v, err: %v\n", reflect.TypeOf(data), data, err)
			}
			//slog.Logf(session, "read: %s\n", slog.Stringify(data))
			serverAddr := fmt.Sprintf("%s:%d", session.c.IP, session.c.Port)
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				slog.Logf(session, "read: lost connection (captured EOF). reconnect to %s\n", serverAddr)
				refreshUntilSuccess(session)
			} else if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					slog.Logf(session, "read: TCP connection closed (%s)\n", err)
					// Two cases
					// 1. on new authentication, 303 PHONE_MIGRATE can require to make a new connection with different
					//   server by closing the connection. -> do nothing, because session will be renewed by MM
					// 2. after authentication, there could be an accidental disconnection. -> need to refreshUntilSuccess
					if session.user == nil {
						// case 1
						// do nothing
					} else {
						// case 2
						refreshUntilSuccess(session)
					}
				} else if strings.Contains(err.Error(), "connection reset by peer") {
					slog.Logf(session, "read: lost connection (%s). reconnect to %s\n", err, serverAddr)
					refreshUntilSuccess(session)
				} else if strings.Contains(err.Error(), "i/o timeout") {
					slog.Logf(session, "read: lost connection (%s). reconnect to %s\n", err, serverAddr)
					refreshUntilSuccess(session)
				} else {
					slog.Logf(session, "read: unknown error, %s. reconnect to %s\n", err, serverAddr)
					refreshUntilSuccess(session)
				}
			} else {
				ch <- data
			}
		}(ch)

		select {
		case <-session.readInterrupter:
			slog.Logln(session, "read: wait for inner routine ...")
			session.isReading = false
			innerRoutineWG.Wait()
			slog.Logln(session, "read: stop")
			return
		case data := <-ch:
			if data == nil {
				return
			}
			session.process(session.msgId, session.seqNo, data)
		}
	}
}
func (session *Session) send(msg TL, resp chan response) error {
	obj := msg.encode()

	x := NewEncodeBuf(256)

	// padding for tcpsize
	x.Int(0)

	if session.c.AuthKey == nil {
		// send un-encrypted packet, e.g., InvokeWithLayer
		x.Long(0)
		x.Long(GenerateMessageId())
		x.Int(int32(len(obj)))
		x.Bytes(obj)
	} else {
		needAck := true
		switch msg.(type) {
		case TL_ping, TL_msgs_ack:
			needAck = false
		}
		z := NewEncodeBuf(256)
		newMsgId := GenerateMessageId()
		z.Bytes(session.c.Salt)
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
		aesKey, aesIV := generateAES(msgKey, session.c.AuthKey, false)

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

		x.Bytes(session.c.AuthKeyHash)
		x.Bytes(msgKey)
		x.Bytes(encryptedData)

		if resp != nil {
			session.mutex.Lock()
			session.msgsIdToResp[newMsgId] = resp
			session.mutex.Unlock()
		}

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

func (session *Session) read() (interface{}, error) {
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
	n, err = tcpconn.Read(b) // Wait for an incoming byte
	if err != nil {
		return nil, err
	}
	slog.Record(b)

	if b[0] < 127 {
		size = int(b[0]) << 2
	} else {
		b := make([]byte, 3)
		n, err = tcpconn.Read(b)
		slog.Record(b)
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
	slog.Record(buf)

	if size == 4 {
		return nil, fmt.Errorf("Server response error: %d", int32(binary.LittleEndian.Uint32(buf)))
	}

	// decrypt incoming packet
	data, session.msgId, session.seqNo, err = decryptMtproto(buf, session.c.AuthKey)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (session *Session) makeAuthKey() ([]byte, []byte, []byte, error) {
	var x []byte
	var err error
	var data interface{}

	// (send) req_pq
	nonceFirst := GenerateNonce(16)
	err = session.send(TL_req_pq{nonceFirst}, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	// (parse) resPQ
	data, err = session.read()
	if err != nil {
		return nil, nil, nil, err
	}
	res, ok := data.(TL_resPQ)
	if !ok {
		return nil, nil, nil, errors.New("Handshake: Need resPQ")
	}
	if !bytes.Equal(nonceFirst, res.nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Nonce")
	}
	found := false
	for _, b := range res.fingerprints {
		if uint64(b) == telegramPublicKey_FP {
			found = true
			break
		}
	}
	if !found {
		return nil, nil, nil, errors.New("Handshake: No fingerprint")
	}

	// (encoding) p_q_inner_data
	p, q := splitPQ(res.pq)
	nonceSecond := GenerateNonce(32)
	nonceServer := res.server_nonce
	innerData1 := (TL_p_q_inner_data{res.pq, p, q, nonceFirst, nonceServer, nonceSecond}).encode()

	x = make([]byte, 255)
	copy(x[0:], sha1(innerData1))
	copy(x[20:], innerData1)
	encryptedData1 := doRSAencrypt(x)
	// (send) req_DH_params
	err = session.send(TL_req_DH_params{nonceFirst, nonceServer, p, q, telegramPublicKey_FP, encryptedData1}, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	// (parse) server_DH_params_{ok, fail}
	data, err = session.read()
	if err != nil {
		return nil, nil, nil, err
	}
	dh, ok := data.(TL_server_DH_params_ok)
	if !ok {
		return nil, nil, nil, errors.New("Handshake: Need server_DH_params_ok")
	}
	if !bytes.Equal(nonceFirst, dh.nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dh.server_nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Server_nonce")
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
	decodedData, err := doAES256IGEdecrypt(dh.encrypted_answer, tmpAESKey, tmpAESIV)
	if err != nil {
		return nil, nil, nil, err
	}
	innerbuf := NewDecodeBuf(decodedData[20:])
	data = innerbuf.Object()
	if innerbuf.err != nil {
		return nil, nil, nil, innerbuf.err
	}
	dhi, ok := data.(TL_server_DH_inner_data)
	if !ok {
		return nil, nil, nil, errors.New("Handshake: Need server_DH_inner_data")
	}
	if !bytes.Equal(nonceFirst, dhi.nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhi.server_nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Server_nonce")
	}

	_, g_b, g_ab := makeGAB(dhi.g, dhi.g_a, dhi.dh_prime)
	authKey := g_ab.Bytes()
	if authKey[0] == 0 {
		authKey = authKey[1:]
	}
	authKeyHash := sha1(authKey)[12:20]
	t4 := make([]byte, 32+1+8)
	copy(t4[0:], nonceSecond)
	t4[32] = 1
	copy(t4[33:], sha1(authKey)[0:8])
	nonceHash1 := sha1(t4)[4:20]
	salt := make([]byte, 8)
	copy(salt, nonceSecond[:8])
	xor(salt, nonceServer[:8])

	// (encoding) client_DH_inner_data
	innerData2 := (TL_client_DH_inner_data{nonceFirst, nonceServer, 0, g_b}).encode()
	x = make([]byte, 20+len(innerData2)+(16-((20+len(innerData2))%16))&15)
	copy(x[0:], sha1(innerData2))
	copy(x[20:], innerData2)
	encryptedData2, err := doAES256IGEencrypt(x, tmpAESKey, tmpAESIV)

	// (send) set_client_DH_params
	err = session.send(TL_set_client_DH_params{nonceFirst, nonceServer, encryptedData2}, nil)
	if err != nil {
		return nil, nil, nil, err
	}

	// (parse) dh_gen_{ok, Retry, fail}
	data, err = session.read()
	if err != nil {
		return nil, nil, nil, err
	}
	dhg, ok := data.(TL_dh_gen_ok)
	if !ok {
		return nil, nil, nil, errors.New("Handshake: Need dh_gen_ok")
	}
	if !bytes.Equal(nonceFirst, dhg.nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhg.server_nonce) {
		return nil, nil, nil, errors.New("Handshake: Wrong Server_nonce")
	}
	if !bytes.Equal(nonceHash1, dhg.new_nonce_hash1) {
		return nil, nil, nil, errors.New("Handshake: Wrong New_nonce_hash1")
	}

	// (all ok)
	//err = session.saveSession()
	//if err != nil {
	//	return nil, err
	//}

	return authKey, authKeyHash, salt, nil
}

func (x *Session) LogPrefix() string {
	return fmt.Sprintf("[%d-%d]", x.connId, x.sessionId)
}

// Implements interface error
func (e TL_rpc_error) Error() string {
	switch e.error_code {
	case errorSeeOther:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.error_code, e.error_message)
	case errorBadRequest, errorUnauthorized, errorFlood, errorInternal:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.error_code, e.error_message)
	default:
		return fmt.Sprintf("mtproto unknow RPC error: %d %s", e.error_code, e.error_message)
	}
}
