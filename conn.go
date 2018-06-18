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

// Conn does not touch sessions.
// binding/unbinding and registration/deregistration of sessions are all handled on Manager.
type Conn struct {
	connId                int32
	session               *Session
	smonitor              chan Event
	interrupter           chan struct{}
	bindWaitGroup         sync.WaitGroup
	listeners             []chan Event
	updateCallbacks       []UpdateCallback
	discardedUpdatesState *PredUpdatesState
}

// open, close, and bind should be done by Manager
func newConnection(connListener chan Event) *Conn {
	//if connListener == nil {
	//	return nil, fmt.Errorf("nil listener")
	//}
	mconn := new(Conn)
	rand.Seed(time.Now().UnixNano())
	mconn.connId = rand.Int31()
	mconn.smonitor = make(chan Event)
	mconn.interrupter = make(chan struct{})
	mconn.AddConnListener(connListener)
	mconn.AddConnListener(mconn.smonitor)
	mconn.bindWaitGroup = sync.WaitGroup{}
	defer mconn.bindWaitGroup.Add(1) // wait for session binding ...

	go mconn.monitorSession()

	mconn.notify(ConnectionOpened{mconn})
	//return mconn, nil
	return mconn
}

func (mconn *Conn) bind(session *Session) error {
	if session == nil {
		return fmt.Errorf("nil ssession")
	}
	session.AddSessionListener(mconn.smonitor)
	session.connId = mconn.connId
	mconn.session = session
	mconn.bindWaitGroup.Done() // stop waiting for new session. Enable querying
	mconn.notify(sessionBound{mconn})

	//TODO: get updates difference on opening session rather than its binding
	// req updates, if exists
	if mconn.discardedUpdatesState != nil {
		//slog.Logf(mconn, "bind: new session seq:%d, unbound session seq:%d\n", session.updatesState.Seq, mconn.discardedUpdatesState.Seq)
		//seqDiff := mconn.discardedUpdatesState.Seq - session.updatesState.Seq
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
				slog.Logf(mconn, "bind: diff: empty\n")
			case *PredUpdatesDifference:
				slog.Logf(mconn, "bind: diff: %v\n", udiff)
				mconn.propagate(udiff)
			case *PredUpdatesDifferenceSlice:
				slog.Logf(mconn, "bind: diff: slice: %v\n", udiff)
				mconn.propagate(udiff)
			case *PredUpdatesDifferenceTooLong:
				slog.Logf(mconn, "bind: diff: too long\n")
			default:
				slog.Logf(mconn, "bind: no diff\n")
			}
			//unstripped := (*updatesDiff).(*PredUpdatesDifference).Unstrip().(US_updates_difference)
			//udiff := (*updatesDiff).(*PredUpdatesDifference)
			//slog.Logf(mconn, "bind: unstripped diff: %v\n", unstripped)
			//mconn.propagate(unstripped)
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
	session, err := mconn.Session()
	if err != nil {
		resp <- response{nil, err}
		return resp
	}
	session.queueSend <- packetToSend{
		msg:  msg,
		resp: resp,
	}
	return resp
}

// CAVEAT:
// Accessing the session without this method does NOT ensure
// the session is alive.
//TODO: fast session failure is better than slow session failure?
//TODO: Think of better way of handling timeout (rather than returning nil + err?)
func (mconn *Conn) Session() (*Session, error) {
	// Start race (waiting-for-binding vs. timeout)
	c := make(chan struct{})
	go func() {
		defer close(c)
		//slog.Logln(mconn, "mconn:", mconn)
		//slog.Logln(mconn, "bindWaitGroup:", mconn.bindWaitGroup)
		mconn.bindWaitGroup.Wait()
		//TODO: ping to prolong session life? Because session can be aborted
	}()
	select {
	case <-c:
		return mconn.session, nil
	case <-time.After(TIMEOUT_SESSION_BINDING):
		return nil, fmt.Errorf("No Session: session binding timeout")
	}
}

// finish connection's internal resource but bound session.
// closing/deregistering session occurs through closeConnection event on Manager
// which is the only caller of this method.
func (mconn *Conn) close() {
	close(mconn.interrupter)
	close(mconn.smonitor)
	mconn.bindWaitGroup.Done()

	// notify the connection is closed
	mconn.notify(connectionClosed{mconn.connId})
}

func (mconn *Conn) AddConnListener(listener chan Event) {
	mconn.listeners = append(mconn.listeners, listener)
}

func (mconn *Conn) AddUpdateCallback(callback UpdateCallback) {
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
	return fmt.Errorf("Listener (%x) doesn't exist", toremove)
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

func (mconn *Conn) notify(e Event) {
	for _, listener := range mconn.listeners {
		// TODO: it doesn't work. think of another solutino to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (mconn *Conn) propagate(u Update) {
	for _, callback := range mconn.updateCallbacks {
		go func() { callback.OnUpdate(u) }()
	}
}

func (mconn *Conn) monitorSession() {
	slog.Logln(mconn, "start")
	for {
		select {
		case <-mconn.interrupter:
			slog.Logf(mconn, "stop")
			return
		case e := <-mconn.smonitor:
			switch e.(type) {
			// Session Events
			case newsession: // never triggered on mconn
			case loadsession: // never triggered on mconn
			case SessionEstablished: // never triggered on mconn
			case discardSession: // triggered only on reconnect (either renewSession or refreshSession)
				go func() {
					// Unbind the session until the connection has new session
					slog.Logf(mconn, "session(%d) will be discarded\n", mconn.session.sessionId)
					e := e.(discardSession)
					mconn.bindWaitGroup.Add(1)
					unbound := sessionUnbound{mconn, e.sessionId}
					mconn.session = nil
					// notify that inside selection needs non-blocking handlers
					mconn.notify(unbound)
				}()
			case SessionDiscarded: // triggered only on reconnect (either renewSession or refreshSession)
			case renewSession:
			case refreshSession:

			// Connection Events
			case ConnectionOpened:
				go func() {
					slog.Logf(mconn, "opened.")
					if mconn.session == nil {
						slog.Logf(mconn, "wait for a session binding ...\n")
					} else {
						slog.Logf(mconn, "with session, %d\n", mconn.session.sessionId)
					}
				}()
			case sessionBound:
				go func() {
					slog.Logf(mconn, "bound to session %d\n", mconn.session.sessionId)
				}()
			case sessionUnbound:
				go func() {
					e := e.(sessionUnbound)
					slog.Logf(mconn, "unbound to session %d\n", e.unboundSessionId)
				}()
			case closeConnection:
				go func() {
					slog.Logln(mconn, "this connection will close")
				}()
			case connectionClosed:
				go func() {
					slog.Logln(mconn, "closed")
				}()

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

	session, err := mconn.Session()
	if err != nil {
		return &TypeAuthAuthorization{Value: auth}, err
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
	return fmt.Sprintf("[mconn %d]", x.connId)
}
