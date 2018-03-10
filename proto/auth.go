package mtp

import (
	"errors"
	"fmt"
	"github.com/cjongseok/slog"
	"log"
)

func (mconn *MConn) authSendCode(phonenumber string) (*TypeAuthSentCode, error) {
	session, err := mconn.Session()
	if err != nil {
		return nil, err
	}
	data, err := mconn.InvokeBlocked(&ReqAuthSendCode{
		//Allow_flashcall: false,
		Flags:         0x00000001,
		PhoneNumber:   phonenumber,
		CurrentNumber: &TypeBool{&TypeBool_BoolTrue{&PredBoolTrue{}}},
		ApiId:         session.appConfig.Id,
		ApiHash:       session.appConfig.Hash,
	})

	if err != nil {
		return nil, err
	}

	switch x := data.(type) {
	case *PredAuthSentCode:
		return &TypeAuthSentCode{x}, nil
	default:
		return nil, fmt.Errorf("authSendCode: Got: %T", data)
	}
}

func (cm *MManager) authSendCode(mconn *MConn, phonenumber string) (*MConn, *TypeAuthSentCode, error) {
	for {
		sentCode, err := mconn.authSendCode(phonenumber)
		if err == nil {
			return mconn, sentCode, nil
		} else {
			// Handle RPC error
			if rpcError, ok := err.(TL_rpc_error); ok {
				switch rpcError.error_code {
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

func (mconn *MConn) AuthSignIn(phoneNumber, phoneCode, phoneCodeHash string) (*TypeAuthAuthorization, error) {
	log.Println("Start of signin")
	if phoneNumber == "" || phoneCode == "" || phoneCodeHash == "" {
		return nil, errors.New("MRProto::AuthSignIn one of function parameters is empty")
	}

	x := <-mconn.InvokeNonBlocked(&ReqAuthSignIn{
		PhoneNumber:   phoneNumber,
		PhoneCodeHash: phoneCodeHash,
		PhoneCode:     phoneCode,
	})
	if x.err != nil {
		return nil, x.err
	}

	auth, ok := x.data.(*PredAuthAuthorization)
	if !ok {
		return nil, fmt.Errorf("RPC: %v", x)
	}

	session, err := mconn.Session()
	if err != nil {
		return &TypeAuthAuthorization{auth}, err
	}

	if auth.GetUser().GetUser() != nil {
		session.user = auth.GetUser().GetUser()
		slog.Logln(mconn, "Signed in as ", session.user)
	} else if auth.GetUser().GetUserEmpty() != nil {
		session.user = &PredUser{}
		slog.Logln(mconn, "Signed in with empty user")
	} else {
		session.user = &PredUser{}
		slog.Logln(mconn, "Signed in without user response: neither user nor user empty")
	}
	return &TypeAuthAuthorization{auth}, nil
}

func (mconn *MConn) AuthLogOut() (bool, error) {
	var result bool
	x := <-mconn.InvokeNonBlocked(&ReqAuthLogOut{})
	if x.err != nil {
		return result, x.err
	}

	if tl, ok := x.data.(TL); ok {
		return toBool(tl), nil
	}
	return false, fmt.Errorf("invalid rpc return: %T: %v", x.data, x.data)
}
