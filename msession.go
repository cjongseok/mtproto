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
	"regexp"
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

	appConfig 		Configuration
	user 			*TL_user

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
	session.notify(SessionDiscarded{session.connId, session.sessionId})
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
func loadSession(phonenumber string, appConfig Configuration, sessionListener chan MEvent) (*MSession, error){
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
			err = session.readSessionFile(sessionfile)
			if err == nil {
				session.open(appConfig, sessionListener)
				return session, nil
			} else {
				return nil, fmt.Errorf("Load Session Failure: open new session")
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
				Api_id:         session.appConfig.Id,
				Device_model:   session.appConfig.DeviceModel,
				System_version: session.appConfig.SystemVersion,
				App_version:    session.appConfig.Version,
				Lang_code:      session.appConfig.Language,
				Query:          TL_help_getConfig{},
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
			if session.useIPv6 && v.Ipv6 {
				session.dclist[v.Id] = fmt.Sprintf("[%s]:%d", v.Ip_address, v.Port)
			} else if !v.Ipv6 {
				session.dclist[v.Id] = fmt.Sprintf("%s:%d", v.Ip_address, v.Port)
			}
		}
	default:
		return fmt.Errorf("Connection error: got: %T", x)
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
		// TODO: it doesn't work. think of another solutino to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (session *MSession) process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).Items
		for _, v := range data {
			session.process(v.Msg_id, v.Seq_no, v.Data)
		}

	case TL_bad_server_salt:
		data := data.(TL_bad_server_salt)
		session.serverSalt = data.New_server_salt
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
		session.serverSalt = data.Server_salt
		// save the session on sign in
		_ = session.saveSession()

	case TL_ping:
		//log.Println("process: got ping. send pong")
		logln(session, "got ping. send pong")
		data := data.(TL_ping)
		session.queueSend <- packetToSend{TL_pong{msgId, data.Ping_id}, nil}

	case TL_pong:
		//log.Println("process: got pong")
		logln(session, "got pong")
		// ignore

	case TL_msgs_ack:
		data := data.(TL_msgs_ack)
		session.mutex.Lock()
		defer session.mutex.Unlock()
		for _, v := range data.MsgIds {
			delete(session.msgsIdToAck, v)
		}

	case TL_rpc_result:
		data := data.(TL_rpc_result)
		x := session.process(msgId, seqNo, data.Obj)
		session.mutex.Lock()
		defer session.mutex.Unlock()
		v, ok := session.msgsIdToResp[data.Req_msg_id]
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
		delete(session.msgsIdToAck, data.Req_msg_id)
	default:
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
			//log.Println("PingRoutine: ping")
			logln(session, "ping")
			session.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
		}
	}
}

func (session *MSession) sendRoutine() {
	//log.Println("SendRoutine: start")
	logln(session, "send: start")
	session.sendWaitGroup.Add(1)
	defer session.sendWaitGroup.Done()
	for {
		select {
		case <-session.sendInterrupter:
			//log.Println("SendRoutine: stop")
			logln(session, "send: stop")
			return
		case x := <-session.queueSend:
			//log.Println("SendRoutine: send ", x, x.msg, session)
			if x.msg != nil {
				err := session.sendPacket(x.msg, x.resp)
				if err != nil {
					//log.Fatalln("SendRoutine:", err)
					fatalln(session, "send: ", err)
				}
			}
		}
	}
}

func (session *MSession) readRoutine() {
	//log.Println("[ReadRoutine: start")
	logln(session, "read: start")
	session.readWaitGroup.Add(1)
	defer session.readWaitGroup.Done()

	innerRoutineWG := sync.WaitGroup{}

	for {
		// Run async wait for data from server
		ch := make(chan interface{}, 1)
		go func(ch chan<- interface{}) {
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := session.read()
			//log.Printf("[%d-%d] read: data: %v, err: %v\n", session.connId, session.sessionId, data, err)
			logf(session, "read: data: %v, err: %v\n", data, err)
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				logf(session, "read: lost connection (captured EOF). reconnect to %s\n", session.addr)
				//err = m.reconnect(m.addr)
				//session.notify(RECONNECT)
				session.notify(refreshSession{
					session.sessionId,
					session.phonenumber,
					nil})
			} else if err != nil {
				//log.Println(err.Error())
				logf(session, "read %v\n", err.Error())
				re, reerr := regexp.Compile("use of closed network connection")
				if reerr != nil {
					//log.Printf("[%d-%d] read: regexp compilation error, %v", session.connId, session.sessionId, reerr)
					logf(session, "read: regexp compilation error, %v", reerr)
				}
				if re.MatchString(err.Error()) {
					//log.Printf("[%d-%d] read: TCP connection closed, %v", session.connId, session.sessionId, err)
					logf(session, "read: TCP connection closed, %v", err)
				} else {
					//log.Fatalln("Unknown session error, ", err)
					fatalln(session, "read: Unknown session error, ", err)
				}
			} else {
				ch <- data
			}
		}(ch)

		select {
		case <-session.readInterrupter:
		//case <-m.stopRoutines:
			//TODO: how about to close conn here?
			//log.Println("ReadRoutine: wait for inner routine")
			logln(session, "read: wait for inner routine ...")
			innerRoutineWG.Wait()
			//log.Println("ReadRoutine: stop")
			logln(session, "read: stop")
			return
		case data := <-ch:
			if data == nil {
				return
			}
			session.process(session.msgId, session.seqNo, data)
		}
	}
}

