package mtproto

import (
	"errors"
	"fmt"
	"log"
)

func (mconn *MConn) authSendCode(phonenumber string) (*TL_auth_sentCode, error) {
	session, err := mconn.Session()
	if err != nil {
		return nil, err
	}
	data, err := mconn.InvokeBlocked(TL_auth_sendCode{
		//Allow_flashcall: false,
		Flags: 0x00000001,
		Phone_number: phonenumber,
		Current_number: TL_boolTrue{},
		Api_id: session.appConfig.Id,
		Api_hash: session.appConfig.Hash,
	})

	if err != nil {
		return nil, err
	}

	switch (*data).(type) {
	case TL_auth_sentCode:
		sentCode := (*data).(TL_auth_sentCode)
		return &sentCode, nil
	default:
		return nil, fmt.Errorf("authSendCode: Got: %T", *data)
	}
}

func (cm *MManager) authSendCode(mconn *MConn, phonenumber string) (*MConn, *TL_auth_sentCode, error) {
	for {
		sentCode, err := mconn.authSendCode(phonenumber)
		if err == nil {
			return mconn, sentCode, nil
		} else {
			// Handle RPC error
			if rpcError, ok := err.(TL_rpc_error); ok {
				switch rpcError.error_code{
				case errorSeeOther:
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
						respch := make(chan sessionResponse)

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
				default:
					return nil, nil, err
				}
			} else {
				return nil, nil, err
			}
		}
	}
}

func (mconn *MConn) AuthSignIn(phoneNumber, phoneCode, phoneCodeHash string) (*TL_auth_authorization, error) {
	log.Println("Start of signin")
	if phoneNumber == "" || phoneCode == "" || phoneCodeHash == "" {
		return nil, errors.New("MRProto::AuthSignIn one of function parameters is empty")
	}

	x := <-mconn.InvokeNonBlocked(TL_auth_signIn{
		Phone_number:    phoneNumber,
		Phone_code_hash: phoneCodeHash,
		Phone_code:      phoneCode,
	})
	if x.err != nil {
		return nil, x.err
	}

	auth, ok := x.data.(TL_auth_authorization)

	if !ok {
		return nil, fmt.Errorf("RPC: %#v", x)
	}

	user := auth.User.(TL_user)
	session, err := mconn.Session()
	if err != nil {
		return &auth, err
	}
	session.user = &user
	log.Println("Signed in as ", auth)

	// save the session
	//session.saveSession()
	return &auth, nil
}

func (mconn *MConn) AuthLogOut() (bool, error) {
	var result bool
	x := <-mconn.InvokeNonBlocked(TL_auth_logOut{})
	if x.err != nil {
		return result, x.err
	}

	result = toBool(x.data)
	return result, nil
}