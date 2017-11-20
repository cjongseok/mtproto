package mtproto

import (
	"errors"
	"fmt"
	"log"
)


func (mconn *MConn) authSendCode(phonenumber string) (*TL_auth_sentCode, error) {
	data, err := mconn.InvokeBlocked(TL_auth_sendCode{
		Allow_flashcall: false,
		Phone_number: phonenumber,
		Current_number: TL_boolTrue{},
		Api_id: mconn.appConfig.Id,
		Api_hash: mconn.appConfig.Hash,
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

func (cm *MManager) AuthSendCode(mconn *MConn, phonenumber string) (*MConn, *TL_auth_sentCode, error) {
	for {
		sentCode, err := mconn.authSendCode(phonenumber)
		if err == nil {
			return mconn, sentCode, nil
		} else {
			// Handle RPC error
			if rpcError, ok := err.(TL_rpc_error); ok {
				switch rpcError.Error_code {
				case errorSeeOther:
					var newdc int32
					n, _ := fmt.Sscanf(rpcError.Error_message, "PHONE_MIGRATE_%d", &newdc)
					if n != 1 {
						n, _ = fmt.Sscanf(rpcError.Error_message, "NETWORK_MIGRATE_%d", &newdc)
					}
					if n != 1 {
						return nil, nil, err
					} else {
						// Reconnect to the new datacenter
						respch := make(chan reconnectResponse)
						mconn.notify(renew{
							phonenumber,
							mconn.dclist[newdc],
							mconn.useIPv6,
							mconn.sessionId,
							respch,
						})

						// Retry with new connection in the next loop
						resp := <-respch
						if resp.err != nil {
							return nil, nil, resp.err
						}
						mconn = resp.mconn
					}
				default:
					return nil, nil, err
				}
			}
		}
	}
}

func (mconn *MConn) AuthSignIn(phoneNumber, phoneCode, phoneCodeHash string) (*TL_auth_authorization, error) {
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
	mconn.user = &user

	log.Println("Signed in as ", auth)


	// save the session
	//mconn.saveSession()

	return &auth, nil
}

func (m *MConn) AuthLogOut() (bool, error) {
	var result bool
	x := <-m.InvokeNonBlocked(TL_auth_logOut{})
	if x.err != nil {
		return result, x.err
	}

	result, err := ToBool(x.data)
	if err != nil {
		return result, err
	}

	return result, err
}
