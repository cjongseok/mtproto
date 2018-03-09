package mtp

import (
	"fmt"
)

func (mconn *MConn) UsersGetFullUsers(id *TypeInputUser) (*TypeUserFull, error) {
	//var user TL_userFull
	x := <-mconn.InvokeNonBlocked(&ReqUsersGetFullUser{id})
	//TL_users_getFullUser{
	//Id: id,
	//})
	if x.err != nil {
		return nil, x.err
	}

	switch casted := x.data.(type) {
	case *PredUserFull:
		return &TypeUserFull{casted}, nil
	}
	return nil, fmt.Errorf("no full user: %T: %v", x, x)
}
