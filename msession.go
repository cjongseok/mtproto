package mtproto

import (
	"net"
	"os"
	"sync"
	"fmt"
	"errors"
	"math/rand"
	"time"
	"io"
	"encoding/json"
	"strings"
	"github.com/cjongseok/slog"
)

const (
	// API Errors
	errorSeeOther     = 303
	errorBadRequest   = 400
	errorUnauthorized = 401
	errorForbidden    = 403
	errorNotFound     = 404
	errorFlood        = 420
	errorInternal     = 500
)

type MSession struct {
	connId 			int32
	sessionId		int64
	phonenumber		string
	addr         	string
	useIPv6      	bool
	listeners	 	[]chan MEvent
	tcpconn      	*net.TCPConn
	f            	*os.File
	queueSend    	chan packetToSend

	readInterrupter	chan struct{}
	sendInterrupter chan struct{}
	pingInterrupter chan struct{}

	readWaitGroup	sync.WaitGroup
	sendWaitGroup	sync.WaitGroup
	pingWaitGroup	sync.WaitGroup

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

	appConfig    Configuration
	user         *TL_user
	updatesState *TL_updates_state

	dclist map[int32]string
}

type packetToSend struct {
	msg  TL
	resp chan response
}

type response struct {
	data TL
	err  error
}

// CHECK: Only mmanager call this method?
func (session *MSession) close() {
	session.stopPing()
	session.stopSend()
	session.pingWaitGroup.Wait()
	session.sendWaitGroup.Wait()

	session.tcpconn.Close()

	session.stopRead()
	session.readWaitGroup.Wait()

	// notify that the connection is gracefully closed
	session.notify(SessionDiscarded{session.connId, session.sessionId, *session.updatesState})
	session.listeners = nil
}

func sessionFilePath(sessionFileHome string, phonenumber string) string {
	return sessionFileHome + "/.telegram_" + phonenumber
}

func newSession(phonenumber string, addr string, useIPv6 bool, appConfig Configuration, sessionListener chan MEvent) (*MSession, error) {
	var err error

	session := new(MSession)
	session.phonenumber = phonenumber
	session.f, err = os.OpenFile(sessionFilePath(appConfig.SessionHome, phonenumber), os.O_WRONLY|os.O_CREATE, 0600)
	if err == nil {
		session.addr = addr
		session.useIPv6 = useIPv6
		session.encrypted = false
		session.open(appConfig, sessionListener)
		return session, nil
	}
	return nil, err
}

// Build a connection from the session file
// returned session contains the same session with the file but session id,
// since the session file does not have session id
func loadSession(phonenumber string, preferredAddr string, appConfig Configuration, sessionListener chan MEvent) (*MSession, error){
	// session file exists?
	sessionfile := sessionFilePath(appConfig.SessionHome, phonenumber)
	_, err := os.Stat(sessionfile)
	sessionExists := !os.IsNotExist(err)

	// load session from session file
	if sessionExists {
		session := new(MSession)
		session.phonenumber = phonenumber
		session.f, err = os.OpenFile(sessionfile, os.O_RDONLY, 0600)
		if err == nil {
			if preferredAddr != "" {
				tcpAddr, err := net.ResolveTCPAddr("tcp", preferredAddr)
				if err == nil {
					if tcpAddr.IP.To4() != nil {
						session.useIPv6 = false
						session.addr = preferredAddr
					} else if tcpAddr.IP.To16() != nil{
						session.useIPv6 = true
						session.addr = preferredAddr
					} else {
						// Invalid IP address. Ignore the preferred ip address
					}
				}
			}
			err = session.readSessionFile(sessionfile)
			if err == nil {
				err = session.open(appConfig, sessionListener)
				if err == nil {
					return session, nil
				}
				return nil, fmt.Errorf("Handshaking Failure: %v", err)
			} else {
				return nil, fmt.Errorf("Cannot Read Session File: open new session")
			}
		}
	}

	// no session file
	return nil, fmt.Errorf("Load Session Failure: cannot find session file %s", sessionfile)
}

func (session *MSession) open(appConfig Configuration, sessionListener chan MEvent) error {
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
	session.sendInterrupter = make(chan struct{})
	session.readInterrupter = make(chan struct{})
	session.pingInterrupter = make(chan struct{})

	session.sendWaitGroup = sync.WaitGroup{}
	session.readWaitGroup = sync.WaitGroup{}
	session.pingWaitGroup = sync.WaitGroup{}

	session.msgsIdToAck = make(map[int64]packetToSend)
	session.msgsIdToResp = make(map[int64]chan response)
	session.mutex = &sync.Mutex{}
	go session.sendRoutine()
	go session.readRoutine()

	// (help_getConfig)
	resp := make(chan response, 1)
	session.queueSend <- packetToSend{
		msg: TL_invokeWithLayer{
			Layer: layer,
			Query: TL_initConnection{
				Api_id:         	session.appConfig.Id,
				Device_model:   	session.appConfig.DeviceModel,
				System_version: 	session.appConfig.SystemVersion,
				App_version:    	session.appConfig.Version,
				System_lang_code: 	session.appConfig.Language,
				Lang_code:      	session.appConfig.Language,
				Query:          	TL_help_getConfig{},
			},
		},
		resp: resp,
	}
	x := <-resp
	if x.err != nil {
		return x.err
	}

	switch x.data.(type) {
	case TL_config:
		session.dclist = make(map[int32]string, 5)
		for _, v := range x.data.(TL_config).Dc_options {
			v := v.(TL_dcOption)
			isIPv6 := false
			tcpAddr, err := net.ResolveTCPAddr("tcp", v.Ip_address)
			if err == nil {
				if tcpAddr.IP.To16() != nil{
					isIPv6 = true
				}
			}
			if session.useIPv6 && isIPv6 {
				session.dclist[v.Id] = fmt.Sprintf("[%s]:%d", v.Ip_address, v.Port)
			} else if !isIPv6 {
				session.dclist[v.Id] = fmt.Sprintf("%s:%d", v.Ip_address, v.Port)
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
	//TODO: From second session, query getUpdatesState with invokeWithLayer and initConnection
	session.updatesState = new(TL_updates_state)
	resp = make(chan response, 1)
	session.queueSend <- packetToSend{
		msg: TL_updates_getState{},
		resp: resp,
	}
	x = <- resp
	if x.err != nil {
		return x.err
	}

	// start keep alive ping
	go session.pingRoutine()

	// notify the connection established
	session.notify(SessionEstablished{session})

	return nil
}

func (session *MSession) AddSessionListener(listener chan MEvent) {
	session.listeners = append(session.listeners, listener)
}

func (session *MSession) RemoveSessionListener(toremove chan MEvent) error {
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

func (session *MSession) readSessionFile(sessionfile string) error {
	// Decode session file
	b := make([]byte, 1024*4)
	n, err := session.f.ReadAt(b, 0)
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

func (session *MSession) notify(e MEvent) {
	for _, listener := range session.listeners {
		// TODO: it doesn't work. think of another solution to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (session *MSession) process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).items
		for _, v := range data {
			session.process(v.msg_id, v.seq_no, v.data)
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
		data := data.(TL_rpc_result)
		x := session.process(msgId, seqNo, data.obj)
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
				}
				resp.data = x.(TL)
				v <- resp
				close(v)
			}()
		}
		delete(session.msgsIdToAck, data.req_msg_id)

	//TODO: Update classifier should be auto-generated from scheme.tl
	//TODO: how to handle seq?
	// Date, Pts, Qts updates
	case TL_updates_state:
		data := data.(TL_updates_state)
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
	case TL_updates:
		data := data.(TL_updates)
		session.updatesState.Date = data.Date
		session.updatesState.Seq = data.Seq
		session.notify(updateReceived{data})
		return data
	case TL_updateShort:
		data := data.(TL_updateShort)
		//session.updatesState.Pts ++	//TODO: need to comment in it?
		session.updatesState.Date = data.Date
		session.notify(updateReceived{data})
		return data

	// Pts updates
	case TL_updateNewMessage:
		data := data.(TL_updateNewMessage)
		session.updatesState.Pts = data.Pts
		session.notify(updateReceived{data})
		return data
	case TL_updateReadMessagesContents:
		data := data.(TL_updateReadMessagesContents)
		session.updatesState.Pts = data.Pts
		session.notify(updateReceived{data})
		return data
	case TL_updateDeleteMessages:
		data := data.(TL_updateDeleteMessages)
		session.updatesState.Pts = data.Pts
		session.notify(updateReceived{data})
		return data

	// Pts and Date updates
	case TL_updateShortMessage:
		data := data.(TL_updateShortMessage)
		session.updatesState.Pts = data.Pts
		session.updatesState.Date = data.Date
		session.notify(updateReceived{data})
		return data
	case TL_updateShortChatMessage:
		data := data.(TL_updateShortChatMessage)
		session.updatesState.Pts = data.Pts
		session.updatesState.Date = data.Date
		session.notify(updateReceived{data})
		return data
	case TL_updateShortSentMessage:
		data := data.(TL_updateShortSentMessage)
		session.updatesState.Pts = data.Pts
		session.updatesState.Date = data.Date
		session.notify(updateReceived{data})
		return data

	// Qts updates
	case TL_updateNewEncryptedMessage:
		data := data.(TL_updateNewEncryptedMessage)
		session.updatesState.Qts = data.Qts
		session.notify(updateReceived{data})
		return data

	// Channel updates
	case TL_updateChannel:
		data := data.(TL_updateChannel)
		session.notify(updateReceived{data})
		return data
	case TL_updateChannelMessageViews:
		data := data.(TL_updateChannelMessageViews)
		session.notify(updateReceived{data})
		return data
	case TL_updateChannelTooLong:
		data := data.(TL_updateChannelTooLong)
		session.updatesState.Pts = data.Pts
		session.notify(updateReceived{data})
		return data
	case TL_updateReadChannelInbox:
		data := data.(TL_updateReadChannelInbox)
		session.notify(updateReceived{data})
		return data
	case TL_updateReadChannelOutbox:
		data := data.(TL_updateReadChannelOutbox)
		session.notify(updateReceived{data})
		return data
	case TL_updateNewChannelMessage:
		data := data.(TL_updateNewChannelMessage)
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

	// TODO: Check why I should do this
	if (seqNo & 1) == 1 {
		session.queueSend <- packetToSend{TL_msgs_ack{[]int64{msgId}}, nil}
	}

	return nil
}

// Save session
//TODO: save channel and datacenter information
func (session *MSession) saveSession() (err error) {
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

func (session *MSession) stopRead() {
	close(session.readInterrupter)
}

func (session *MSession) stopSend() {
	close(session.sendInterrupter)
	close(session.queueSend)
}

func (session *MSession) stopPing() {
	close(session.pingInterrupter)
}

func (session *MSession) pingRoutine() {
	session.pingWaitGroup.Add(1)
	defer session.pingWaitGroup.Done()
	for {
		select {
		case <-session.pingInterrupter:
			return
		case <-time.After(60 * time.Second):
			slog.Logln(session, "ping")
			session.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
		}
	}
}

func (session *MSession) sendRoutine() {
	slog.Logln(session, "send: start")
	session.sendWaitGroup.Add(1)
	defer session.sendWaitGroup.Done()
	for {
		select {
		case <-session.sendInterrupter:
			slog.Logln(session, "send: stop")
			return
		case x := <-session.queueSend:
			//slog.Logf(session, "send: type: %v, data: %v", reflect.TypeOf(x.msg), x.msg)
			slog.Logf(session, "send %s\n", slog.Stringify(x.msg))
			if x.msg != nil {
				err := session.sendPacket(x.msg, x.resp)
				if err != nil {
					slog.Fatalln(session, "send: ", err)
				}
			}
		}
	}
}

func (session *MSession) readRoutine() {
	slog.Logln(session, "read: start")
	session.readWaitGroup.Add(1)
	defer session.readWaitGroup.Done()

	innerRoutineWG := sync.WaitGroup{}

	for {
		// Run async wait for data from server
		ch := make(chan interface{}, 1)
		go func(ch chan<- interface{}) {
			refresh := func(session *MSession) {
				session.notify(refreshSession{
					session.sessionId,
					session.phonenumber,
					nil,
				})
			}
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := session.read()
			//slog.Logf(session, "read: type: %v, data: %v, err: %v\n", reflect.TypeOf(data), data, err)
			slog.Logf(session, "read: %s\n", slog.Stringify(data))
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
					refresh(session)
				} else if strings.Contains(err.Error(), "connection reset by peer") {
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

// Implements interface error
func (e TL_rpc_error) Error() string {
	switch e.error_code{
	case errorSeeOther:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.error_code, e.error_message)
	case errorBadRequest, errorUnauthorized, errorFlood, errorInternal:
		return fmt.Sprintf("mtproto RPC error: %d %s", e.error_code, e.error_message)
	default:
		return fmt.Sprintf("mtproto unknow RPC error: %d %s", e.error_code, e.error_message)
	}
}

