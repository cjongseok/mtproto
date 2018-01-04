package mtproto

import (
	"fmt"
	"github.com/cjongseok/slog"
	"math/rand"
	"sync"
	"time"
)

const (
	TIMEOUT_RPC               = 5 * time.Second
	TIMEOUT_INVOKE_WITH_LAYER = 3 * time.Second
	TIMEOUT_UPDATES_GETSTATE  = 3 * time.Second
	TIMEOUT_SESSION_BINDING   = TIMEOUT_INVOKE_WITH_LAYER + TIMEOUT_UPDATES_GETSTATE
	//DELAY_RETRY_OPEN_SESSION  = 1 * time.Second
)

// MConn does not touch sessions.
// binding/unbinding and registration/deregistration of sessions are all handled on MManager.
type MConn struct {
	connId                int32
	session               *MSession
	smonitor              chan MEvent
	interrupter           chan struct{}
	bindWaitGroup         sync.WaitGroup
	listeners             []chan MEvent
	updateCallbacks       []MUpdateCallback
	discardedUpdatesState *TL_updates_state
}

// open, close, and bind should be done by MManager
func newConnection(connListener chan MEvent) *MConn {
	//if connListener == nil {
	//	return nil, fmt.Errorf("nil listener")
	//}
	mconn := new(MConn)
	rand.Seed(time.Now().UnixNano())
	mconn.connId = rand.Int31()
	mconn.smonitor = make(chan MEvent)
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

func (mconn *MConn) bind(session *MSession) error {
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
			updatesDiff, err := mconn.UpdatesGetDifference(
				mconn.discardedUpdatesState.Pts,
				0,
				mconn.discardedUpdatesState.Date,
				mconn.discardedUpdatesState.Qts)
			if err != nil {
				return fmt.Errorf("failed to get update difference")
			}
			unstripped := (*updatesDiff).(TL_updates_difference).Unstrip().(US_updates_difference)
			slog.Logf(mconn, "bind: unstripped diff: %v\n", unstripped)
			mconn.propagate(unstripped)
		}
		mconn.discardedUpdatesState = nil
	} else {
		slog.Logln(mconn, "bind: mconn.discardedUpdatesState is nil")
	}
	return nil
}

func (mconn *MConn) InvokeBlocked(msg TL) (*TL, error) {
	// TODO: timeout the call
	select {
	case x := <-mconn.InvokeNonBlocked(msg):
		if x.err == nil {
			return &x.data, nil
		}
		return nil, x.err
	case <-time.After(TIMEOUT_RPC):
		return nil, fmt.Errorf("RPC Timeout(%f s)", TIMEOUT_RPC.Seconds())
	}
}

func (mconn *MConn) InvokeNonBlocked(msg TL) chan response {
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
func (mconn *MConn) Session() (*MSession, error) {
	// Start race (waiting-for-binding vs. timeout)
	c := make(chan struct{})
	go func() {
		defer close(c)
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
// closing/deregistering session occurs through closeConnection event on MManager
// which is the only caller of this method.
func (mconn *MConn) close() {
	close(mconn.interrupter)
	close(mconn.smonitor)
	mconn.bindWaitGroup.Done()

	// notify the connection is closed
	mconn.notify(connectionClosed{mconn.connId})
}

func (mconn *MConn) AddConnListener(listener chan MEvent) {
	mconn.listeners = append(mconn.listeners, listener)
}

func (mconn *MConn) AddUpdateCallback(callback MUpdateCallback) {
	mconn.updateCallbacks = append(mconn.updateCallbacks, callback)

}

func (mconn *MConn) RemoveConnListener(toremove chan MEvent) error {
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

func (mconn *MConn) RemoveUpdateListener(toremove MUpdateCallback) error {
	for index, registered := range mconn.updateCallbacks {
		if registered == toremove {
			copy(mconn.updateCallbacks[index:], mconn.updateCallbacks[index+1:])
			mconn.updateCallbacks[len(mconn.updateCallbacks)-1] = nil
			mconn.updateCallbacks = mconn.updateCallbacks[:len(mconn.updateCallbacks)-1]
			return nil
		}
	}
	return fmt.Errorf("MUpdateCallback (%x) doesn't exist", toremove)
}

func (mconn *MConn) notify(e MEvent) {
	for _, listener := range mconn.listeners {
		// TODO: it doesn't work. think of another solutino to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (mconn *MConn) propagate(u MUpdate) {
	for _, callback := range mconn.updateCallbacks {
		go func() { callback.OnUpdate(u) }()
	}
}

func (mconn *MConn) monitorSession() {
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
					slog.Logf(mconn, "session will be discarded%d\n", mconn.session.sessionId)
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
