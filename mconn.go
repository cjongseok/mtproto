package mtproto

import (
	"time"
	"math/rand"
	"sync"
	"fmt"
)

const (
	TIMEOUT_SESSION_BINDING = 3 * time.Second
)

// MConn does not touch sessions.
// binding/unbinding and registration/deregistration of sessions are all handled on MManager.
type MConn struct {
	connId			int32
	session 		*MSession
	smonitor		chan MEvent
	interrupter		chan struct{}
	bindWaitGroup	sync.WaitGroup
	listeners		[]chan MEvent
}

// open, close, and bind should be done by MManager
func newConnection(connListener chan MEvent) (*MConn, error) {
	if connListener == nil {
		return nil, fmt.Errorf("nil listener")
	}
	mconn := new(MConn)
	rand.Seed(time.Now().UnixNano())
	mconn.connId = rand.Int31()
	mconn.smonitor = make(chan MEvent)
	mconn.interrupter = make(chan struct{})
	mconn.AddConnListener(connListener)
	mconn.AddConnListener(mconn.smonitor)
	mconn.bindWaitGroup = sync.WaitGroup{}
	defer mconn.bindWaitGroup.Add(1)		// wait for session binding ...

	go mconn.monitorSession()

	mconn.notify(ConnectionOpened{mconn})
	return mconn, nil
}

func (mconn *MConn) bind(session *MSession) error {
	if session == nil {
		return fmt.Errorf("nil ssession")
	}
	defer mconn.bindWaitGroup.Done()	// stop waiting for new session
	session.AddSessionListener(mconn.smonitor)
	session.connId = mconn.connId
	mconn.session = session
	mconn.notify(sessionBound{mconn})

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
	session, err := mconn.Session()
	if err != nil {
		resp <- response{nil, err}
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
	case <- c:
		return mconn.session, nil
	case <- time.After(TIMEOUT_SESSION_BINDING):
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

func (mconn *MConn) notify(e MEvent) {
	for _, listener := range mconn.listeners {
		// TODO: it doesn't work. think of another solutino to handle a deadlock on channel
		//go func(){listener <- e}()
		listener <- e
	}
}

func (mconn *MConn) monitorSession() {
	//log.Printf("[mconn %d] start\n", mconn.connId)
	logln(mconn, "start")
	for {
		select {
		case <-mconn.interrupter:
			//log.Printf("[mconn %d] stop\n", mconn.connId)
			logf(mconn, "stop")
			return
		case e := <- mconn.smonitor:
			switch e.(type) {
			// Session Events
			case newsession:	// triggered only on reconnect (either renewSession or refreshSession)
			case loadsession:	// triggered only on reconnect (either renewSession or refreshSession)
			case SessionEstablished: // triggered only on reconnect (either renewSession or refreshSession)
			case discardSession: // triggered only on reconnect (either renewSession or refreshSession)
			go func() {
				// Unbind the session until the connection has new session
				//log.Printf("[mconn %d] discard session(%d)\n", mconn.connId, mconn.session.sessionId)
				logf(mconn, "the session will be discarded%d\n", mconn.session.sessionId)
				e := e.(discardSession)
				mconn.bindWaitGroup.Add(1)
				// notify inside selection needs non-blocking handlers
				mconn.notify(sessionUnbound{mconn, e.sessionId})
			}()
			case SessionDiscarded: // triggered only on reconnect (either renewSession or refreshSession)
			case renewSession:
			case refreshSession:

			// Connection Events
			case ConnectionOpened:
				go func() {
					//log.Printf("[mconn %d] opened. wait for a session binding ...\n", mconn.connId)
					logln(mconn, "opened. wait for a session binding ...")
				}()
			case sessionBound:
				go func() {
					//log.Printf("[mconn %d] bound to session(%d)\n", mconn.connId, mconn.session.sessionId)
					logf(mconn, "bound to session %d\n", mconn.session.sessionId)
				}()
			case sessionUnbound:
				go func() {
					//log.Printf("[mconn %d] unbound to session(%d)\n", mconn.connId, mconn.session.sessionId)
					logf(mconn, "unbound to session %d\n", mconn.session.sessionId)
				}()
			case closeConnection:
				go func() {
					//log.Printf("[mconn %d] this connection will close\n")
					logln(mconn, "this connection will close")
				}()
			case connectionClosed:
				go func() {
					//log.Printf("[mconn %d] closed\n")
					logln(mconn, "closed")
				}()
			default:
			}
		}
	}

}