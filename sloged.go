package mtproto

import "fmt"

func (x *MManager) logprefixed() string {
	return fmt.Sprintf("[MM %d]", x.managerId)
}

func (x *MConn) logprefixed() string {
	return fmt.Sprintf("[mconn %d]", x.connId)
}

func (x *MSession) logprefixed() string {
	return fmt.Sprintf("[%d-%d]", x.connId, x.sessionId)
}
