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
	connId      int32
	sessionId   int64
	phonenumber string
	addr        string
	useIPv6     bool
	listeners   []chan Event
	tcpconn     *net.TCPConn
	f           *os.File
	queueSend   chan packetToSend

	readInterrupter chan struct{}
	sendInterrupter chan struct{}
	pingInterrupter chan struct{}

	isReading bool
	isSending bool
	isPing    bool

	readWaitGroup sync.WaitGroup
	sendWaitGroup sync.WaitGroup
	pingWaitGroup sync.WaitGroup

	authKey     []byte
	authKeyHash []byte
	serverSalt  []byte
	encrypted   bool

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

	dclist map[int32]string
}

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

func newSession(phonenumber string, addr string, useIPv6 bool, appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event) (*Session, error) {
	var err error

	session := new(Session)
	session.phonenumber = phonenumber
	session.f, err = os.OpenFile(appConfig.KeyPath, os.O_WRONLY|os.O_CREATE, 0600)
	if err == nil {
		session.addr = addr
		session.useIPv6 = useIPv6
		session.encrypted = false
		err = session.open(appConfig /*sendQueue,*/, sessionListener, false)
		if err != nil {
			return nil, err
		}
		return session, nil
	}
	return nil, err
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

// Build a connection from the session file
// returned session contains the same session with the file but session id,
// since the session file does not have session id
func loadSession(phonenumber string, preferredAddr string, appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event) (*Session, error) {
	// session file exists?
	//sessionfile := sessionFilePath(appConfig.SessionHome, phonenumber)
	//_, err := os.Stat(sessionfile)
	//sessionExists := !os.IsNotExist(err)
	//err = nil

	// load session info from either session file or env
	// its precedence is; preferredAddr > sessionFile > env
	session := new(Session)
	session.phonenumber = phonenumber
	var err error
	if appConfig.KeyPath == "" {
		// load key from env
		session.authKey = byteArrayString2byteArray(os.Getenv(ENV_AUTHKEY))
		session.authKeyHash = byteArrayString2byteArray(os.Getenv(ENV_AUTHHASH))
		session.serverSalt = byteArrayString2byteArray(os.Getenv(ENV_SERVER_SALT))
		session.addr = os.Getenv(ENV_ADDR)
		session.useIPv6, _ = strconv.ParseBool(os.Getenv(ENV_USE_IPV6))
		session.encrypted = true
		if session.authKey == nil || session.authKeyHash == nil || session.serverSalt == nil || session.addr == "" {
			return nil, fmt.Errorf("cannot find mtproto key from neither session file nor env: open new session: %v", err)
		} else {
			session.f, err = os.OpenFile(appConfig.KeyPath, os.O_WRONLY|os.O_CREATE, 0600)
		}
	} else {
		session.f, err = os.OpenFile(appConfig.KeyPath, os.O_RDONLY, 0600)
		if err == nil {
			err = session.readSessionFile(session.f)
		}
		if err != nil {
			return nil, fmt.Errorf("read mtproto key failure: %v", err)
		}
	}

	if preferredAddr != "" {
		tcpAddr, err := net.ResolveTCPAddr("tcp", preferredAddr)
		if err != nil {
			return nil, fmt.Errorf("resolve the telegram server address failure: %v", err)
		}
		if tcpAddr.IP.To4() != nil {
			session.useIPv6 = false
			session.addr = preferredAddr
		} else if tcpAddr.IP.To16() != nil {
			session.useIPv6 = true
			session.addr = preferredAddr
		} else {
			// Invalid IP address. Ignore the preferred ip address
			slog.Logln(session, "invalid preferred preferred Telegram server address. ignore it.")
		}
	}
	err = session.open(appConfig /*sendQueue,*/, sessionListener, true)
	if err != nil {
		return session, handshakingFailure{fmt.Sprintf("Handshaking Failure: %v", err)}
	}
	return session, nil
}

func (session *Session) open(appConfig Configuration /*sendQueue chan packetToSend,*/, sessionListener chan Event, getUpdateStates bool) error {
	var err error
	var tcpAddr *net.TCPAddr

	// set up rest of session
	session.appConfig = appConfig
	rand.Seed(time.Now().UnixNano())
	session.sessionId = rand.Int63()
	session.AddSessionListener(sessionListener)

	// connect
	tcpAddr, err = net.ResolveTCPAddr("tcp", session.addr)
	if err != nil {
		return err
	}
	slog.Logf(session, "dial TCP to %s\n", session.addr)
	session.tcpconn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	// Packet Length is encoded by a single byte (see: https://core.telegram.org/mtproto)
	_, err = session.tcpconn.Write([]byte{0xef})
	if err != nil {
		return err
	}
	// get new authKey if need
	if !session.encrypted {
		err = session.makeAuthKey()
		if err != nil {
			return err
		}
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
				ApiId:          session.appConfig.Id,
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
			return fmt.Errorf("TL_invokeWithLayer Failure:", x.err)
		}
	case <-time.After(TIMEOUT_INVOKE_WITH_LAYER):
		return fmt.Errorf("TL_invokeWithLayer Timeout(%f s)", TIMEOUT_INVOKE_WITH_LAYER.Seconds())
		//slog.Logf(session, "TL_invokeWithLayer Timeout(%f s)\n", TIMEOUT_INVOKE_WITH_LAYER.Seconds())
	}

	switch x.data.(type) {
	case *PredConfig:
		session.dclist = make(map[int32]string, 5)
		for _, v := range x.data.(*PredConfig).DcOptions {
			isIPv6 := true
			dcOption := v.GetValue()
			tcpAddr, err := net.ResolveIPAddr("ip", dcOption.GetIpAddress())
			if err == nil {
				ip := tcpAddr.IP.To4()
				if ip != nil {
					isIPv6 = false
				}
				if session.useIPv6 {
					if isIPv6 {
						session.dclist[dcOption.GetId()] = fmt.Sprintf("[%s]:%d", dcOption.GetIpAddress(), dcOption.GetPort())
					}
				} else {
					if !isIPv6 {
						session.dclist[dcOption.GetId()] = fmt.Sprintf("%s:%d", dcOption.GetIpAddress(), dcOption.GetPort())
					}
				}
			}
		}
		marshaled, err := json.Marshal(x.data)
		if err == nil {
			slog.Logf(session, "config: %s\n", marshaled)
		}
	default:
		return fmt.Errorf("Connection error: Failed to get config. got: %T", x)
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
				return fmt.Errorf("TL_updates_getState Failure: %s", x.err)
			}
		case <-time.After(TIMEOUT_UPDATES_GETSTATE):
			//session.close()
			return fmt.Errorf("TL_updates_getState Timeout(%f s)", TIMEOUT_UPDATES_GETSTATE.Seconds())
		}
	}

	// start keep alive ping
	session.isPing = true
	session.pingWaitGroup.Add(1)
	go session.pingRoutine()

	// notify the connection established
	session.notify(SessionEstablished{session})

	//fmt.Println("authKey:", session.authKey)
	//fmt.Println("authKeyHash:", session.authKeyHash)
	//fmt.Println("salt:", session.serverSalt)
	//fmt.Println("addf:", session.addr)
	//fmt.Println("ipv6:", session.useIPv6)

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
	return fmt.Errorf("Listener (%x) doesn't exist", toremove)
}

func (session *Session) readSessionFile(f *os.File) error {
	// Decode session file
	b := make([]byte, 1024*4)
	n, err := f.ReadAt(b, 0)
	if n <= 0 || (err != nil && err.Error() != "EOF") {
		return errors.New("New session")
	}

	d := NewDecodeBuf(b)
	session.authKey = d.StringBytes()
	session.authKeyHash = d.StringBytes()
	session.serverSalt = d.StringBytes()
	session.addr = d.String()
	session.useIPv6 = false
	if d.UInt() == 1 {
		session.useIPv6 = true
	}

	if d.err != nil {
		// Failed to load session
		return d.err
	}

	session.encrypted = true
	return nil
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
			session.serverSalt = data.new_server_salt
			// save the session on sign in
			_ = session.saveSession()
			session.mutex.Lock()
			defer session.mutex.Unlock()
			for k, v := range session.msgsIdToAck {
				delete(session.msgsIdToAck, k)
				session.queueSend <- v
			}

		case TL_new_session_created:
			data := data.(TL_new_session_created)
			session.serverSalt = data.server_salt
			// save the session on sign in
			_ = session.saveSession()

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

// Save session
//TODO: save channel and datacenter information
func (session *Session) saveSession() (err error) {
	session.encrypted = true

	b := NewEncodeBuf(1024)
	b.StringBytes(session.authKey)
	b.StringBytes(session.authKeyHash)
	b.StringBytes(session.serverSalt)
	b.String(session.addr)
	var useIPv6UInt uint32
	if session.useIPv6 {
		useIPv6UInt = 1
	}
	b.UInt(useIPv6UInt)

	err = session.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = session.f.WriteAt(b.buf, 0)
	if err != nil {
		return err
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
				err := session.sendPacket(x.msg, x.resp)
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
			refresh := func(session *Session) {
				session.notify(refreshSession{
					session.sessionId,
					session.phonenumber,
					nil,
				})
			}
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := session.read()
			if _, ok := data.(TL_pong); !ok {
				slog.Logf(session, "read: type: %v, data: %v, err: %v\n", reflect.TypeOf(data), data, err)
			}
			//slog.Logf(session, "read: %s\n", slog.Stringify(data))
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				slog.Logf(session, "read: lost connection (captured EOF). reconnect to %s\n", session.addr)
				refresh(session)
			} else if err != nil {
				if strings.Contains(err.Error(), "use of closed network connection") {
					slog.Logf(session, "read: TCP connection closed (%s)\n", err)
					// Two cases
					// 1. on new authentication, 303 PHONE_MIGRATE can require to make a new connection with different
					//   server by closing the connection. -> do nothing, because session will be renewed by MM
					// 2. after authentication, there could be an accidental disconnection. -> need to refresh
					if session.user == nil {
						// case 1
						// do nothing
					} else {
						// case 2
						refresh(session)
					}
				} else if strings.Contains(err.Error(), "connection reset by peer") {
					slog.Logf(session, "read: lost connection (%s). reconnect to %s\n", err, session.addr)
					refresh(session)
				} else if strings.Contains(err.Error(), "i/o timeout") {
					slog.Logf(session, "read: lost connection (%s). reconnect to %s\n", err, session.addr)
					refresh(session)
				} else {
					slog.Logf(session, "read: unknown error, %s. reconnect to %s\n", err, session.addr)
					refresh(session)
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
func (session *Session) sendPacket(msg TL, resp chan response) error {
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
	data, session.msgId, session.seqNo, err = decryptMtproto(buf, session.authKey)
	if err != nil {
		return nil, err
	}
	return data, nil

}

func (session *Session) makeAuthKey() error {
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
	if !bytes.Equal(nonceFirst, res.nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	found := false
	for _, b := range res.fingerprints {
		if uint64(b) == telegramPublicKey_FP {
			found = true
			break
		}
	}
	if !found {
		return errors.New("Handshake: No fingerprint")
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
	if !bytes.Equal(nonceFirst, dh.nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dh.server_nonce) {
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
	decodedData, err := doAES256IGEdecrypt(dh.encrypted_answer, tmpAESKey, tmpAESIV)
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
	if !bytes.Equal(nonceFirst, dhi.nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhi.server_nonce) {
		return errors.New("Handshake: Wrong Server_nonce")
	}

	_, g_b, g_ab := makeGAB(dhi.g, dhi.g_a, dhi.dh_prime)
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
	if !bytes.Equal(nonceFirst, dhg.nonce) {
		return errors.New("Handshake: Wrong Nonce")
	}
	if !bytes.Equal(nonceServer, dhg.server_nonce) {
		return errors.New("Handshake: Wrong Server_nonce")
	}
	if !bytes.Equal(nonceHash1, dhg.new_nonce_hash1) {
		return errors.New("Handshake: Wrong New_nonce_hash1")
	}

	// (all ok)
	err = session.saveSession()
	if err != nil {
		return err
	}

	return nil
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
