package mtproto

import (
	"encoding/json"
	"fmt"
	"github.com/cjongseok/slog"
	"math/rand"
	"sync"
	"time"
)

const (
	DEBUG_LEVEL_NETWORK         = 0x01
	DEBUG_LEVEL_NETWORK_DETAILS = 0x02
	DEBUG_LEVEL_DECODE          = 0x04
	DEBUG_LEVEL_DECODE_DETAILS  = 0x08
	DEBUG_LEVEL_ENCODE_DETAILS  = 0x10
)

var (
	__debug = 0
)

type Manager struct {
	managerId     int32
	appConfig     Configuration
	conns         map[int32]*Conn
	sessions      map[int64]*Session
	stuckSessions map[int64]int32
	eventq        chan Event
	//refreshSessionThrottle map[int64]int
	//queueSend chan packetToSend

	manageInterrupter chan struct{}
	manageWaitGroup   sync.WaitGroup
}

func NewManager(appConfig Configuration) (*Manager, error) {
	var err error

	err = appConfig.Check()
	if err != nil {
		return nil, err
	}

	mm := new(Manager)
	rand.Seed(time.Now().UnixNano())
	mm.managerId = rand.Int31()
	mm.appConfig = appConfig
	//TODO: set proper buf size to channels
	mm.conns = make(map[int32]*Conn)
	mm.sessions = make(map[int64]*Session)
	mm.stuckSessions = make(map[int64]int32)
	mm.eventq = make(chan Event)
	//mm.refreshSessionThrottle = make(map[int64]int)
	//mm.queueSend = make(chan packetToSend, 64)
	mm.manageInterrupter = make(chan struct{})
	mm.manageWaitGroup = sync.WaitGroup{}

	go mm.manageRoutine()

	return mm, nil
}

func (mm *Manager) Finish() {
	// close all connections
	for id, _ := range mm.conns {
		mm.eventq <- closeConnection{id, nil}
	}

	// Send stop signal to manage routine
	close(mm.manageInterrupter)

	// Wait for event routines + manage routine
	mm.manageWaitGroup.Wait()
}

//func (mm *Manager) IsAuthenticated(phonenumber string) bool {
//	sessionfile := sessionFilePath(mm.appConfig.SessionHome, phonenumber)
//	_, err := os.Stat(sessionfile)
//	if os.IsNotExist(err) {
//		return false
//	}
//	return true
//}

func (mm *Manager) LoadAuthentication(phonenumber string) (*Conn, error) {
	// req connect
	respCh := make(chan sessionResponse, 1)
	mm.eventq <- loadsession{0, phonenumber, respCh}

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

	// Request full user
	inputUser := &TypeInputUser{&TypeInputUser_InputUserSelf{&PredInputUserSelf{}}}
	var userFull *TypeUserFull
	x := <-mconn.InvokeNonBlocked(&ReqUsersGetFullUser{inputUser})
	if x.err != nil {
		return nil, x.err
	}

	switch casted := x.data.(type) {
	case *PredUserFull:
		userFull = &TypeUserFull{casted}
	default:
		return nil, fmt.Errorf("no full user: %T: %v", x, x)
	}

	// Already authenticated
	typeUser := userFull.GetValue().GetUser()
	session, err := mconn.Session()
	if err != nil {
		return mconn, err
	}
	if typeUser.GetUser() != nil {
		user := typeUser.GetUser()
		session.user = user
		slog.Logln(mm, "Auth as ", user)
	} else if typeUser.GetUserEmpty() != nil {
		session.user = &PredUser{}
		slog.Logln(mm, "Authenticated, but failed to get user")
	}
	return mm.conns[resp.connId], nil
}

func (mm *Manager) NewAuthentication(phonenumber string, addr string, useIPv6 bool) (*Conn, *TypeAuthSentCode, error) {
	// req connect
	respCh := make(chan sessionResponse, 1)
	mm.eventq <- newsession{0, phonenumber, addr, useIPv6, respCh}

	// Wait for connection
	resp := <-respCh
	if resp.err != nil {
		return nil, nil, resp.err
	}

	// sendAuthCode
	mconn := mm.conns[resp.connId]
	for {
		//sentCode, err := mconn.authSendCode(phonenumber)
		session, err := mconn.Session()
		if err != nil {
			return nil, nil, err
		}

		// request to send code
		data, err := mconn.InvokeBlocked(&ReqAuthSendCode{
			//Allow_flashcall: false,
			Flags:         0x00000001,
			PhoneNumber:   phonenumber,
			CurrentNumber: &TypeBool{&TypeBool_BoolTrue{&PredBoolTrue{}}},
			ApiId:         session.appConfig.Id,
			ApiHash:       session.appConfig.Hash,
		})
		switch x := data.(type) {
		case *PredAuthSentCode:
			return mconn, &TypeAuthSentCode{x}, nil
			//default:
			//	return nil, nil, fmt.Errorf("authSendCode: Got: %T", data)
		}

		// retry the send code request to another server
		if err != nil {
			rpcError, ok := err.(TL_rpc_error)
			if !ok {
				return nil, nil, err
			}
			if rpcError.error_code != errorSeeOther {
				return nil, nil, err
			}

			var newdc int32
			n, _ := fmt.Sscanf(rpcError.error_message, "PHONE_MIGRATE_%d", &newdc)
			if n != 1 {
				n, _ = fmt.Sscanf(rpcError.error_message, "NETWORK_MIGRATE_%d", &newdc)
			}
			if n != 1 {
				return nil, nil, err
			} else {
				// Reconnect to the new datacenter
				session, err := mconn.Session()
				if err != nil {
					return nil, nil, err
				}
				respch := make(chan sessionResponse, 1)

				//TODO: Check if renewSession event works with mconn.notify()
				mconn.notify(renewSession{
					session.sessionId,
					phonenumber,
					session.dclist[newdc],
					session.useIPv6,
					respch,
				})

				// Wait for binding with new session
				resp := <-respch
				if resp.err != nil {
					return nil, nil, resp.err
				}
			}
		}
	}
}

func (mm *Manager) manageRoutine() {
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
					session, err := newSession(e.phonenumber, e.addr, e.useIPv6, mm.appConfig /*mm.queueSend,*/, mm.eventq)
					var resp sessionResponse
					if err != nil {
						slog.Logln(mm, "connect failure:", err)
						//TODO: need to handle nil resp channel?
						//e.resp <- sessionResponse{0, nil, err}
						resp = sessionResponse{0, nil, err}
					} else {
						// Bind the session with mconn and mmanager
						mm.sessions[session.sessionId] = session // Immediate registration
						var mconn *Conn
						if e.connId != 0 {
							mconn = mm.conns[e.connId]
						} else {
							// Create new connection, if not exist
							mconn = newConnection(mm.eventq)
							if err != nil {
								//e.resp <- sessionResponse{0, nil, err}
								if e.resp != nil {
									e.resp <- sessionResponse{0, nil, err}
								}
								return
							}
							mm.conns[mconn.connId] = mconn // Immediate registration
						}
						mconn.bind(session)
						//TODO: need to handle nil resp channel?
						//e.resp <- sessionResponse{mconn.connId, session, nil}
						resp = sessionResponse{mconn.connId, session, nil}
					}
					if e.resp != nil {
						e.resp <- resp
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
					session, err := loadSession(e.phonenumber, mm.appConfig /*mm.queueSend,*/, mm.eventq)
					var resp sessionResponse
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
						resp = sessionResponse{0, session, err}
					} else {
						// Bind the session with mconn and mmanager
						mm.sessions[session.sessionId] = session // Immediate registration
						var mconn *Conn
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
						resp = sessionResponse{mconn.connId, session, nil}
					}
					if e.resp != nil {
						e.resp <- resp
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
					if e.connId != 0 {
						mconn := mm.conns[e.connId]
						mconn.discardedUpdatesState = &PredUpdatesState{}
						*mconn.discardedUpdatesState = *session.updatesState
					}
					if e.resp != nil {
						e.resp <- sessionResponse{e.connId, session, nil}
					}
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
					disconnectRespCh := make(chan sessionResponse, 1)
					//mm.eventq <- discardSession{e.SessionId(), disconnectRespCh}
					mm.sessions[e.sessionId].notify(discardSession{connId, e.sessionId, disconnectRespCh})

					// Wait for disconnection
					disconnectResp := <-disconnectRespCh
					if disconnectResp.err != nil {
						slog.Logf(mm, "renewSession failure: cannot discardSession %d. %v\n", e.sessionId, disconnectResp.err)
						if e.resp != nil {
							e.resp <- sessionResponse{0, nil, fmt.Errorf("cannot discardSession %d. %v", e.sessionId, disconnectResp.err)}
						}
						return
					}

					// Req newsession
					slog.Logln(mm, "renewRoutine: req newsession")
					connectRespCh := make(chan sessionResponse, 1)
					mm.eventq <- newsession{connId, e.phonenumber, e.addr, e.useIPv6, connectRespCh}
					connectResp := <-connectRespCh
					if connectResp.err != nil {
						slog.Logf(mm, "renewSession failure: cannot connect to %s. %v\n", e.addr, connectResp.err)
						if e.resp != nil {
							e.resp <- sessionResponse{0, nil, fmt.Errorf("cannot connect to %s. %v", e.addr, connectResp.err)}
						}
						return
					}
					slog.Logln(mm, "renewSession done")
					//TODO: need to handle nil resp channel?
					if e.resp != nil {
						e.resp <- sessionResponse{connectResp.connId, connectResp.session, nil}
					}
					//TODO: figure out missed updates
				}()

				// In normal case, five events,
				// discardSesseion, (SessionDiscarded), newsession, (SessionEstablished, ConnectionOpened, sessionBound),
				// are generated and propagated.
			case refreshSession:
				// throttle the refreshSession
				//if mm.refreshSessionThrottle[e.(refreshSession).sessionId] > 0 {
				//	continue
				//}
				//mm.refreshSessionThrottle[e.(refreshSession).sessionId] = 1

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
								// session is registered
								if mm.sessions[e.sessionId].connId != 0 {
									// session is bound to a connection
									spinLock = false
									connId = mm.sessions[e.sessionId].connId
									slog.Logln(mm, "spinlocked. session(%d) is bound. Release the lock now.", e.sessionId)
								} else {
									// session is not bound to a connection yet
									slog.Logf(mm, "spinlocked. wait for the session(%d) binding.\n", e.sessionId)
								}
							} else if stuckSessionConnId, ok := mm.stuckSessions[e.sessionId]; ok {
								// session is not registered yet,
								// even the session would not be registered forever,
								// because either invokeWithLayer or updatesGetState does not respond.
								spinLock = false
								skipDiscardSession = true
								connId = stuckSessionConnId
								delete(mm.stuckSessions, e.sessionId)
								slog.Logf(mm, "spinlocked. Session(%d) is stuck on either invokeWithLayer or "+
									"updatesGetState. Release the lock now and skip discardSession.\n", e.sessionId)
							} else {
								// session is not registered yet. wait for the registration.
								slog.Logf(mm, "spinlocked. Session(%d) is waiting for a response from either "+
									"invokeWithLayer or updatesGetState.\n", e.sessionId)
							}
						}
					}

					if !skipDiscardSession {
						// Req discardSession
						disconnectRespCh := make(chan sessionResponse, 1)
						mm.sessions[e.sessionId].notify(discardSession{connId, e.sessionId, disconnectRespCh})

						// Wait for disconnected event
						disconnectResp := <-disconnectRespCh
						if disconnectResp.err != nil {
							slog.Logf(mm, "refreshSession failure: cannot discardSession %d. %v\n", e.sessionId, disconnectResp.err)
							return
						}
					}

					// Req loadsession
					connectRespCh := make(chan sessionResponse, 1)
					var connectResp sessionResponse
					slog.Logln(mm, "req loadsession")
					mm.eventq <- loadsession{connId, "", connectRespCh}
					connectResp = <-connectRespCh
					var sessionResp sessionResponse
					if connectResp.err != nil {
						slog.Logln(mm, "loadsession failure on refreshSession: ", connectResp.err)
						sessionResp = sessionResponse{0, nil, connectResp.err}
					} else {
						//TODO: need to handle nil resp channel?
						sessionResp = sessionResponse{connectResp.connId, connectResp.session, nil}
					}
					if sessionResp.err != nil && e.policy == untilSuccess {
						slog.Logln(mm, "retry refreshSession")
						mm.eventq <- refreshSession{
							sessionResp.session.sessionId,
							e.phonenumber,
							untilSuccess,
							make(chan sessionResponse),
						}
					} else {
						slog.Logln(mm, "refreshSession is done.")
						//mm.refreshSessionThrottle[e.sessionId] = 0
					}
					if e.resp != nil {
						e.resp <- sessionResp
					}
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
						if e.resp != nil {
							e.resp <- err
						}
						return
					}
					discardSessionRespCh := make(chan sessionResponse, 1)
					//mm.eventq <- discardSession{closeE.connId, session.sessionId, discardSessionRespCh}
					mconn.notify(discardSession{e.connId, session.sessionId, discardSessionRespCh})

					// close and deregister connection
					discardSessionResp := <-discardSessionRespCh
					if discardSessionResp.err == nil {
						mconn.close()
						if e.resp != nil {
							e.resp <- nil
						}
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

func (x *Manager) LogPrefix() string {
	return fmt.Sprintf("[MM %d]", x.managerId)
}
