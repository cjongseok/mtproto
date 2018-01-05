package mtproto

import (
	"encoding/json"
	"fmt"
	"github.com/cjongseok/slog"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	DEBUG_LEVEL_NETWORK         = 0x01
	DEBUG_LEVEL_NETWORK_DETAILS = 0x02
	DEBUG_LEVEL_DECODE          = 0x04
	DEBUG_LEVEL_DECODE_DETAILS  = 0x08
)

var (
	__debug = 0
)

type MManager struct {
	managerId               int32
	appConfig               Configuration
	conns                   map[int32]*MConn
	sessions                map[int64]*MSession
	stuckSessions           map[int64]int32
	eventq                  chan MEvent
	refreshSessionThrotttle map[int64]int

	manageInterrupter chan struct{}
	manageWaitGroup   sync.WaitGroup
}

func NewManager(appConfig Configuration) (*MManager, error) {
	var err error

	err = appConfig.Check()
	if err != nil {
		return nil, err
	}

	mm := new(MManager)
	rand.Seed(time.Now().UnixNano())
	mm.managerId = rand.Int31()
	mm.appConfig = appConfig
	//TODO: set proper buf size to channels
	mm.conns = make(map[int32]*MConn)
	mm.sessions = make(map[int64]*MSession)
	mm.stuckSessions = make(map[int64]int32)
	mm.eventq = make(chan MEvent)
	mm.refreshSessionThrotttle = make(map[int64]int)
	mm.manageInterrupter = make(chan struct{})
	mm.manageWaitGroup = sync.WaitGroup{}

	go mm.manageRoutine()

	return mm, nil
}

func (mm *MManager) Finish() {
	// close all connections
	for id, _ := range mm.conns {
		mm.eventq <- closeConnection{id, nil}
	}

	// Send stop signal to manage routine
	close(mm.manageInterrupter)

	// Wait for event routines + manage routine
	mm.manageWaitGroup.Wait()
}

func (mm *MManager) IsAuthenticated(phonenumber string) bool {
	sessionfile := sessionFilePath(mm.appConfig.SessionHome, phonenumber)
	_, err := os.Stat(sessionfile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (mm *MManager) LoadAuthentication(phonenumber string, preferredAddr string) (*MConn, error) {
	// req connect
	respCh := make(chan sessionResponse)
	mm.eventq <- loadsession{0, phonenumber, preferredAddr, respCh}

	// Wait for connection built
	resp := <-respCh
	if resp.err != nil {
		return nil, resp.err
	}

	// Check user authentication by user info
	mconn := mm.conns[resp.connId]
	//state, err := mconn.UpdatesGetState()
	//if err != nil {
	//	return nil, err
	//}

	userFull, err := mconn.UsersGetFullUsers(TL_inputUserSelf{})
	if err != nil {
		// Need to authenticate
		return nil, err
	}

	// Already authenticated
	user := userFull.User.(TL_user)
	session, err := mconn.Session()
	if err != nil {
		return mconn, err
	}
	session.user = &user
	slog.Logln(mm, "Auth as ", user)
	return mm.conns[resp.connId], nil
}

func (mm *MManager) NewAuthentication(phonenumber string, addr string, useIPv6 bool) (*MConn, *TL_auth_sentCode, error) {
	// req connect
	respCh := make(chan sessionResponse)
	mm.eventq <- newsession{0, phonenumber, addr, useIPv6, respCh}

	// Wait for connection
	resp := <-respCh
	if resp.err != nil {
		return nil, nil, resp.err
	}

	// sendAuthCode
	mconn := mm.conns[resp.connId]
	mconn, sentCode, err := mm.authSendCode(mconn, phonenumber)
	if err != nil {
		return nil, nil, err
	}

	return mconn, sentCode, nil
}

func (mm *MManager) manageRoutine() {
	slog.Logln(mm, "start")
	mm.manageWaitGroup.Add(1)
	defer mm.manageWaitGroup.Done()

	for {
		select {
		case <-mm.manageInterrupter:
			// Default interrupt is STOP
			slog.Logln(mm, "stop")
			return

		case e := <-mm.eventq:
			// Delegate event handlings to go routines
			switch e.(type) {
			// Session Event Handlers
			// In normal case, three resp events,
			// SessionEstablished, ConnectionOpened, sessionBound,
			// are generated and propagated.
			case newsession:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(newsession)
					slog.Logln(mm, "newsession to ", e.addr)
					session, err := newSession(e.phonenumber, e.addr, e.useIPv6, mm.appConfig, mm.eventq)
					if err != nil {
						//log.Fatalln("ManageRoutine: Connect Failure", err)
						//slog.Fatalln(mm, "connect failure: ", err)
						slog.Logln(mm, "connect failure:", err)
						//TODO: need to handle nil resp channel?
						e.resp <- sessionResponse{0, nil, err}
					} else {
						// Bind the session with mconn and mmanager
						mm.sessions[session.sessionId] = session // Immediate registration
						var mconn *MConn
						if e.connId != 0 {
							mconn = mm.conns[e.connId]
						} else {
							// Create new connection, if not exist
							mconn = newConnection(mm.eventq)
							if err != nil {
								e.resp <- sessionResponse{0, nil, err}
								return
							}
							mm.conns[mconn.connId] = mconn // Immediate registration
						}
						mconn.bind(session)
						//TODO: need to handle nil resp channel?
						e.resp <- sessionResponse{mconn.connId, session, nil}
					}
				}()

				// In normal case, three resp events,
				// SessionEstablished, ConnectionOpened, sessionBound,
				// are generated and propagated.
			case loadsession:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(loadsession)
					slog.Logln(mm, "loadsession of ", e.phonenumber)
					session, err := loadSession(e.phonenumber, e.preferredAddr, mm.appConfig, mm.eventq)
					if err != nil {
						//log.Fatalln("ManageRoutine: Connect Failure", err)
						//slog.Fatalln(mm, "connect failure", err)
						slog.Logln(mm, "connect failure:", err)
						switch err.(type) {
						case handshakingFailure:
							mm.stuckSessions[session.sessionId] = e.connId // register the stuck session
							// usually TCP resets causes stuck sessions, and the sessions are refreshed in the cases.
							// Sometimes TCP t/o makes stuck sessions, and the sessions are refreshed as well,
							// however it takes too long to be identified.
							// So trigger the refresh session by closing the TCP connection
              //mm.eventq <- refreshSession{session.sessionId, session.phonenumber, nil}
              session.close()
						}
						//TODO: separate the handshaking error into two cases and trigger refreshSession on tcp dialing
						// failure
						e.resp <- sessionResponse{0, nil, err}
					} else {
						// Bind the session with mconn and mmanager
						mm.sessions[session.sessionId] = session // Immediate registration
						var mconn *MConn
						if e.connId != 0 {
							mconn = mm.conns[e.connId]
						} else {
							//mconn, err = newConnection(mm.eventq)
							//if err != nil {
							//	e.resp <- sessionResponse{0, nil, err}
							//	return
							//}
							mconn = newConnection(mm.eventq)
							mm.conns[mconn.connId] = mconn // Immediate registration
						}
						mconn.bind(session)
						//TODO: need to handle nil resp channel?
						e.resp <- sessionResponse{mconn.connId, session, nil}
					}
				}()

			case SessionEstablished:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(SessionEstablished)
					slog.Logf(mm, "session established %d\n", e.session.sessionId)
				}()

				// In normal case, an event,
				// SessionDiscarded,
				// is generated and propagated.
			case discardSession:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(discardSession)
					slog.Logln(mm, "discard session ", e.sessionId)
					session := mm.sessions[e.sessionId]
					session.close()

					// Immediate assignment of discarded session's updates state
					// The assignment on handling SessionDiscarded event is sometimes slower than new sessionBound
					// event, so that it results in either nil discardedUpdateState or a lot of duplicated updates.
					marshaled, err := json.Marshal(session.updatesState)
					if err == nil {
						slog.Logf(mm, "session is discarded. keep its updates state, (json): %s\n", marshaled)
					} else {
						slog.Logf(mm, "session is discarded. keep its updates state, %v\n", session.updatesState)
					}
					mconn := mm.conns[e.connId]
					mconn.discardedUpdatesState = new(TL_updates_state)
					*mconn.discardedUpdatesState = *session.updatesState
					e.resp <- sessionResponse{e.connId, session, nil}
				}()

			case SessionDiscarded:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(SessionDiscarded)
					slog.Logln(mm, "session discarded ", e.discardedSessionId)
					delete(mm.sessions, e.discardedSessionId) // Late deregistration
				}()

				// In normal case, five events,
				// discardSesseion, (SessionDiscarded), newsession, (SessionEstablished, ConnectionOpened, sessionBound),
				// are generated and propagated.
			case renewSession:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(renewSession)
					slog.Logln(mm, "renewSession to ", e.addr)
					connId := mm.sessions[e.sessionId].connId

					// Req discardSession
					disconnectRespCh := make(chan sessionResponse)
					//mm.eventq <- discardSession{e.SessionId(), disconnectRespCh}
					mm.sessions[e.sessionId].notify(discardSession{connId, e.sessionId, disconnectRespCh})

					// Wait for disconnection
					disconnectResp := <-disconnectRespCh
					if disconnectResp.err != nil {
						slog.Logf(mm, "renewSession failure: cannot discardSession %d. %v\n", e.sessionId, disconnectResp.err)
						e.resp <- sessionResponse{0, nil, fmt.Errorf("cannot discardSession %d. %v", e.sessionId, disconnectResp.err)}
						return
					}

					// Req newsession
					slog.Logln(mm, "renewRoutine: req newsession")
					connectRespCh := make(chan sessionResponse)
					mm.eventq <- newsession{connId, e.phonenumber, e.addr, e.useIPv6, connectRespCh}
					connectResp := <-connectRespCh
					if connectResp.err != nil {
						slog.Logf(mm, "renewSession failure: cannot connect to %s. %v\n", e.addr, connectResp.err)
						e.resp <- sessionResponse{0, nil, fmt.Errorf("cannot connect to %s. %v", e.addr, connectResp.err)}
						return
					}
					//TODO: need to handle nil resp channel?
					e.resp <- sessionResponse{connectResp.connId, connectResp.session, nil}
					//TODO: figure out missed updates
					slog.Logln(mm, "renewSession done")
				}()

				// In normal case, five events,
				// discardSesseion, (SessionDiscarded), newsession, (SessionEstablished, ConnectionOpened, sessionBound),
				// are generated and propagated.
			case refreshSession:
				// throttle the refreshSession
				if mm.refreshSessionThrotttle[e.(refreshSession).sessionId] > 0 {
					continue
				}
				mm.refreshSessionThrotttle[e.(refreshSession).sessionId] = 1

				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(refreshSession)
					slog.Logln(mm, "refreshSession ", e.sessionId)
					//TODO: alternate the spin lock
					// Wait for session registration and binding for graceful refreshing
					var connId int32
					spinLock := true
					skipDiscardSession := false
					if mm.sessions[e.sessionId] != nil {
						connId = mm.sessions[e.sessionId].connId
						spinLock = false
					}
					for spinLock {
						select {
						// sleep timer
						case <-time.After(1 * time.Second):
							if mm.sessions[e.sessionId] != nil {
								if mm.sessions[e.sessionId].connId != 0 {
									spinLock = false
									connId = mm.sessions[e.sessionId].connId
									slog.Logln(mm, "spinlocked. session(%d) is bound. Release the lock now.", e.sessionId)
								} else {
									slog.Logf(mm, "spinlocked. wait for the session(%d) binding.\n", e.sessionId)
								}
							} else if mm.stuckSessions[e.sessionId] != 0 {
								spinLock = false
								skipDiscardSession = true
								connId = mm.stuckSessions[e.sessionId]
								delete(mm.stuckSessions, e.sessionId)
								slog.Logf(mm, "spinlocked. Session(%d) is stuck on either invokeWithLayer or "+
									"updatesGetState. Release the lock now and skip discardSession.\n", e.sessionId)
							} else {
								slog.Logf(mm, "spinlocked. Session(%d) is waiting for a response from either "+
									"invokeWithLayer or updatesGetState.\n", e.sessionId)
							}
						}
					}

					if !skipDiscardSession {
						// Req discardSession
						disconnectRespCh := make(chan sessionResponse)
						mm.sessions[e.sessionId].notify(discardSession{connId, e.sessionId, disconnectRespCh})

						// Wait for disconnected event
						disconnectResp := <-disconnectRespCh
						if disconnectResp.err != nil {
							slog.Logf(mm, "refreshSession failure: cannot discardSession %d. %v\n", e.sessionId, disconnectResp.err)
							return
						}
					}

					// Req loadsession
					connectRespCh := make(chan sessionResponse)
					var connectResp sessionResponse
					//for {
					slog.Logln(mm, "req loadsession")
					mm.eventq <- loadsession{connId, e.phonenumber, "", connectRespCh}
					connectResp = <-connectRespCh
					if connectResp.err != nil {
						//switch connectResp.err.(type) {
						//case handshakingFailure:
						//  slog.Logf(mm, "retry loadsession after %f seconds: %s", DELAY_RETRY_OPEN_SESSION.Seconds(), connectResp.err)
						//  time.Sleep(DELAY_RETRY_OPEN_SESSION)
						//  continue
						//default:
						slog.Logln(mm, "loadsession failure on refreshSession: ", connectResp.err)
						//slog.Logln(mm ,"retry loadsession")
						e.resp <- sessionResponse{0, nil, connectResp.err}
						//}
					} else {
						//TODO: need to handle nil resp channel?
						e.resp <- sessionResponse{connectResp.connId, connectResp.session, nil}
					}
					//break
					//}
					//TODO: figure out missed updates
					slog.Logln(mm, "refreshSession done. Release the throttle")
					mm.refreshSessionThrotttle[e.sessionId] = 0
				}()

				// Connection Event Handlers
			case ConnectionOpened:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(ConnectionOpened)
					slog.Logln(mm, "connectionOpened ", e.mconn.connId)
				}()

			case sessionBound:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(sessionBound)
					connId := e.mconn.connId
					sessionId := e.mconn.session.sessionId
					slog.Logf(mm, "sessionBound: session %d is bound to mconn %d\n", sessionId, connId)
				}()
			case sessionUnbound:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(sessionUnbound)
					slog.Logf(mm, "sessionUnbound: session %d is unbound from mconn %d\n", e.unboundSessionId, e.mconn.connId)
				}()
			case closeConnection:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(closeConnection)
					slog.Logln(mm, "closeConnection ", e.connId)

					// close, unbound, and deregister session
					mconn := mm.conns[e.connId]
					session, err := mconn.Session()
					if err != nil {
						e.resp <- err
						return
					}
					discardSessionRespCh := make(chan sessionResponse)
					//mm.eventq <- discardSession{closeE.connId, session.sessionId, discardSessionRespCh}
					mconn.notify(discardSession{e.connId, session.sessionId, discardSessionRespCh})

					// close and deregister connection
					discardSessionResp := <-discardSessionRespCh
					if discardSessionResp.err == nil {
						mconn.close()
						e.resp <- nil
						return
					}
					slog.Logln(mm, "closeConnection failure: cannot discard its session ", session.sessionId)
					e.resp <- fmt.Errorf("Failed to discard its session %d", session.sessionId)
				}()
			case connectionClosed:
				go func() {
					mm.manageWaitGroup.Add(1)
					defer mm.manageWaitGroup.Done()
					e := e.(connectionClosed)
					slog.Logln(mm, "connectionClosed ", e.closedConnId)
					delete(mm.conns, e.closedConnId) // Late deregistration
				}()
			case updateReceived:
			default:
			}
		}
	}
	slog.Logln(mm, "done")
}
