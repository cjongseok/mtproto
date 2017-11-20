package mtproto

type MEvent interface {
	getSessionId() (int64)
}

type newsession struct {
	phonenumber		string
	addr 			string
	useIPv6 		bool
	resp 			chan connectResponse
}

type loadsession struct {
	phonenumber		string
	resp			chan connectResponse
}

type connectResponse struct {
	mconn 	*MConn
	err 	error
}

type Connected struct {
	sessionId int64
}

type disconnect struct {
	sessionId 	int64
	resp 		chan disconnectResponse
}
type disconnectResponse struct {
	err 	error
}

type Disconnected struct {
	sessionId int64
}

// disconnect + newsession
type renew struct {
	phonenumber	string
	addr 		string
	useIPv6		bool
	sessionId	int64
	resp		chan reconnectResponse
}

// disconnect + loadsession
type refresh struct {
	phonenumber	string
	sessionId 	int64
	resp 		chan reconnectResponse
}
type reconnectResponse struct {
	mconn 	*MConn
	err		error
}

func (e newsession) getSessionId() (int64) {return 0}
func (e loadsession) getSessionId() (int64) {return 0}
func (e Connected) getSessionId() (int64) {return e.sessionId}
func (e renew) getSessionId() (int64) {return e.sessionId}
func (e refresh) getSessionId() (int64) {return e.sessionId}
func (e disconnect) getSessionId() (int64) {return e.sessionId}
func (e Disconnected) getSessionId() (int64) {return e.sessionId}