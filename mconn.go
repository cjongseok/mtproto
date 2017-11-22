package mtproto

import (
	"net"
	"os"
	"sync"
	"log"
	"fmt"
	"errors"
	"math/rand"
	"time"
	"io"
	"regexp"
)

type MConn struct {

	phonenumber		string
	listeners	 	[]chan MEvent
	addr         	string
	useIPv6      	bool
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
	sessionId	int64

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

func (mconn *MConn) Close() {
	mconn.stopPing()
	mconn.stopSend()
	mconn.pingWaitGroup.Wait()
	mconn.sendWaitGroup.Wait()

	mconn.tcpconn.Close()

	mconn.stopRead()
	mconn.readWaitGroup.Wait()

	// notify the connection gracefully closed
	mconn.notify(Disconnected{mconn.sessionId})
}

func sessionFilePath(sessionFileHome string, phonenumber string) string {
	return sessionFileHome + "/.telegram_" + phonenumber
}

func newSession(phonenumber string, addr string, useIPv6 bool, appConfig Configuration, connListener chan MEvent) (*MConn, error) {
	var err error

	mconn := new(MConn)
	mconn.phonenumber = phonenumber
	mconn.f, err = os.OpenFile(sessionFilePath(appConfig.SessionHome, phonenumber), os.O_WRONLY|os.O_CREATE, 0600)
	if err == nil {
		mconn.addr = addr
		mconn.useIPv6 = useIPv6
		mconn.encrypted = false
		mconn.open(appConfig, connListener)
		return mconn, nil
	}
	return nil, err
}

// Build a connection from the session file
// returned mconn contains the same session with the file but session id,
// since the session file does not have session id
func loadSession(phonenumber string, appConfig Configuration, connListener chan MEvent) (*MConn, error){
	// session file exists?
	sessionfile := sessionFilePath(appConfig.SessionHome, phonenumber)
	_, err := os.Stat(sessionfile)
	sessionExists := !os.IsNotExist(err)

	// load mconn from session file
	if sessionExists {
		mconn := new(MConn)
		mconn.phonenumber = phonenumber
		mconn.f, err = os.OpenFile(sessionfile, os.O_RDONLY, 0600)
		if err == nil {
			err = mconn.readSessionFile(sessionfile)
			if err == nil {
				mconn.open(appConfig, connListener)
				return mconn, nil
			} else {
				return nil, fmt.Errorf("Load Session Failure: open new session")
			}
		}
	}

	// no session file
	return nil, fmt.Errorf("Load Session Failure: cannot find session file %s", sessionfile)
}

func (mconn *MConn) open(appConfig Configuration, connListener chan MEvent) error {
	var err error
	var tcpAddr *net.TCPAddr

	// set up rest of mconn
	mconn.appConfig = appConfig
	rand.Seed(time.Now().UnixNano())
	mconn.sessionId = rand.Int63()
	mconn.addConnListener(connListener)

	// connect
	tcpAddr, err = net.ResolveTCPAddr("tcp", mconn.addr)
	if err != nil {
		return err
	}
	mconn.tcpconn, err = net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return err
	}
	// Packet Length is encoded by a single byte (see: https://core.telegram.org/mtproto)
	_, err = mconn.tcpconn.Write([]byte{0xef})
	if err != nil {
		return err
	}
	// get new authKey if need
	if !mconn.encrypted {
		err = mconn.makeAuthKey()
		if err != nil {
			return err
		}
	}

	// start goroutines
	mconn.queueSend = make(chan packetToSend, 64)
	mconn.sendInterrupter = make(chan struct{})
	mconn.readInterrupter = make(chan struct{})
	mconn.pingInterrupter = make(chan struct{})

	mconn.sendWaitGroup = sync.WaitGroup{}
	mconn.readWaitGroup = sync.WaitGroup{}
	mconn.pingWaitGroup = sync.WaitGroup{}

	mconn.msgsIdToAck = make(map[int64]packetToSend)
	mconn.msgsIdToResp = make(map[int64]chan response)
	mconn.mutex = &sync.Mutex{}
	go mconn.sendRoutine()
	go mconn.readRoutine()

	// (help_getConfig)
	resp := make(chan response, 1)
	mconn.queueSend <- packetToSend{
		msg: TL_invokeWithLayer{
			Layer: layer,
			Query: TL_initConnection{
				Api_id:         mconn.appConfig.Id,
				Device_model:   mconn.appConfig.DeviceModel,
				System_version: mconn.appConfig.SystemVersion,
				App_version:    mconn.appConfig.Version,
				Lang_code:      mconn.appConfig.Language,
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
		mconn.dclist = make(map[int32]string, 5)
		for _, v := range x.data.(TL_config).Dc_options {
			v := v.(TL_dcOption)
			if mconn.useIPv6 && v.Ipv6 {
				mconn.dclist[v.Id] = fmt.Sprintf("[%s]:%d", v.Ip_address, v.Port)
			} else if !v.Ipv6 {
				mconn.dclist[v.Id] = fmt.Sprintf("%s:%d", v.Ip_address, v.Port)
			}
		}
	default:
		return fmt.Errorf("Connection error: got: %T", x)
	}

	// start keep alive ping
	go mconn.pingRoutine()

	// notify the connection established
	mconn.notify(Connected{mconn.sessionId})

	return nil
}

func (mconn *MConn) addConnListener(connListener chan MEvent) {
	mconn.listeners = append(mconn.listeners, connListener)
}

func (mconn *MConn) removeConnListener(toremove chan MEvent) error {
	// find the listener index in the list
	for index, registered := range mconn.listeners {
		if registered == toremove {
			// remove the element by the index
			copy(mconn.listeners[index:], mconn.listeners[index+1:])
			mconn.listeners[len(mconn.listeners)-1] = nil
			mconn.listeners = mconn.listeners[:len(mconn.listeners)-1]
			return nil
		}
	}
	return fmt.Errorf("Connection Listener (%x) doesn't exist", toremove)
}

func (mconn *MConn) readSessionFile(sessionfile string) error {
	// Decode session file
	b := make([]byte, 1024*4)
	n, err := mconn.f.ReadAt(b, 0)
	if n <= 0 || (err != nil && err.Error() != "EOF") {
		return errors.New("New session")
	}

	d := NewDecodeBuf(b)
	mconn.authKey = d.StringBytes()
	mconn.authKeyHash = d.StringBytes()
	mconn.serverSalt = d.StringBytes()
	mconn.addr = d.String()
	mconn.useIPv6 = false
	if d.UInt() == 1 {
		mconn.useIPv6 = true
	}

	if d.err != nil {
		// Failed to load session
		return d.err
	}

	mconn.encrypted = true
	return nil
}

func (mconn *MConn) notify(event MEvent) {
	for _, listener := range mconn.listeners {
		listener <- event
	}
}

func (mconn *MConn) process(msgId int64, seqNo int32, data interface{}) interface{} {
	switch data.(type) {
	case TL_msg_container:
		data := data.(TL_msg_container).Items
		for _, v := range data {
			mconn.process(v.Msg_id, v.Seq_no, v.Data)
		}

	case TL_bad_server_salt:
		data := data.(TL_bad_server_salt)
		mconn.serverSalt = data.New_server_salt
		// save the session on sign in
		_ = mconn.saveSession()
		mconn.mutex.Lock()
		defer mconn.mutex.Unlock()
		for k, v := range mconn.msgsIdToAck {
			delete(mconn.msgsIdToAck, k)
			mconn.queueSend <- v
		}

	case TL_new_session_created:
		data := data.(TL_new_session_created)
		mconn.serverSalt = data.Server_salt
		// save the session on sign in
		_ = mconn.saveSession()

	case TL_ping:
		log.Println("process: got ping. send pong")
		data := data.(TL_ping)
		mconn.queueSend <- packetToSend{TL_pong{msgId, data.Ping_id}, nil}

	case TL_pong:
		log.Println("process: got pong")
		// ignore

	case TL_msgs_ack:
		data := data.(TL_msgs_ack)
		mconn.mutex.Lock()
		defer mconn.mutex.Unlock()
		for _, v := range data.MsgIds {
			delete(mconn.msgsIdToAck, v)
		}

	case TL_rpc_result:
		data := data.(TL_rpc_result)
		x := mconn.process(msgId, seqNo, data.Obj)
		mconn.mutex.Lock()
		defer mconn.mutex.Unlock()
		v, ok := mconn.msgsIdToResp[data.Req_msg_id]
		if ok {
			go func() {
				var resp response
				rpcError, ok := x.(TL_rpc_error)
				if ok {
					//resp.err = mconn.handleRPCError(rpcError)
					resp.err = rpcError
				}
				resp.data = x.(TL)
				v <- resp
				close(v)
			}()
		}
		delete(mconn.msgsIdToAck, data.Req_msg_id)
	default:
		return data
	}

	// TODO: Check why I should do this
	if (seqNo & 1) == 1 {
		mconn.queueSend <- packetToSend{TL_msgs_ack{[]int64{msgId}}, nil}
	}

	return nil
}

// Save session
//TODO: save channel and datacenter information
func (mconn *MConn) saveSession() (err error) {
	mconn.encrypted = true

	b := NewEncodeBuf(1024)
	b.StringBytes(mconn.authKey)
	b.StringBytes(mconn.authKeyHash)
	b.StringBytes(mconn.serverSalt)
	b.String(mconn.addr)
	var useIPv6UInt uint32
	if mconn.useIPv6 {
		useIPv6UInt = 1
	}
	b.UInt(useIPv6UInt)

	err = mconn.f.Truncate(0)
	if err != nil {
		return err
	}

	_, err = mconn.f.WriteAt(b.buf, 0)
	if err != nil {
		return err
	}

	return nil
}

func (mconn *MConn) InvokeBlocked(msg TL) (*TL, error) {
	x := <-mconn.InvokeNonBlocked(msg)

	if x.err != nil {
		return nil, x.err
	}

	return &x.data, nil
}

func (mconn *MConn) InvokeNonBlocked(msg TL) chan response {
	resp := make(chan response, 1)
	mconn.queueSend <- packetToSend{
		msg:  msg,
		resp: resp,
	}
	return resp
}

func (mconn *MConn) stopRead() {
	close(mconn.readInterrupter)
}

func (mconn *MConn) stopSend() {
	close(mconn.sendInterrupter)
	close(mconn.queueSend)
}

func (mconn *MConn) stopPing() {
	close(mconn.pingInterrupter)
}

func (mconn *MConn) pingRoutine() {
	mconn.pingWaitGroup.Add(1)
	defer mconn.pingWaitGroup.Done()
	for {
		select {
		case <-mconn.pingInterrupter:
			return
		case <-time.After(60 * time.Second):
			log.Println("PingRoutine: ping")
			mconn.queueSend <- packetToSend{TL_ping{0xCADACADA}, nil}
		}
	}
}

func (mconn *MConn) sendRoutine() {
	log.Println("SendRoutine: start")
	mconn.sendWaitGroup.Add(1)
	defer mconn.sendWaitGroup.Done()
	for {
		select {
		case <-mconn.sendInterrupter:
			log.Println("SendRoutine: stop")
			return
		case x := <-mconn.queueSend:
			//log.Println("SendRoutine: send ", x, x.msg, mconn)
			if x.msg != nil {
				err := mconn.sendPacket(x.msg, x.resp)
				if err != nil {
					log.Fatalln("SendRoutine:", err)
				}
			}
		}
	}
}

func (mconn *MConn) readRoutine() {
	log.Println("ReadRoutine: start")
	mconn.readWaitGroup.Add(1)
	defer mconn.readWaitGroup.Done()

	innerRoutineWG := sync.WaitGroup{}

	for {
		// Run async wait for data from server
		ch := make(chan interface{}, 1)
		go func(ch chan<- interface{}) {
			innerRoutineWG.Add(1)
			defer innerRoutineWG.Done()

			data, err := mconn.read()
			if err == io.EOF {
				// Connection closed by server, trying to reconnect
				log.Println("ReadRoutine: innerRoutine: Lost connection (captured EOF). reconnect to ", mconn.addr)
				//err = m.reconnect(m.addr)
				//mconn.notify(RECONNECT)
				mconn.notify(refresh{
					mconn.phonenumber,
					mconn.sessionId,
					nil})
			} else if err != nil {
				log.Println(err.Error())
				re, reerr := regexp.Compile("use of closed network connection")
				if reerr != nil {
					log.Println("ReadRoutine: regexp compilation error, ", reerr)
				}
				if re.MatchString(err.Error()) {
					log.Println("ReadRoutine: TCP connection closed, ", err)
				} else {
					log.Fatalln("ReadRoutine: Unknown error, ", err)
				}
			} else {
				ch <- data
			}
		}(ch)

		select {
		case <-mconn.readInterrupter:
		//case <-m.stopRoutines:
			//TODO: how about to close conn here?
			log.Println("ReadRoutine: wait for inner routine")
			innerRoutineWG.Wait()
			log.Println("ReadRoutine: stop")
			return
		case data := <-ch:
			if data == nil {
				return
			}
			mconn.process(mconn.msgId, mconn.seqNo, data)
		}
	}
}
