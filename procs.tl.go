package mtproto

import (
	"fmt"
	"github.com/golang/protobuf/ptypes/any"
	"golang.org/x/net/context"
)

// Procedures
func (caller RPCaller) InvokeAfterMsg(ctx context.Context, req *ReqInvokeAfterMsg) (*any.Any, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if pred, ok := resp.(Predicate); ok {
		packed := Pack(pred.ToType())
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) InvokeAfterMsgs(ctx context.Context, req *ReqInvokeAfterMsgs) (*any.Any, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if pred, ok := resp.(Predicate); ok {
		packed := Pack(pred.ToType())
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthCheckPhone(ctx context.Context, req *ReqAuthCheckPhone) (*TypeAuthCheckedPhone, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthCheckedPhone:
		return x.ToType().(*TypeAuthCheckedPhone), nil
	}
	return &TypeAuthCheckedPhone{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthSendCode(ctx context.Context, req *ReqAuthSendCode) (*TypeAuthSentCode, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthSentCode:
		return x.ToType().(*TypeAuthSentCode), nil
	}
	return &TypeAuthSentCode{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthSignUp(ctx context.Context, req *ReqAuthSignUp) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthSignIn(ctx context.Context, req *ReqAuthSignIn) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthLogOut(ctx context.Context, req *ReqAuthLogOut) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthResetAuthorizations(ctx context.Context, req *ReqAuthResetAuthorizations) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthSendInvites(ctx context.Context, req *ReqAuthSendInvites) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthExportAuthorization(ctx context.Context, req *ReqAuthExportAuthorization) (*TypeAuthExportedAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthExportedAuthorization:
		return x.ToType().(*TypeAuthExportedAuthorization), nil
	}
	return &TypeAuthExportedAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthImportAuthorization(ctx context.Context, req *ReqAuthImportAuthorization) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountRegisterDevice(ctx context.Context, req *ReqAccountRegisterDevice) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUnregisterDevice(ctx context.Context, req *ReqAccountUnregisterDevice) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdateNotifySettings(ctx context.Context, req *ReqAccountUpdateNotifySettings) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetNotifySettings(ctx context.Context, req *ReqAccountGetNotifySettings) (*TypePeerNotifySettings, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPeerNotifySettingsEmpty:
		return x.ToType().(*TypePeerNotifySettings), nil
	case *PredPeerNotifySettings:
		return x.ToType().(*TypePeerNotifySettings), nil
	}
	return &TypePeerNotifySettings{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountResetNotifySettings(ctx context.Context, req *ReqAccountResetNotifySettings) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdateProfile(ctx context.Context, req *ReqAccountUpdateProfile) (*TypeUser, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserEmpty:
		return x.ToType().(*TypeUser), nil
	case *PredUser:
		return x.ToType().(*TypeUser), nil
	}
	return &TypeUser{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdateStatus(ctx context.Context, req *ReqAccountUpdateStatus) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetWallPapers(ctx context.Context, req *ReqAccountGetWallPapers) (*TypeVectorWallPaper, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorWallPaper{}
		v.WallPaper = make([]*TypeWallPaper, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredWallPaper:
				v.WallPaper[i] = x.ToType().(*TypeWallPaper)
			case *PredWallPaperSolid:
				v.WallPaper[i] = x.ToType().(*TypeWallPaper)
			default:
				return nil, fmt.Errorf("invalid WallPaper vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorWallPaper{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UsersGetUsers(ctx context.Context, req *ReqUsersGetUsers) (*TypeVectorUser, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorUser{}
		v.User = make([]*TypeUser, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredUserEmpty:
				v.User[i] = x.ToType().(*TypeUser)
			case *PredUser:
				v.User[i] = x.ToType().(*TypeUser)
			default:
				return nil, fmt.Errorf("invalid User vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorUser{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UsersGetFullUser(ctx context.Context, req *ReqUsersGetFullUser) (*TypeUserFull, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserFull:
		return x.ToType().(*TypeUserFull), nil
	}
	return &TypeUserFull{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsGetStatuses(ctx context.Context, req *ReqContactsGetStatuses) (*TypeVectorContactStatus, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorContactStatus{}
		v.ContactStatus = make([]*TypeContactStatus, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredContactStatus:
				v.ContactStatus[i] = x.ToType().(*TypeContactStatus)
			default:
				return nil, fmt.Errorf("invalid ContactStatus vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorContactStatus{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsGetContacts(ctx context.Context, req *ReqContactsGetContacts) (*TypeContactsContacts, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsContacts:
		return x.ToType().(*TypeContactsContacts), nil
	case *PredContactsContactsNotModified:
		return x.ToType().(*TypeContactsContacts), nil
	}
	return &TypeContactsContacts{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsImportContacts(ctx context.Context, req *ReqContactsImportContacts) (*TypeContactsImportedContacts, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsImportedContacts:
		return x.ToType().(*TypeContactsImportedContacts), nil
	}
	return &TypeContactsImportedContacts{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsSearch(ctx context.Context, req *ReqContactsSearch) (*TypeContactsFound, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsFound:
		return x.ToType().(*TypeContactsFound), nil
	}
	return &TypeContactsFound{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsDeleteContact(ctx context.Context, req *ReqContactsDeleteContact) (*TypeContactsLink, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsLink:
		return x.ToType().(*TypeContactsLink), nil
	}
	return &TypeContactsLink{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsDeleteContacts(ctx context.Context, req *ReqContactsDeleteContacts) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsBlock(ctx context.Context, req *ReqContactsBlock) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsUnblock(ctx context.Context, req *ReqContactsUnblock) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsGetBlocked(ctx context.Context, req *ReqContactsGetBlocked) (*TypeContactsBlocked, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsBlocked:
		return x.ToType().(*TypeContactsBlocked), nil
	case *PredContactsBlockedSlice:
		return x.ToType().(*TypeContactsBlocked), nil
	}
	return &TypeContactsBlocked{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetMessages(ctx context.Context, req *ReqMessagesGetMessages) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetDialogs(ctx context.Context, req *ReqMessagesGetDialogs) (*TypeMessagesDialogs, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesDialogs:
		return x.ToType().(*TypeMessagesDialogs), nil
	case *PredMessagesDialogsSlice:
		return x.ToType().(*TypeMessagesDialogs), nil
	}
	return &TypeMessagesDialogs{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetHistory(ctx context.Context, req *ReqMessagesGetHistory) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSearch(ctx context.Context, req *ReqMessagesSearch) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReadHistory(ctx context.Context, req *ReqMessagesReadHistory) (*TypeMessagesAffectedMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedMessages:
		return x.ToType().(*TypeMessagesAffectedMessages), nil
	}
	return &TypeMessagesAffectedMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesDeleteHistory(ctx context.Context, req *ReqMessagesDeleteHistory) (*TypeMessagesAffectedHistory, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedHistory:
		return x.ToType().(*TypeMessagesAffectedHistory), nil
	}
	return &TypeMessagesAffectedHistory{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesDeleteMessages(ctx context.Context, req *ReqMessagesDeleteMessages) (*TypeMessagesAffectedMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedMessages:
		return x.ToType().(*TypeMessagesAffectedMessages), nil
	}
	return &TypeMessagesAffectedMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReceivedMessages(ctx context.Context, req *ReqMessagesReceivedMessages) (*TypeVectorReceivedNotifyMessage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorReceivedNotifyMessage{}
		v.ReceivedNotifyMessage = make([]*TypeReceivedNotifyMessage, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredReceivedNotifyMessage:
				v.ReceivedNotifyMessage[i] = x.ToType().(*TypeReceivedNotifyMessage)
			default:
				return nil, fmt.Errorf("invalid ReceivedNotifyMessage vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorReceivedNotifyMessage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetTyping(ctx context.Context, req *ReqMessagesSetTyping) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendMessage(ctx context.Context, req *ReqMessagesSendMessage) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendMedia(ctx context.Context, req *ReqMessagesSendMedia) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesForwardMessages(ctx context.Context, req *ReqMessagesForwardMessages) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetChats(ctx context.Context, req *ReqMessagesGetChats) (*TypeMessagesChats, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChats:
		return x.ToType().(*TypeMessagesChats), nil
	case *PredMessagesChatsSlice:
		return x.ToType().(*TypeMessagesChats), nil
	}
	return &TypeMessagesChats{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetFullChat(ctx context.Context, req *ReqMessagesGetFullChat) (*TypeMessagesChatFull, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChatFull:
		return x.ToType().(*TypeMessagesChatFull), nil
	}
	return &TypeMessagesChatFull{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesEditChatTitle(ctx context.Context, req *ReqMessagesEditChatTitle) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesEditChatPhoto(ctx context.Context, req *ReqMessagesEditChatPhoto) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesAddChatUser(ctx context.Context, req *ReqMessagesAddChatUser) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesDeleteChatUser(ctx context.Context, req *ReqMessagesDeleteChatUser) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesCreateChat(ctx context.Context, req *ReqMessagesCreateChat) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UpdatesGetState(ctx context.Context, req *ReqUpdatesGetState) (*TypeUpdatesState, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesState:
		return x.ToType().(*TypeUpdatesState), nil
	}
	return &TypeUpdatesState{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UpdatesGetDifference(ctx context.Context, req *ReqUpdatesGetDifference) (*TypeUpdatesDifference, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesDifferenceEmpty:
		return x.ToType().(*TypeUpdatesDifference), nil
	case *PredUpdatesDifference:
		return x.ToType().(*TypeUpdatesDifference), nil
	case *PredUpdatesDifferenceSlice:
		return x.ToType().(*TypeUpdatesDifference), nil
	case *PredUpdatesDifferenceTooLong:
		return x.ToType().(*TypeUpdatesDifference), nil
	}
	return &TypeUpdatesDifference{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhotosUpdateProfilePhoto(ctx context.Context, req *ReqPhotosUpdateProfilePhoto) (*TypeUserProfilePhoto, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserProfilePhotoEmpty:
		return x.ToType().(*TypeUserProfilePhoto), nil
	case *PredUserProfilePhoto:
		return x.ToType().(*TypeUserProfilePhoto), nil
	}
	return &TypeUserProfilePhoto{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhotosUploadProfilePhoto(ctx context.Context, req *ReqPhotosUploadProfilePhoto) (*TypePhotosPhoto, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPhotosPhoto:
		return x.ToType().(*TypePhotosPhoto), nil
	}
	return &TypePhotosPhoto{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadSaveFilePart(ctx context.Context, req *ReqUploadSaveFilePart) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadGetFile(ctx context.Context, req *ReqUploadGetFile) (*TypeUploadFile, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUploadFile:
		return x.ToType().(*TypeUploadFile), nil
	case *PredUploadFileCdnRedirect:
		return x.ToType().(*TypeUploadFile), nil
	}
	return &TypeUploadFile{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetConfig(ctx context.Context, req *ReqHelpGetConfig) (*TypeConfig, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredConfig:
		return x.ToType().(*TypeConfig), nil
	}
	return &TypeConfig{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetNearestDc(ctx context.Context, req *ReqHelpGetNearestDc) (*TypeNearestDc, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredNearestDc:
		return x.ToType().(*TypeNearestDc), nil
	}
	return &TypeNearestDc{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetAppUpdate(ctx context.Context, req *ReqHelpGetAppUpdate) (*TypeHelpAppUpdate, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredHelpAppUpdate:
		return x.ToType().(*TypeHelpAppUpdate), nil
	case *PredHelpNoAppUpdate:
		return x.ToType().(*TypeHelpAppUpdate), nil
	}
	return &TypeHelpAppUpdate{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpSaveAppLog(ctx context.Context, req *ReqHelpSaveAppLog) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetInviteText(ctx context.Context, req *ReqHelpGetInviteText) (*TypeHelpInviteText, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredHelpInviteText:
		return x.ToType().(*TypeHelpInviteText), nil
	}
	return &TypeHelpInviteText{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhotosDeletePhotos(ctx context.Context, req *ReqPhotosDeletePhotos) (*TypeVectorLong, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int64); ok {
		return &TypeVectorLong{Values: resp}, nil
	}
	return &TypeVectorLong{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhotosGetUserPhotos(ctx context.Context, req *ReqPhotosGetUserPhotos) (*TypePhotosPhotos, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPhotosPhotos:
		return x.ToType().(*TypePhotosPhotos), nil
	case *PredPhotosPhotosSlice:
		return x.ToType().(*TypePhotosPhotos), nil
	}
	return &TypePhotosPhotos{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesForwardMessage(ctx context.Context, req *ReqMessagesForwardMessage) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetDhConfig(ctx context.Context, req *ReqMessagesGetDhConfig) (*TypeMessagesDhConfig, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesDhConfigNotModified:
		return x.ToType().(*TypeMessagesDhConfig), nil
	case *PredMessagesDhConfig:
		return x.ToType().(*TypeMessagesDhConfig), nil
	}
	return &TypeMessagesDhConfig{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesRequestEncryption(ctx context.Context, req *ReqMessagesRequestEncryption) (*TypeEncryptedChat, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredEncryptedChatEmpty:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatWaiting:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatRequested:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChat:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatDiscarded:
		return x.ToType().(*TypeEncryptedChat), nil
	}
	return &TypeEncryptedChat{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesAcceptEncryption(ctx context.Context, req *ReqMessagesAcceptEncryption) (*TypeEncryptedChat, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredEncryptedChatEmpty:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatWaiting:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatRequested:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChat:
		return x.ToType().(*TypeEncryptedChat), nil
	case *PredEncryptedChatDiscarded:
		return x.ToType().(*TypeEncryptedChat), nil
	}
	return &TypeEncryptedChat{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesDiscardEncryption(ctx context.Context, req *ReqMessagesDiscardEncryption) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetEncryptedTyping(ctx context.Context, req *ReqMessagesSetEncryptedTyping) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReadEncryptedHistory(ctx context.Context, req *ReqMessagesReadEncryptedHistory) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendEncrypted(ctx context.Context, req *ReqMessagesSendEncrypted) (*TypeMessagesSentEncryptedMessage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesSentEncryptedMessage:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	case *PredMessagesSentEncryptedFile:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	}
	return &TypeMessagesSentEncryptedMessage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendEncryptedFile(ctx context.Context, req *ReqMessagesSendEncryptedFile) (*TypeMessagesSentEncryptedMessage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesSentEncryptedMessage:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	case *PredMessagesSentEncryptedFile:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	}
	return &TypeMessagesSentEncryptedMessage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendEncryptedService(ctx context.Context, req *ReqMessagesSendEncryptedService) (*TypeMessagesSentEncryptedMessage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesSentEncryptedMessage:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	case *PredMessagesSentEncryptedFile:
		return x.ToType().(*TypeMessagesSentEncryptedMessage), nil
	}
	return &TypeMessagesSentEncryptedMessage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReceivedQueue(ctx context.Context, req *ReqMessagesReceivedQueue) (*TypeVectorLong, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int64); ok {
		return &TypeVectorLong{Values: resp}, nil
	}
	return &TypeVectorLong{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadSaveBigFilePart(ctx context.Context, req *ReqUploadSaveBigFilePart) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) InitConnection(ctx context.Context, req *ReqInitConnection) (*any.Any, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if pred, ok := resp.(Predicate); ok {
		packed := Pack(pred.ToType())
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetSupport(ctx context.Context, req *ReqHelpGetSupport) (*TypeHelpSupport, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredHelpSupport:
		return x.ToType().(*TypeHelpSupport), nil
	}
	return &TypeHelpSupport{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthBindTempAuthKey(ctx context.Context, req *ReqAuthBindTempAuthKey) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsExportCard(ctx context.Context, req *ReqContactsExportCard) (*TypeVectorInt, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int32); ok {
		return &TypeVectorInt{Values: resp}, nil
	}
	return &TypeVectorInt{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsImportCard(ctx context.Context, req *ReqContactsImportCard) (*TypeUser, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserEmpty:
		return x.ToType().(*TypeUser), nil
	case *PredUser:
		return x.ToType().(*TypeUser), nil
	}
	return &TypeUser{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReadMessageContents(ctx context.Context, req *ReqMessagesReadMessageContents) (*TypeMessagesAffectedMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedMessages:
		return x.ToType().(*TypeMessagesAffectedMessages), nil
	}
	return &TypeMessagesAffectedMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountCheckUsername(ctx context.Context, req *ReqAccountCheckUsername) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdateUsername(ctx context.Context, req *ReqAccountUpdateUsername) (*TypeUser, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserEmpty:
		return x.ToType().(*TypeUser), nil
	case *PredUser:
		return x.ToType().(*TypeUser), nil
	}
	return &TypeUser{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetPrivacy(ctx context.Context, req *ReqAccountGetPrivacy) (*TypeAccountPrivacyRules, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountPrivacyRules:
		return x.ToType().(*TypeAccountPrivacyRules), nil
	}
	return &TypeAccountPrivacyRules{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountSetPrivacy(ctx context.Context, req *ReqAccountSetPrivacy) (*TypeAccountPrivacyRules, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountPrivacyRules:
		return x.ToType().(*TypeAccountPrivacyRules), nil
	}
	return &TypeAccountPrivacyRules{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountDeleteAccount(ctx context.Context, req *ReqAccountDeleteAccount) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetAccountTTL(ctx context.Context, req *ReqAccountGetAccountTTL) (*TypeAccountDaysTTL, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountDaysTTL:
		return x.ToType().(*TypeAccountDaysTTL), nil
	}
	return &TypeAccountDaysTTL{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountSetAccountTTL(ctx context.Context, req *ReqAccountSetAccountTTL) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) InvokeWithLayer(ctx context.Context, req *ReqInvokeWithLayer) (*any.Any, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if pred, ok := resp.(Predicate); ok {
		packed := Pack(pred.ToType())
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsResolveUsername(ctx context.Context, req *ReqContactsResolveUsername) (*TypeContactsResolvedPeer, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsResolvedPeer:
		return x.ToType().(*TypeContactsResolvedPeer), nil
	}
	return &TypeContactsResolvedPeer{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountSendChangePhoneCode(ctx context.Context, req *ReqAccountSendChangePhoneCode) (*TypeAuthSentCode, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthSentCode:
		return x.ToType().(*TypeAuthSentCode), nil
	}
	return &TypeAuthSentCode{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountChangePhone(ctx context.Context, req *ReqAccountChangePhone) (*TypeUser, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUserEmpty:
		return x.ToType().(*TypeUser), nil
	case *PredUser:
		return x.ToType().(*TypeUser), nil
	}
	return &TypeUser{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetAllStickers(ctx context.Context, req *ReqMessagesGetAllStickers) (*TypeMessagesAllStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAllStickersNotModified:
		return x.ToType().(*TypeMessagesAllStickers), nil
	case *PredMessagesAllStickers:
		return x.ToType().(*TypeMessagesAllStickers), nil
	}
	return &TypeMessagesAllStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdateDeviceLocked(ctx context.Context, req *ReqAccountUpdateDeviceLocked) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetPassword(ctx context.Context, req *ReqAccountGetPassword) (*TypeAccountPassword, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountNoPassword:
		return x.ToType().(*TypeAccountPassword), nil
	case *PredAccountPassword:
		return x.ToType().(*TypeAccountPassword), nil
	}
	return &TypeAccountPassword{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthCheckPassword(ctx context.Context, req *ReqAuthCheckPassword) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetWebPagePreview(ctx context.Context, req *ReqMessagesGetWebPagePreview) (*TypeMessageMedia, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessageMediaEmpty:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaPhoto:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaGeo:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaContact:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaUnsupported:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaDocument:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaWebPage:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaVenue:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaGame:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaInvoice:
		return x.ToType().(*TypeMessageMedia), nil
	}
	return &TypeMessageMedia{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetAuthorizations(ctx context.Context, req *ReqAccountGetAuthorizations) (*TypeAccountAuthorizations, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountAuthorizations:
		return x.ToType().(*TypeAccountAuthorizations), nil
	}
	return &TypeAccountAuthorizations{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountResetAuthorization(ctx context.Context, req *ReqAccountResetAuthorization) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetPasswordSettings(ctx context.Context, req *ReqAccountGetPasswordSettings) (*TypeAccountPasswordSettings, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountPasswordSettings:
		return x.ToType().(*TypeAccountPasswordSettings), nil
	}
	return &TypeAccountPasswordSettings{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountUpdatePasswordSettings(ctx context.Context, req *ReqAccountUpdatePasswordSettings) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthRequestPasswordRecovery(ctx context.Context, req *ReqAuthRequestPasswordRecovery) (*TypeAuthPasswordRecovery, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthPasswordRecovery:
		return x.ToType().(*TypeAuthPasswordRecovery), nil
	}
	return &TypeAuthPasswordRecovery{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthRecoverPassword(ctx context.Context, req *ReqAuthRecoverPassword) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) InvokeWithoutUpdates(ctx context.Context, req *ReqInvokeWithoutUpdates) (*any.Any, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if pred, ok := resp.(Predicate); ok {
		packed := Pack(pred.ToType())
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesExportChatInvite(ctx context.Context, req *ReqMessagesExportChatInvite) (*TypeExportedChatInvite, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChatInviteEmpty:
		return x.ToType().(*TypeExportedChatInvite), nil
	case *PredChatInviteExported:
		return x.ToType().(*TypeExportedChatInvite), nil
	}
	return &TypeExportedChatInvite{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesCheckChatInvite(ctx context.Context, req *ReqMessagesCheckChatInvite) (*TypeChatInvite, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChatInviteAlready:
		return x.ToType().(*TypeChatInvite), nil
	case *PredChatInvite:
		return x.ToType().(*TypeChatInvite), nil
	}
	return &TypeChatInvite{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesImportChatInvite(ctx context.Context, req *ReqMessagesImportChatInvite) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetStickerSet(ctx context.Context, req *ReqMessagesGetStickerSet) (*TypeMessagesStickerSet, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSet:
		return x.ToType().(*TypeMessagesStickerSet), nil
	}
	return &TypeMessagesStickerSet{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesInstallStickerSet(ctx context.Context, req *ReqMessagesInstallStickerSet) (*TypeMessagesStickerSetInstallResult, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSetInstallResultSuccess:
		return x.ToType().(*TypeMessagesStickerSetInstallResult), nil
	case *PredMessagesStickerSetInstallResultArchive:
		return x.ToType().(*TypeMessagesStickerSetInstallResult), nil
	}
	return &TypeMessagesStickerSetInstallResult{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesUninstallStickerSet(ctx context.Context, req *ReqMessagesUninstallStickerSet) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthImportBotAuthorization(ctx context.Context, req *ReqAuthImportBotAuthorization) (*TypeAuthAuthorization, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthAuthorization:
		return x.ToType().(*TypeAuthAuthorization), nil
	}
	return &TypeAuthAuthorization{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesStartBot(ctx context.Context, req *ReqMessagesStartBot) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetAppChangelog(ctx context.Context, req *ReqHelpGetAppChangelog) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReportSpam(ctx context.Context, req *ReqMessagesReportSpam) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetMessagesViews(ctx context.Context, req *ReqMessagesGetMessagesViews) (*TypeVectorInt, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int32); ok {
		return &TypeVectorInt{Values: resp}, nil
	}
	return &TypeVectorInt{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UpdatesGetChannelDifference(ctx context.Context, req *ReqUpdatesGetChannelDifference) (*TypeUpdatesChannelDifference, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesChannelDifferenceEmpty:
		return x.ToType().(*TypeUpdatesChannelDifference), nil
	case *PredUpdatesChannelDifferenceTooLong:
		return x.ToType().(*TypeUpdatesChannelDifference), nil
	case *PredUpdatesChannelDifference:
		return x.ToType().(*TypeUpdatesChannelDifference), nil
	}
	return &TypeUpdatesChannelDifference{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsReadHistory(ctx context.Context, req *ReqChannelsReadHistory) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsDeleteMessages(ctx context.Context, req *ReqChannelsDeleteMessages) (*TypeMessagesAffectedMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedMessages:
		return x.ToType().(*TypeMessagesAffectedMessages), nil
	}
	return &TypeMessagesAffectedMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsDeleteUserHistory(ctx context.Context, req *ReqChannelsDeleteUserHistory) (*TypeMessagesAffectedHistory, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAffectedHistory:
		return x.ToType().(*TypeMessagesAffectedHistory), nil
	}
	return &TypeMessagesAffectedHistory{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsReportSpam(ctx context.Context, req *ReqChannelsReportSpam) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetMessages(ctx context.Context, req *ReqChannelsGetMessages) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetParticipants(ctx context.Context, req *ReqChannelsGetParticipants) (*TypeChannelsChannelParticipants, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChannelsChannelParticipants:
		return x.ToType().(*TypeChannelsChannelParticipants), nil
	}
	return &TypeChannelsChannelParticipants{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetParticipant(ctx context.Context, req *ReqChannelsGetParticipant) (*TypeChannelsChannelParticipant, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChannelsChannelParticipant:
		return x.ToType().(*TypeChannelsChannelParticipant), nil
	}
	return &TypeChannelsChannelParticipant{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetChannels(ctx context.Context, req *ReqChannelsGetChannels) (*TypeMessagesChats, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChats:
		return x.ToType().(*TypeMessagesChats), nil
	case *PredMessagesChatsSlice:
		return x.ToType().(*TypeMessagesChats), nil
	}
	return &TypeMessagesChats{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetFullChannel(ctx context.Context, req *ReqChannelsGetFullChannel) (*TypeMessagesChatFull, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChatFull:
		return x.ToType().(*TypeMessagesChatFull), nil
	}
	return &TypeMessagesChatFull{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsCreateChannel(ctx context.Context, req *ReqChannelsCreateChannel) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsEditAbout(ctx context.Context, req *ReqChannelsEditAbout) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsEditAdmin(ctx context.Context, req *ReqChannelsEditAdmin) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsEditTitle(ctx context.Context, req *ReqChannelsEditTitle) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsEditPhoto(ctx context.Context, req *ReqChannelsEditPhoto) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsCheckUsername(ctx context.Context, req *ReqChannelsCheckUsername) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsUpdateUsername(ctx context.Context, req *ReqChannelsUpdateUsername) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsJoinChannel(ctx context.Context, req *ReqChannelsJoinChannel) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsLeaveChannel(ctx context.Context, req *ReqChannelsLeaveChannel) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsInviteToChannel(ctx context.Context, req *ReqChannelsInviteToChannel) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsExportInvite(ctx context.Context, req *ReqChannelsExportInvite) (*TypeExportedChatInvite, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChatInviteEmpty:
		return x.ToType().(*TypeExportedChatInvite), nil
	case *PredChatInviteExported:
		return x.ToType().(*TypeExportedChatInvite), nil
	}
	return &TypeExportedChatInvite{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsDeleteChannel(ctx context.Context, req *ReqChannelsDeleteChannel) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesToggleChatAdmins(ctx context.Context, req *ReqMessagesToggleChatAdmins) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesEditChatAdmin(ctx context.Context, req *ReqMessagesEditChatAdmin) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesMigrateChat(ctx context.Context, req *ReqMessagesMigrateChat) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSearchGlobal(ctx context.Context, req *ReqMessagesSearchGlobal) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountReportPeer(ctx context.Context, req *ReqAccountReportPeer) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReorderStickerSets(ctx context.Context, req *ReqMessagesReorderStickerSets) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetTermsOfService(ctx context.Context, req *ReqHelpGetTermsOfService) (*TypeHelpTermsOfService, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredHelpTermsOfService:
		return x.ToType().(*TypeHelpTermsOfService), nil
	}
	return &TypeHelpTermsOfService{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetDocumentByHash(ctx context.Context, req *ReqMessagesGetDocumentByHash) (*TypeDocument, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredDocumentEmpty:
		return x.ToType().(*TypeDocument), nil
	case *PredDocument:
		return x.ToType().(*TypeDocument), nil
	}
	return &TypeDocument{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSearchGifs(ctx context.Context, req *ReqMessagesSearchGifs) (*TypeMessagesFoundGifs, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesFoundGifs:
		return x.ToType().(*TypeMessagesFoundGifs), nil
	}
	return &TypeMessagesFoundGifs{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetSavedGifs(ctx context.Context, req *ReqMessagesGetSavedGifs) (*TypeMessagesSavedGifs, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesSavedGifsNotModified:
		return x.ToType().(*TypeMessagesSavedGifs), nil
	case *PredMessagesSavedGifs:
		return x.ToType().(*TypeMessagesSavedGifs), nil
	}
	return &TypeMessagesSavedGifs{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSaveGif(ctx context.Context, req *ReqMessagesSaveGif) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetInlineBotResults(ctx context.Context, req *ReqMessagesGetInlineBotResults) (*TypeMessagesBotResults, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesBotResults:
		return x.ToType().(*TypeMessagesBotResults), nil
	}
	return &TypeMessagesBotResults{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetInlineBotResults(ctx context.Context, req *ReqMessagesSetInlineBotResults) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendInlineBotResult(ctx context.Context, req *ReqMessagesSendInlineBotResult) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsToggleInvites(ctx context.Context, req *ReqChannelsToggleInvites) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsExportMessageLink(ctx context.Context, req *ReqChannelsExportMessageLink) (*TypeExportedMessageLink, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredExportedMessageLink:
		return x.ToType().(*TypeExportedMessageLink), nil
	}
	return &TypeExportedMessageLink{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsToggleSignatures(ctx context.Context, req *ReqChannelsToggleSignatures) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesHideReportSpam(ctx context.Context, req *ReqMessagesHideReportSpam) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetPeerSettings(ctx context.Context, req *ReqMessagesGetPeerSettings) (*TypePeerSettings, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPeerSettings:
		return x.ToType().(*TypePeerSettings), nil
	}
	return &TypePeerSettings{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsUpdatePinnedMessage(ctx context.Context, req *ReqChannelsUpdatePinnedMessage) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthResendCode(ctx context.Context, req *ReqAuthResendCode) (*TypeAuthSentCode, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthSentCode:
		return x.ToType().(*TypeAuthSentCode), nil
	}
	return &TypeAuthSentCode{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthCancelCode(ctx context.Context, req *ReqAuthCancelCode) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetMessageEditData(ctx context.Context, req *ReqMessagesGetMessageEditData) (*TypeMessagesMessageEditData, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessageEditData:
		return x.ToType().(*TypeMessagesMessageEditData), nil
	}
	return &TypeMessagesMessageEditData{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesEditMessage(ctx context.Context, req *ReqMessagesEditMessage) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesEditInlineBotMessage(ctx context.Context, req *ReqMessagesEditInlineBotMessage) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetBotCallbackAnswer(ctx context.Context, req *ReqMessagesGetBotCallbackAnswer) (*TypeMessagesBotCallbackAnswer, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesBotCallbackAnswer:
		return x.ToType().(*TypeMessagesBotCallbackAnswer), nil
	}
	return &TypeMessagesBotCallbackAnswer{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetBotCallbackAnswer(ctx context.Context, req *ReqMessagesSetBotCallbackAnswer) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsGetTopPeers(ctx context.Context, req *ReqContactsGetTopPeers) (*TypeContactsTopPeers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredContactsTopPeersNotModified:
		return x.ToType().(*TypeContactsTopPeers), nil
	case *PredContactsTopPeers:
		return x.ToType().(*TypeContactsTopPeers), nil
	}
	return &TypeContactsTopPeers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsResetTopPeerRating(ctx context.Context, req *ReqContactsResetTopPeerRating) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetPeerDialogs(ctx context.Context, req *ReqMessagesGetPeerDialogs) (*TypeMessagesPeerDialogs, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesPeerDialogs:
		return x.ToType().(*TypeMessagesPeerDialogs), nil
	}
	return &TypeMessagesPeerDialogs{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSaveDraft(ctx context.Context, req *ReqMessagesSaveDraft) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetAllDrafts(ctx context.Context, req *ReqMessagesGetAllDrafts) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountSendConfirmPhoneCode(ctx context.Context, req *ReqAccountSendConfirmPhoneCode) (*TypeAuthSentCode, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAuthSentCode:
		return x.ToType().(*TypeAuthSentCode), nil
	}
	return &TypeAuthSentCode{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountConfirmPhone(ctx context.Context, req *ReqAccountConfirmPhone) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetFeaturedStickers(ctx context.Context, req *ReqMessagesGetFeaturedStickers) (*TypeMessagesFeaturedStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesFeaturedStickersNotModified:
		return x.ToType().(*TypeMessagesFeaturedStickers), nil
	case *PredMessagesFeaturedStickers:
		return x.ToType().(*TypeMessagesFeaturedStickers), nil
	}
	return &TypeMessagesFeaturedStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReadFeaturedStickers(ctx context.Context, req *ReqMessagesReadFeaturedStickers) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetRecentStickers(ctx context.Context, req *ReqMessagesGetRecentStickers) (*TypeMessagesRecentStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesRecentStickersNotModified:
		return x.ToType().(*TypeMessagesRecentStickers), nil
	case *PredMessagesRecentStickers:
		return x.ToType().(*TypeMessagesRecentStickers), nil
	}
	return &TypeMessagesRecentStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSaveRecentSticker(ctx context.Context, req *ReqMessagesSaveRecentSticker) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesClearRecentStickers(ctx context.Context, req *ReqMessagesClearRecentStickers) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetArchivedStickers(ctx context.Context, req *ReqMessagesGetArchivedStickers) (*TypeMessagesArchivedStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesArchivedStickers:
		return x.ToType().(*TypeMessagesArchivedStickers), nil
	}
	return &TypeMessagesArchivedStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetAdminedPublicChannels(ctx context.Context, req *ReqChannelsGetAdminedPublicChannels) (*TypeMessagesChats, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChats:
		return x.ToType().(*TypeMessagesChats), nil
	case *PredMessagesChatsSlice:
		return x.ToType().(*TypeMessagesChats), nil
	}
	return &TypeMessagesChats{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AuthDropTempAuthKeys(ctx context.Context, req *ReqAuthDropTempAuthKeys) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetGameScore(ctx context.Context, req *ReqMessagesSetGameScore) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetInlineGameScore(ctx context.Context, req *ReqMessagesSetInlineGameScore) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetMaskStickers(ctx context.Context, req *ReqMessagesGetMaskStickers) (*TypeMessagesAllStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesAllStickersNotModified:
		return x.ToType().(*TypeMessagesAllStickers), nil
	case *PredMessagesAllStickers:
		return x.ToType().(*TypeMessagesAllStickers), nil
	}
	return &TypeMessagesAllStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetAttachedStickers(ctx context.Context, req *ReqMessagesGetAttachedStickers) (*TypeVectorStickerSetCovered, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorStickerSetCovered{}
		v.StickerSetCovered = make([]*TypeStickerSetCovered, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredStickerSetCovered:
				v.StickerSetCovered[i] = x.ToType().(*TypeStickerSetCovered)
			case *PredStickerSetMultiCovered:
				v.StickerSetCovered[i] = x.ToType().(*TypeStickerSetCovered)
			default:
				return nil, fmt.Errorf("invalid StickerSetCovered vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorStickerSetCovered{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetGameHighScores(ctx context.Context, req *ReqMessagesGetGameHighScores) (*TypeMessagesHighScores, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesHighScores:
		return x.ToType().(*TypeMessagesHighScores), nil
	}
	return &TypeMessagesHighScores{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetInlineGameHighScores(ctx context.Context, req *ReqMessagesGetInlineGameHighScores) (*TypeMessagesHighScores, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesHighScores:
		return x.ToType().(*TypeMessagesHighScores), nil
	}
	return &TypeMessagesHighScores{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetCommonChats(ctx context.Context, req *ReqMessagesGetCommonChats) (*TypeMessagesChats, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChats:
		return x.ToType().(*TypeMessagesChats), nil
	case *PredMessagesChatsSlice:
		return x.ToType().(*TypeMessagesChats), nil
	}
	return &TypeMessagesChats{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetAllChats(ctx context.Context, req *ReqMessagesGetAllChats) (*TypeMessagesChats, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesChats:
		return x.ToType().(*TypeMessagesChats), nil
	case *PredMessagesChatsSlice:
		return x.ToType().(*TypeMessagesChats), nil
	}
	return &TypeMessagesChats{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpSetBotUpdatesStatus(ctx context.Context, req *ReqHelpSetBotUpdatesStatus) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetWebPage(ctx context.Context, req *ReqMessagesGetWebPage) (*TypeWebPage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredWebPageEmpty:
		return x.ToType().(*TypeWebPage), nil
	case *PredWebPagePending:
		return x.ToType().(*TypeWebPage), nil
	case *PredWebPage:
		return x.ToType().(*TypeWebPage), nil
	case *PredWebPageNotModified:
		return x.ToType().(*TypeWebPage), nil
	}
	return &TypeWebPage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesToggleDialogPin(ctx context.Context, req *ReqMessagesToggleDialogPin) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReorderPinnedDialogs(ctx context.Context, req *ReqMessagesReorderPinnedDialogs) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetPinnedDialogs(ctx context.Context, req *ReqMessagesGetPinnedDialogs) (*TypeMessagesPeerDialogs, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesPeerDialogs:
		return x.ToType().(*TypeMessagesPeerDialogs), nil
	}
	return &TypeMessagesPeerDialogs{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneRequestCall(ctx context.Context, req *ReqPhoneRequestCall) (*TypePhonePhoneCall, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPhonePhoneCall:
		return x.ToType().(*TypePhonePhoneCall), nil
	}
	return &TypePhonePhoneCall{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneAcceptCall(ctx context.Context, req *ReqPhoneAcceptCall) (*TypePhonePhoneCall, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPhonePhoneCall:
		return x.ToType().(*TypePhonePhoneCall), nil
	}
	return &TypePhonePhoneCall{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneDiscardCall(ctx context.Context, req *ReqPhoneDiscardCall) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneReceivedCall(ctx context.Context, req *ReqPhoneReceivedCall) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesReportEncryptedSpam(ctx context.Context, req *ReqMessagesReportEncryptedSpam) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsGetPaymentForm(ctx context.Context, req *ReqPaymentsGetPaymentForm) (*TypePaymentsPaymentForm, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPaymentsPaymentForm:
		return x.ToType().(*TypePaymentsPaymentForm), nil
	}
	return &TypePaymentsPaymentForm{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsSendPaymentForm(ctx context.Context, req *ReqPaymentsSendPaymentForm) (*TypePaymentsPaymentResult, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPaymentsPaymentResult:
		return x.ToType().(*TypePaymentsPaymentResult), nil
	case *PredPaymentsPaymentVerficationNeeded:
		return x.ToType().(*TypePaymentsPaymentResult), nil
	}
	return &TypePaymentsPaymentResult{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) AccountGetTmpPassword(ctx context.Context, req *ReqAccountGetTmpPassword) (*TypeAccountTmpPassword, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredAccountTmpPassword:
		return x.ToType().(*TypeAccountTmpPassword), nil
	}
	return &TypeAccountTmpPassword{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetBotShippingResults(ctx context.Context, req *ReqMessagesSetBotShippingResults) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSetBotPrecheckoutResults(ctx context.Context, req *ReqMessagesSetBotPrecheckoutResults) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadGetWebFile(ctx context.Context, req *ReqUploadGetWebFile) (*TypeUploadWebFile, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUploadWebFile:
		return x.ToType().(*TypeUploadWebFile), nil
	}
	return &TypeUploadWebFile{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) BotsSendCustomRequest(ctx context.Context, req *ReqBotsSendCustomRequest) (*TypeDataJSON, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredDataJSON:
		return x.ToType().(*TypeDataJSON), nil
	}
	return &TypeDataJSON{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) BotsAnswerWebhookJSONQuery(ctx context.Context, req *ReqBotsAnswerWebhookJSONQuery) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsGetPaymentReceipt(ctx context.Context, req *ReqPaymentsGetPaymentReceipt) (*TypePaymentsPaymentReceipt, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPaymentsPaymentReceipt:
		return x.ToType().(*TypePaymentsPaymentReceipt), nil
	}
	return &TypePaymentsPaymentReceipt{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsValidateRequestedInfo(ctx context.Context, req *ReqPaymentsValidateRequestedInfo) (*TypePaymentsValidatedRequestedInfo, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPaymentsValidatedRequestedInfo:
		return x.ToType().(*TypePaymentsValidatedRequestedInfo), nil
	}
	return &TypePaymentsValidatedRequestedInfo{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsGetSavedInfo(ctx context.Context, req *ReqPaymentsGetSavedInfo) (*TypePaymentsSavedInfo, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPaymentsSavedInfo:
		return x.ToType().(*TypePaymentsSavedInfo), nil
	}
	return &TypePaymentsSavedInfo{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PaymentsClearSavedInfo(ctx context.Context, req *ReqPaymentsClearSavedInfo) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneGetCallConfig(ctx context.Context, req *ReqPhoneGetCallConfig) (*TypeDataJSON, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredDataJSON:
		return x.ToType().(*TypeDataJSON), nil
	}
	return &TypeDataJSON{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneConfirmCall(ctx context.Context, req *ReqPhoneConfirmCall) (*TypePhonePhoneCall, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredPhonePhoneCall:
		return x.ToType().(*TypePhonePhoneCall), nil
	}
	return &TypePhonePhoneCall{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneSetCallRating(ctx context.Context, req *ReqPhoneSetCallRating) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) PhoneSaveCallDebug(ctx context.Context, req *ReqPhoneSaveCallDebug) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadGetCdnFile(ctx context.Context, req *ReqUploadGetCdnFile) (*TypeUploadCdnFile, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUploadCdnFileReuploadNeeded:
		return x.ToType().(*TypeUploadCdnFile), nil
	case *PredUploadCdnFile:
		return x.ToType().(*TypeUploadCdnFile), nil
	}
	return &TypeUploadCdnFile{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadReuploadCdnFile(ctx context.Context, req *ReqUploadReuploadCdnFile) (*TypeVectorCdnFileHash, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorCdnFileHash{}
		v.CdnFileHash = make([]*TypeCdnFileHash, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredCdnFileHash:
				v.CdnFileHash[i] = x.ToType().(*TypeCdnFileHash)
			default:
				return nil, fmt.Errorf("invalid CdnFileHash vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorCdnFileHash{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) HelpGetCdnConfig(ctx context.Context, req *ReqHelpGetCdnConfig) (*TypeCdnConfig, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredCdnConfig:
		return x.ToType().(*TypeCdnConfig), nil
	}
	return &TypeCdnConfig{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesUploadMedia(ctx context.Context, req *ReqMessagesUploadMedia) (*TypeMessageMedia, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessageMediaEmpty:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaPhoto:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaGeo:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaContact:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaUnsupported:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaDocument:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaWebPage:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaVenue:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaGame:
		return x.ToType().(*TypeMessageMedia), nil
	case *PredMessageMediaInvoice:
		return x.ToType().(*TypeMessageMedia), nil
	}
	return &TypeMessageMedia{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) StickersCreateStickerSet(ctx context.Context, req *ReqStickersCreateStickerSet) (*TypeMessagesStickerSet, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSet:
		return x.ToType().(*TypeMessagesStickerSet), nil
	}
	return &TypeMessagesStickerSet{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) LangpackGetLangPack(ctx context.Context, req *ReqLangpackGetLangPack) (*TypeLangPackDifference, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredLangPackDifference:
		return x.ToType().(*TypeLangPackDifference), nil
	}
	return &TypeLangPackDifference{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) LangpackGetStrings(ctx context.Context, req *ReqLangpackGetStrings) (*TypeVectorLangPackString, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorLangPackString{}
		v.LangPackString = make([]*TypeLangPackString, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredLangPackString:
				v.LangPackString[i] = x.ToType().(*TypeLangPackString)
			case *PredLangPackStringPluralized:
				v.LangPackString[i] = x.ToType().(*TypeLangPackString)
			case *PredLangPackStringDeleted:
				v.LangPackString[i] = x.ToType().(*TypeLangPackString)
			default:
				return nil, fmt.Errorf("invalid LangPackString vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorLangPackString{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) LangpackGetDifference(ctx context.Context, req *ReqLangpackGetDifference) (*TypeLangPackDifference, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredLangPackDifference:
		return x.ToType().(*TypeLangPackDifference), nil
	}
	return &TypeLangPackDifference{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) LangpackGetLanguages(ctx context.Context, req *ReqLangpackGetLanguages) (*TypeVectorLangPackLanguage, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorLangPackLanguage{}
		v.LangPackLanguage = make([]*TypeLangPackLanguage, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredLangPackLanguage:
				v.LangPackLanguage[i] = x.ToType().(*TypeLangPackLanguage)
			default:
				return nil, fmt.Errorf("invalid LangPackLanguage vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorLangPackLanguage{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsEditBanned(ctx context.Context, req *ReqChannelsEditBanned) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsGetAdminLog(ctx context.Context, req *ReqChannelsGetAdminLog) (*TypeChannelsAdminLogResults, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredChannelsAdminLogResults:
		return x.ToType().(*TypeChannelsAdminLogResults), nil
	}
	return &TypeChannelsAdminLogResults{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) StickersRemoveStickerFromSet(ctx context.Context, req *ReqStickersRemoveStickerFromSet) (*TypeMessagesStickerSet, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSet:
		return x.ToType().(*TypeMessagesStickerSet), nil
	}
	return &TypeMessagesStickerSet{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) StickersChangeStickerPosition(ctx context.Context, req *ReqStickersChangeStickerPosition) (*TypeMessagesStickerSet, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSet:
		return x.ToType().(*TypeMessagesStickerSet), nil
	}
	return &TypeMessagesStickerSet{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) StickersAddStickerToSet(ctx context.Context, req *ReqStickersAddStickerToSet) (*TypeMessagesStickerSet, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesStickerSet:
		return x.ToType().(*TypeMessagesStickerSet), nil
	}
	return &TypeMessagesStickerSet{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesSendScreenshotNotification(ctx context.Context, req *ReqMessagesSendScreenshotNotification) (*TypeUpdates, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredUpdatesTooLong:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortChatMessage:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShort:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdatesCombined:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdates:
		return x.ToType().(*TypeUpdates), nil
	case *PredUpdateShortSentMessage:
		return x.ToType().(*TypeUpdates), nil
	}
	return &TypeUpdates{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) UploadGetCdnFileHashes(ctx context.Context, req *ReqUploadGetCdnFileHashes) (*TypeVectorCdnFileHash, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]TL); ok {
		v := &TypeVectorCdnFileHash{}
		v.CdnFileHash = make([]*TypeCdnFileHash, len(resp))
		for i, tl := range resp {
			switch x := tl.(type) {
			case *PredCdnFileHash:
				v.CdnFileHash[i] = x.ToType().(*TypeCdnFileHash)
			default:
				return nil, fmt.Errorf("invalid CdnFileHash vector element type:%T", resp)
			}
		}
		return v, nil
	}

	return &TypeVectorCdnFileHash{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetUnreadMentions(ctx context.Context, req *ReqMessagesGetUnreadMentions) (*TypeMessagesMessages, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesMessagesSlice:
		return x.ToType().(*TypeMessagesMessages), nil
	case *PredMessagesChannelMessages:
		return x.ToType().(*TypeMessagesMessages), nil
	}
	return &TypeMessagesMessages{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesFaveSticker(ctx context.Context, req *ReqMessagesFaveSticker) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsSetStickers(ctx context.Context, req *ReqChannelsSetStickers) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ContactsResetSaved(ctx context.Context, req *ReqContactsResetSaved) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) MessagesGetFavedStickers(ctx context.Context, req *ReqMessagesGetFavedStickers) (*TypeMessagesFavedStickers, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredMessagesFavedStickers:
		return x.ToType().(*TypeMessagesFavedStickers), nil
	case *PredMessagesFavedStickersNotModified:
		return x.ToType().(*TypeMessagesFavedStickers), nil
	}
	return &TypeMessagesFavedStickers{}, fmt.Errorf("unexpected return: %T", resp)
}
func (caller RPCaller) ChannelsReadMessageContents(ctx context.Context, req *ReqChannelsReadMessageContents) (*TypeBool, error) {
	resp, err := caller.RPC.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	switch x := resp.(type) {
	case *PredBoolFalse:
		return x.ToType().(*TypeBool), nil
	case *PredBoolTrue:
		return x.ToType().(*TypeBool), nil
	}
	return &TypeBool{}, fmt.Errorf("unexpected return: %T", resp)
}
