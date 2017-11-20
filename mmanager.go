package mtproto

import (
	"log"
	"os"
	"sync"
)

type MManager struct {
	appConfig 	Configuration
	conns 		map[int64]*MConn
	eventq 		chan MEvent

	manageInterrupter chan struct{}
	manageWaitGroup   sync.WaitGroup
}

const (
	// Current API Layer Version
	layer = 65

	// API Errors
	errorSeeOther     = 303
	errorBadRequest   = 400
	errorUnauthorized = 401
	errorForbidden    = 403
	errorNotFound     = 404
	errorFlood        = 420
	errorInternal     = 500
)

func NewManager (appConfig Configuration) (*MManager, error) {
	var err error

	err = appConfig.Check()
	if err != nil {
		return nil, err
	}

	cm := new(MManager)
	cm.appConfig = appConfig
	//TODO: set proper buf size to channels
	cm.conns = make(map[int64]*MConn)
	cm.eventq = make(chan MEvent)
	cm.manageInterrupter = make(chan struct{})
	cm.manageWaitGroup = sync.WaitGroup{}

	go cm.manageRoutine()

	return cm, nil
}

func (cm *MManager) Finish() {
	// Send stop signal to manage routine
	close(cm.manageInterrupter)

	// Wait for event routines + manage routine
	cm.manageWaitGroup.Wait()
}

func (cm *MManager) IsAuthenticated(phonenumber string) bool {
	sessionfile := sessionFilePath(cm.appConfig.SessionHome, phonenumber)
	_, err := os.Stat(sessionfile)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (cm *MManager) LoadAuthentication(phonenumber string) (*MConn, error){
	// req connect
	respCh := make(chan connectResponse)
	cm.eventq  <- loadsession{phonenumber, respCh}

	// Wait for connection built
	resp := <- respCh
	if resp.err != nil {
		return nil, resp.err
	}

	// Check user authentication by user info
	mconn := resp.mconn
	userFull, err := mconn.UsersGetFullUsers(TL_inputUserSelf{})
	if err != nil {
		// Need to authenticate
		return nil, err
	}
	// Already authenticated
	user := userFull.User.(TL_user)
	mconn.user = &user
	log.Println("Signed in as ", mconn.user.Username)
	return mconn, nil
}

func (cm *MManager) NewAuthentication(phonenumber string, addr string, useIPv6 bool) (*MConn, *TL_auth_sentCode, error) {
	// req connect
	respCh := make(chan connectResponse)
	cm.eventq <- newsession{phonenumber, addr, useIPv6, respCh}

	// Wait for connection
	resp := <- respCh
	if resp.err != nil {
		return nil, nil, resp.err
	}

	// sendAuthCode
	mconn := resp.mconn
	mconn, sentCode, err := cm.AuthSendCode(mconn, phonenumber)
	if err != nil {
		return nil, nil, err
	}
	return mconn, sentCode, nil
}

func (cm *MManager) manageRoutine() {
	log.Println("ManageRoutine: start")
	cm.manageWaitGroup.Add(1)
	defer cm.manageWaitGroup.Done()

	for {
		select {
		case <-cm.manageInterrupter:
			// Default interrupt is STOP
			log.Println("ManageRoutine: stop")
			return

		case e := <-cm.eventq :
			// Delegate event handlings to go routines
			switch e.(type) {
			case newsession:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					e := e.(newsession)
					log.Println("ManageRoutine: newsession to ", e.addr)
					mconn, err := newSession(e.phonenumber, e.addr, e.useIPv6, cm.appConfig, cm.eventq)
					if err != nil {
						log.Fatalln("ManageRoutine: Connect Failure", err)
						//TODO: need to handle nil resp channel?
						e.resp <- connectResponse{nil, err}
					} else {
						cm.conns[mconn.sessionId] = mconn
						//TODO: need to handle nil resp channel?
						e.resp <- connectResponse{mconn, nil}
					}
				}()

			case loadsession:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					e := e.(loadsession)
					log.Println("ManageRoutine: loadsession of ", e.phonenumber)
					mconn, err := loadSession(e.phonenumber, cm.appConfig, cm.eventq)
					if err != nil {
						log.Fatalln("ManageRoutine: Connect Failure", err)
						//TODO: need to handle nil resp channel?
						e.resp <- connectResponse{nil, err}
					} else {
						log.Println("mconn = ", mconn)
						log.Println("mconn.sessionId = ", mconn.sessionId)
						log.Println("cm.conns = ", cm.conns)
						cm.conns[mconn.sessionId] = mconn
						//TODO: need to handle nil resp channel?
						e.resp <- connectResponse{mconn, nil}
					}
				}()

			case Connected:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					log.Println("ManageRoutine: connected ", e.getSessionId())
					log.Println()
					log.Println()
					log.Println()
					log.Println()
				}()

			case disconnect:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					e := e.(disconnect)
					log.Println("ManageRoutine: disconnect ", e.sessionId)
					mconn := cm.conns[e.sessionId]
					mconn.Close()
					//TODO: removal timing matters?
					//delete(cm.conns, e.getSessionId())
					e.resp <- disconnectResponse{nil}
				}()

			case Disconnected:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					log.Println("ManageRoutine: disconnected ", e.getSessionId())
					//TODO: removal timing matters?
					delete(cm.conns, e.getSessionId())
				}()

			case renew:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					log.Println("ManageRoutine: renew to ", e.(renew).addr)
					renewE:= e.(renew)

					// Req disconnect
					disconnectRespCh := make(chan disconnectResponse)
					cm.eventq <- disconnect{e.getSessionId(), disconnectRespCh}

					// Wait for disconnected event
					disconnectResp := <- disconnectRespCh
					if disconnectResp.err != nil {
						log.Printf("ManageRoutine: renew failure: can not disconnect %d. %v\n", e.getSessionId(), disconnectResp.err)
						return
					}

					// Req newsession
					log.Println("ManageRoutine: RenewRoutine: req newsession")
					connectRespCh := make(chan connectResponse)
					cm.eventq <- newsession{renewE.phonenumber,	renewE.addr, renewE.useIPv6, connectRespCh}
					connectResp := <-connectRespCh
					if connectResp.err != nil {
						log.Printf("ManageRoutine: renew failure: can not connect to %s. %v\n", renewE.addr, connectResp.err)
						return
					}
					//TODO: need to handle nil resp channel?
					renewE.resp <- reconnectResponse{connectResp.mconn, nil}
					log.Println("ManageRoutine: RenewRoutine: done")
				}()

			case refresh:
				go func() {
					cm.manageWaitGroup.Add(1)
					defer cm.manageWaitGroup.Done()
					log.Println("ManageRoutine: refresh session ", e.(refresh).sessionId)
					refreshE := e.(refresh)

					// Req disconnect
					disconnectRespCh := make(chan disconnectResponse)
					cm.eventq <- disconnect{e.getSessionId(), disconnectRespCh}

					// Wait for disconnected event
					disconnectResp := <- disconnectRespCh
					if disconnectResp.err != nil {
						log.Printf("ManageRoutine: refresh failure: can not disconnect %d. %v\n", e.getSessionId(), disconnectResp.err)
						return
					}

					// Req loadsession
					log.Println("ManageRoutine: RefreshRoutine: req loadsession")
					connectRespCh := make(chan connectResponse)
					cm.eventq <- loadsession{refreshE.phonenumber, connectRespCh}
					connectResp := <- connectRespCh
					if connectResp.err != nil {
						log.Println("ManageRoutine: refresh failure: ", connectResp.err)
						return
					}
					//TODO: need to handle nil resp channel?
					refreshE.resp <- reconnectResponse{connectResp.mconn, nil}
					log.Println("ManageRoutine: RefreshRoutine: done")

				}()
			}
		}
	}
	log.Println("Manageroutine: done")
}
