package mtproto

import (
	"fmt"
	"github.com/cjongseok/slog"
	"math/rand"
	"sync"
	"time"
)

const (
	//TODO: elastic timeout
	TIMEOUT_RPC               = 5 * time.Second
	TIMEOUT_INVOKE_WITH_LAYER = 10 * time.Second
	TIMEOUT_UPDATES_GETSTATE  = 7 * time.Second
	TIMEOUT_SESSION_BINDING   = TIMEOUT_INVOKE_WITH_LAYER + TIMEOUT_UPDATES_GETSTATE
	//DELAY_RETRY_OPEN_SESSION  = 1 * time.Second
)

// Conn is mtproto connection. Conn guarantees it is always wired-up with Telegram server, although
// Session can expire anytime without notice.
type Conn struct {
	connID                int32
	session               Session
	smonitor              chan Event
	interrupter           chan struct{}
	bindSession           chan Session
	sessionReqs           []chan Session
	sessionReqsMux        *sync.Mutex
	listeners             []chan Event
	updateCallbacks       []UpdateCallback
	callbackMux           *sync.Mutex
	discardedUpdatesState *PredUpdatesState
}

// newConnection creates a Session instance. Other session actions such as 'open', 'close',
// 'bind (to a conn)', 'unbind (from a conn)', 'register (to the manager)', and 'de-register (from
// the manager)' are controlled by Manager.
func newConnection(connListener chan Event) *Conn {
	mconn := new(Conn)
	rand.Seed(time.Now().UnixNano())
	mconn.connID = rand.Int31()
	mconn.smonitor = make(chan Event)
	mconn.interrupter = make(chan struct{})
	mconn.AddConnListener(connListener)
	mconn.AddConnListener(mconn.smonitor)
	mconn.bindSession = make(chan Session)
	mconn.sessionReqsMux = &sync.Mutex{}
	mconn.callbackMux = &sync.Mutex{}

	go mconn.monitorSession()

	mconn.notify(ConnectionOpened{mconn, 0})
	//return mconn, nil
	return mconn
}

// bind attaches a Session to the Conn. If the Conn already has a Session, it alternates the old
// one.
func (mconn *Conn) bind(session *Session) error {
	slog.Logln(mconn, "bind session", session.sessionID)
	if session == nil {
		return fmt.Errorf("nil ssession")
	}
	session.AddSessionListener(mconn.smonitor)
	session.connID = mconn.connID
	slog.Logln(mconn, "pass the session to bindSession ch")
	mconn.bindSession <- *session
	slog.Logln(mconn, "sesssin passed")
	mconn.notify(sessionBound{mconn, session.sessionID})

	//TODO: get updates difference on opening session rather than its binding
	// req updates, if exists
	if mconn.discardedUpdatesState != nil {
		ptsDiff := session.updatesState.Pts - mconn.discardedUpdatesState.Pts
		qtsDiff := session.updatesState.Qts - mconn.discardedUpdatesState.Qts
		seqDiff := session.updatesState.Seq - mconn.discardedUpdatesState.Seq
		if ptsDiff > 0 || qtsDiff > 0 || seqDiff > 0 {
			// missed updates exist. Propagate updates to callbacks
			updatesDiff, err := mconn.InvokeBlocked(&ReqUpdatesGetDifference{
				Pts:           mconn.discardedUpdatesState.Pts,
				PtsTotalLimit: 0,
				Date:          mconn.discardedUpdatesState.Date,
				Qts:           mconn.discardedUpdatesState.Qts})
			if err != nil {
				return fmt.Errorf("failed to get update difference")
			}

			switch udiff := updatesDiff.(type) {
			case *PredUpdatesDifferenceEmpty:
				slog.Logln(mconn, "bind: diff: empty")
			case *PredUpdatesDifference:
				slog.Logln(mconn, "bind: diff:", udiff)
				mconn.propagate(udiff)
			case *PredUpdatesDifferenceSlice:
				slog.Logln(mconn, "bind: diff: slice:", udiff)
				mconn.propagate(udiff)
			case *PredUpdatesDifferenceTooLong:
				slog.Logln(mconn, "bind: diff: too long")
			default:
				slog.Logln(mconn, "bind: no diff")
			}
		}
		mconn.discardedUpdatesState = nil
	} else {
		slog.Logln(mconn, "bind: mconn.discardedUpdatesState is nil")
	}
	return nil
}

func (mconn *Conn) InvokeBlocked(msg TL) (interface{}, error) {
	// TODO: timeout the call
	select {
	case x := <-mconn.InvokeNonBlocked(msg):
		if x.err == nil {
			return x.data, nil
		}
		return nil, x.err

	case <-time.After(TIMEOUT_RPC):
		return nil, fmt.Errorf("RPC Timeout(%f s)", TIMEOUT_RPC.Seconds())
	}
}

func (mconn *Conn) InvokeNonBlocked(msg TL) chan response {
	resp := make(chan response, 1)

	// get session
	var session Session
	res := <-mconn.Session()
	switch res.(type) {
	case Session:
		session = res.(Session)
	case error:
		resp <- response{nil, res.(error)}
		return resp
	}

	session.queueSend <- packetToSend{
		msg:  msg,
		resp: resp,
	}
	return resp
}

// Session returns the bound bound of the conn. The direct access to the session (using '.') does
// not guarantee both not-nil and data racing. The returned session can expire any time, so that it
// cannot match with the latest bound session of the conn.
func (mconn *Conn) Session() <-chan interface{} {
	sessionCh := make(chan Session, 1)
	slog.Logln(mconn, "session request", sessionCh)
	if mconn.session.sessionID != 0 {
		go func(send chan<- Session, session Session) {
			send <- session
		}(sessionCh, mconn.session)
	} else {
		mconn.sessionReqsMux.Lock()
		mconn.sessionReqs = append(mconn.sessionReqs, sessionCh)
		mconn.sessionReqsMux.Unlock()
	}

	promise := make(chan interface{})
	go func(raceResult chan interface{}, receiveSession chan Session) {
		select {
		case newSession := <-receiveSession:
			raceResult <- newSession
		case <-time.After(TIMEOUT_SESSION_BINDING):
			raceResult <- fmt.Errorf("session binding timeout (%v)", TIMEOUT_SESSION_BINDING)
		}
	}(promise, sessionCh)
	return promise
}

func (mconn *Conn) MigrateSessionTo(newdc int32) error {
	// get session
	var session Session
	res := <-mconn.Session()
	switch res.(type) {
	case Session:
		session = res.(Session)
	case error:
		return res.(error)
	}

	// reconnect to the new datacenter
	respch := make(chan sessionResponse, 1)
	ipVersion := ipv4
	if isIPv6(session.c.IP) {
		ipVersion = ipv6
	}
	dcOption, err := session.apiDcOption(ipVersion, newdc)
	if err != nil {
		return err
	}
	slog.Logln(mconn, "migrate session to", dcOption)

	//TODO: Check if renewSession event works with mconn.notify()
	mconn.notify(renewSession{
		session.sessionID,
		session.c.Phone,
		session.c.ApiID,
		session.c.ApiHash,
		dcOption.IpAddress,
		int(dcOption.Port),
		respch,
	})

	// Wait for binding the new session
	resp := <-respch
	return resp.err
}

// finish connection's internal resource but bound session.
// closing/deregistering session occurs through closeConnection event on Manager
// which is the only caller of this method.
func (mconn *Conn) close() {
	close(mconn.interrupter)
	close(mconn.smonitor)

	// notify the connection is closed
	mconn.notify(connectionClosed{mconn.connID})
}

func (mconn *Conn) AddConnListener(listener chan Event) {
	mconn.listeners = append(mconn.listeners, listener)
}

func (mconn *Conn) AddUpdateCallback(callback UpdateCallback) {
	mconn.callbackMux.Lock()
	defer mconn.callbackMux.Unlock()
	mconn.updateCallbacks = append(mconn.updateCallbacks, callback)

}

func (mconn *Conn) RemoveConnListener(toremove chan Event) error {
	for index, registered := range mconn.listeners {
		if registered == toremove {
			copy(mconn.listeners[index:], mconn.listeners[index+1:])
			mconn.listeners[len(mconn.listeners)-1] = nil
			mconn.listeners = mconn.listeners[:len(mconn.listeners)-1]
			return nil
		}
	}
	return fmt.Errorf("listener (%v) doesn't exist", toremove)
}

func (mconn *Conn) RemoveUpdateListener(toremove UpdateCallback) error {
	for index, registered := range mconn.updateCallbacks {
		if registered == toremove {
			copy(mconn.updateCallbacks[index:], mconn.updateCallbacks[index+1:])
			mconn.updateCallbacks[len(mconn.updateCallbacks)-1] = nil
			mconn.updateCallbacks = mconn.updateCallbacks[:len(mconn.updateCallbacks)-1]
			return nil
		}
	}
	return fmt.Errorf("UpdateCallback (%x) doesn't exist", toremove)
}

func (mconn *Conn) notify(event Event) {
	mconn.callbackMux.Lock()
	defer mconn.callbackMux.Unlock()
	for _, listener := range mconn.listeners {
		go func(l chan Event, e Event) {
			l <- e
		}(listener, event)
	}
}

func (mconn *Conn) propagate(u Update) {
	for _, callback := range mconn.updateCallbacks {
		go func(cb UpdateCallback) {
			cb.OnUpdate(u)
		}(callback)
	}
}

func (mconn *Conn) monitorSession() {
	slog.Logln(mconn, "start")
	for {
		select {
		case <-mconn.interrupter:
			slog.Logln(mconn, "stop")
			return
		case newSession := <-mconn.bindSession:
			mconn.session = newSession
			go func(session Session) {
				mconn.sessionReqsMux.Lock()
				defer mconn.sessionReqsMux.Unlock()
				for _, req := range mconn.sessionReqs {
					go func(c chan Session, s Session){
						c <- s
					}(req, session)
				}
				mconn.sessionReqs = nil
			}(mconn.session)
		case e := <-mconn.smonitor:
			slog.Logf(mconn, "session event %T(%v)\n", e, e)
			switch e.(type) {
			// Session Events
			case newsession: // never triggered on mconn
			case loadsession: // never triggered on mconn
			case SessionEstablished: // never triggered on mconn
			case discardSession: // triggered only on reconnect (either renewSession or refreshSession)
				// Unbind the session until the connection has new session
				mconn.session = Session{}
				slog.Logf(mconn, "session(%d) will be discarded\n", e.(discardSession).sessionId)
				go func(sid int64) {
					unbound := sessionUnbound{mconn, sid}
					// notify that inside selection needs non-blocking handlers
					mconn.notify(unbound)
				}(e.(discardSession).sessionId)
			case SessionDiscarded: // triggered only on reconnect (either renewSession or refreshSession)
			case renewSession:
			case refreshSession:

			// Connection Events
			case ConnectionOpened:
				slog.Logln(mconn, "opened")
				if e.(ConnectionOpened).sessionID != 0 {
					slog.Logln(mconn, "wait for a session binding ...")
				} else {
					slog.Logln(mconn, "with session,", e.(ConnectionOpened).sessionID)
				}
			case sessionBound:
				slog.Logln(mconn, "bound to session", e.(sessionBound).sessionID)
			case sessionUnbound:
				e := e.(sessionUnbound)
				slog.Logln(mconn, "unbound to session", e.unboundSessionID)
			case closeConnection:
				slog.Logln(mconn, "this connection will close")
			case connectionClosed:
				slog.Logln(mconn, "closed")

				// Update Event
			case updateReceived:
				go func() {
					slog.Logln(mconn, "received an update, ", e.(updateReceived).update)
					mconn.propagate(e.(updateReceived).update)
				}()
			default:
			}
		}
	}
}

func (mconn *Conn) SignIn(phoneNumber, phoneCode, phoneCodeHash string) (*TypeAuthAuthorization, error) {
	if phoneNumber == "" || phoneCode == "" || phoneCodeHash == "" {
		return nil, fmt.Errorf("empty sign-in argument")
	}

	x := <-mconn.InvokeNonBlocked(&ReqAuthSignIn{
		PhoneNumber:   phoneNumber,
		PhoneCodeHash: phoneCodeHash,
		PhoneCode:     phoneCode,
	})
	if x.err != nil {
		return nil, x.err
	}

	auth, ok := x.data.(*PredAuthAuthorization)
	if !ok {
		return nil, fmt.Errorf("RPC: %v", x)
	}

	// get session
	var session Session
	res := <-mconn.Session()
	switch res.(type) {
	case Session:
		session = res.(Session)
	case error:
		return &TypeAuthAuthorization{Value: auth}, res.(error)
	}

	if auth.GetUser().GetUser() != nil {
		session.user = auth.GetUser().GetUser()
		slog.Logln(mconn, "Signed in as ", session.user)
	} else if auth.GetUser().GetUserEmpty() != nil {
		session.user = &PredUser{}
		slog.Logln(mconn, "Signed in with empty user")
	} else {
		session.user = &PredUser{}
		slog.Logln(mconn, "Signed in without user response: neither user nor user empty")
	}
	return &TypeAuthAuthorization{Value: auth}, nil
}

func (mconn *Conn) SignOut() (bool, error) {
	var result bool
	x := <-mconn.InvokeNonBlocked(&ReqAuthLogOut{})
	if x.err != nil {
		return result, x.err
	}

	if tl, ok := x.data.(TL); ok {
		return toBool(tl), nil
	}
	return false, fmt.Errorf("invalid rpc return: %T: %v", x.data, x.data)
}

func (x *Conn) LogPrefix() string {
	return fmt.Sprintf("[mconn %d]", x.connID)
}
