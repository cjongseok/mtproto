package mtp

import "fmt"

func (x *MManager) LogPrefix() string {
	return fmt.Sprintf("[MM %d]", x.managerId)
}

func (x *MConn) LogPrefix() string {
	return fmt.Sprintf("[mconn %d]", x.connId)
}

func (x *MSession) LogPrefix() string {
	return fmt.Sprintf("[%d-%d]", x.connId, x.sessionId)
}
