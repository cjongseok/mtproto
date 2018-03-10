package proxy

import (
	"fmt"
	"github.com/cjongseok/mtproto/proto"
	"github.com/golang/protobuf/ptypes/any"
	"golang.org/x/net/context"
)

// Procedures
func (p *MProxy) InvokeAfterMsg(ctx context.Context, req *mtp.ReqInvokeAfterMsg) (*any.Any, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(mtp.TL); ok {
		packed := mtp.Pack(resp)
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T: %v", resp, resp)
}
func (p *MProxy) InvokeAfterMsgs(ctx context.Context, req *mtp.ReqInvokeAfterMsgs) (*any.Any, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(mtp.TL); ok {
		packed := mtp.Pack(resp)
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T: %v", resp, resp)
}
func (p *MProxy) AuthCheckPhone(ctx context.Context, req *mtp.ReqAuthCheckPhone) (*mtp.TypeAuthCheckedPhone, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthCheckedPhone); ok {
		return resp, nil
	}
	return &mtp.TypeAuthCheckedPhone{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthSendCode(ctx context.Context, req *mtp.ReqAuthSendCode) (*mtp.TypeAuthSentCode, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthSentCode); ok {
		return resp, nil
	}
	return &mtp.TypeAuthSentCode{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthSignUp(ctx context.Context, req *mtp.ReqAuthSignUp) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthSignIn(ctx context.Context, req *mtp.ReqAuthSignIn) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthLogOut(ctx context.Context, req *mtp.ReqAuthLogOut) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthResetAuthorizations(ctx context.Context, req *mtp.ReqAuthResetAuthorizations) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthSendInvites(ctx context.Context, req *mtp.ReqAuthSendInvites) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthExportAuthorization(ctx context.Context, req *mtp.ReqAuthExportAuthorization) (*mtp.TypeAuthExportedAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthExportedAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthExportedAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthImportAuthorization(ctx context.Context, req *mtp.ReqAuthImportAuthorization) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountRegisterDevice(ctx context.Context, req *mtp.ReqAccountRegisterDevice) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUnregisterDevice(ctx context.Context, req *mtp.ReqAccountUnregisterDevice) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdateNotifySettings(ctx context.Context, req *mtp.ReqAccountUpdateNotifySettings) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetNotifySettings(ctx context.Context, req *mtp.ReqAccountGetNotifySettings) (*mtp.TypePeerNotifySettings, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePeerNotifySettings); ok {
		return resp, nil
	}
	return &mtp.TypePeerNotifySettings{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountResetNotifySettings(ctx context.Context, req *mtp.ReqAccountResetNotifySettings) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdateProfile(ctx context.Context, req *mtp.ReqAccountUpdateProfile) (*mtp.TypeUser, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUser); ok {
		return resp, nil
	}
	return &mtp.TypeUser{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdateStatus(ctx context.Context, req *mtp.ReqAccountUpdateStatus) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetWallPapers(ctx context.Context, req *mtp.ReqAccountGetWallPapers) (*mtp.TypeVectorWallPaper, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorWallPaper{}
		v.WallPaper = make([]*mtp.TypeWallPaper, len(resp))
		for i, tl := range resp {
			v.WallPaper[i] = tl.(*mtp.TypeWallPaper)
		}
		return v, nil
	}
	return &mtp.TypeVectorWallPaper{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UsersGetUsers(ctx context.Context, req *mtp.ReqUsersGetUsers) (*mtp.TypeVectorUser, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorUser{}
		v.User = make([]*mtp.TypeUser, len(resp))
		for i, tl := range resp {
			v.User[i] = tl.(*mtp.TypeUser)
		}
		return v, nil
	}
	return &mtp.TypeVectorUser{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UsersGetFullUser(ctx context.Context, req *mtp.ReqUsersGetFullUser) (*mtp.TypeUserFull, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUserFull); ok {
		return resp, nil
	}
	return &mtp.TypeUserFull{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsGetStatuses(ctx context.Context, req *mtp.ReqContactsGetStatuses) (*mtp.TypeVectorContactStatus, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorContactStatus{}
		v.ContactStatus = make([]*mtp.TypeContactStatus, len(resp))
		for i, tl := range resp {
			v.ContactStatus[i] = tl.(*mtp.TypeContactStatus)
		}
		return v, nil
	}
	return &mtp.TypeVectorContactStatus{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsGetContacts(ctx context.Context, req *mtp.ReqContactsGetContacts) (*mtp.TypeContactsContacts, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsContacts); ok {
		return resp, nil
	}
	return &mtp.TypeContactsContacts{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsImportContacts(ctx context.Context, req *mtp.ReqContactsImportContacts) (*mtp.TypeContactsImportedContacts, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsImportedContacts); ok {
		return resp, nil
	}
	return &mtp.TypeContactsImportedContacts{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsSearch(ctx context.Context, req *mtp.ReqContactsSearch) (*mtp.TypeContactsFound, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsFound); ok {
		return resp, nil
	}
	return &mtp.TypeContactsFound{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsDeleteContact(ctx context.Context, req *mtp.ReqContactsDeleteContact) (*mtp.TypeContactsLink, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsLink); ok {
		return resp, nil
	}
	return &mtp.TypeContactsLink{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsDeleteContacts(ctx context.Context, req *mtp.ReqContactsDeleteContacts) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsBlock(ctx context.Context, req *mtp.ReqContactsBlock) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsUnblock(ctx context.Context, req *mtp.ReqContactsUnblock) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsGetBlocked(ctx context.Context, req *mtp.ReqContactsGetBlocked) (*mtp.TypeContactsBlocked, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsBlocked); ok {
		return resp, nil
	}
	return &mtp.TypeContactsBlocked{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetMessages(ctx context.Context, req *mtp.ReqMessagesGetMessages) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetDialogs(ctx context.Context, req *mtp.ReqMessagesGetDialogs) (*mtp.TypeMessagesDialogs, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesDialogs); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesDialogs{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetHistory(ctx context.Context, req *mtp.ReqMessagesGetHistory) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSearch(ctx context.Context, req *mtp.ReqMessagesSearch) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReadHistory(ctx context.Context, req *mtp.ReqMessagesReadHistory) (*mtp.TypeMessagesAffectedMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesDeleteHistory(ctx context.Context, req *mtp.ReqMessagesDeleteHistory) (*mtp.TypeMessagesAffectedHistory, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedHistory); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedHistory{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesDeleteMessages(ctx context.Context, req *mtp.ReqMessagesDeleteMessages) (*mtp.TypeMessagesAffectedMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReceivedMessages(ctx context.Context, req *mtp.ReqMessagesReceivedMessages) (*mtp.TypeVectorReceivedNotifyMessage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorReceivedNotifyMessage{}
		v.ReceivedNotifyMessage = make([]*mtp.TypeReceivedNotifyMessage, len(resp))
		for i, tl := range resp {
			v.ReceivedNotifyMessage[i] = tl.(*mtp.TypeReceivedNotifyMessage)
		}
		return v, nil
	}
	return &mtp.TypeVectorReceivedNotifyMessage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetTyping(ctx context.Context, req *mtp.ReqMessagesSetTyping) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendMessage(ctx context.Context, req *mtp.ReqMessagesSendMessage) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendMedia(ctx context.Context, req *mtp.ReqMessagesSendMedia) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesForwardMessages(ctx context.Context, req *mtp.ReqMessagesForwardMessages) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetChats(ctx context.Context, req *mtp.ReqMessagesGetChats) (*mtp.TypeMessagesChats, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChats); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChats{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetFullChat(ctx context.Context, req *mtp.ReqMessagesGetFullChat) (*mtp.TypeMessagesChatFull, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChatFull); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChatFull{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesEditChatTitle(ctx context.Context, req *mtp.ReqMessagesEditChatTitle) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesEditChatPhoto(ctx context.Context, req *mtp.ReqMessagesEditChatPhoto) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesAddChatUser(ctx context.Context, req *mtp.ReqMessagesAddChatUser) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesDeleteChatUser(ctx context.Context, req *mtp.ReqMessagesDeleteChatUser) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesCreateChat(ctx context.Context, req *mtp.ReqMessagesCreateChat) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UpdatesGetState(ctx context.Context, req *mtp.ReqUpdatesGetState) (*mtp.TypeUpdatesState, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdatesState); ok {
		return resp, nil
	}
	return &mtp.TypeUpdatesState{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UpdatesGetDifference(ctx context.Context, req *mtp.ReqUpdatesGetDifference) (*mtp.TypeUpdatesDifference, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdatesDifference); ok {
		return resp, nil
	}
	return &mtp.TypeUpdatesDifference{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhotosUpdateProfilePhoto(ctx context.Context, req *mtp.ReqPhotosUpdateProfilePhoto) (*mtp.TypeUserProfilePhoto, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUserProfilePhoto); ok {
		return resp, nil
	}
	return &mtp.TypeUserProfilePhoto{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhotosUploadProfilePhoto(ctx context.Context, req *mtp.ReqPhotosUploadProfilePhoto) (*mtp.TypePhotosPhoto, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePhotosPhoto); ok {
		return resp, nil
	}
	return &mtp.TypePhotosPhoto{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadSaveFilePart(ctx context.Context, req *mtp.ReqUploadSaveFilePart) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadGetFile(ctx context.Context, req *mtp.ReqUploadGetFile) (*mtp.TypeUploadFile, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUploadFile); ok {
		return resp, nil
	}
	return &mtp.TypeUploadFile{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetConfig(ctx context.Context, req *mtp.ReqHelpGetConfig) (*mtp.TypeConfig, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeConfig); ok {
		return resp, nil
	}
	return &mtp.TypeConfig{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetNearestDc(ctx context.Context, req *mtp.ReqHelpGetNearestDc) (*mtp.TypeNearestDc, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeNearestDc); ok {
		return resp, nil
	}
	return &mtp.TypeNearestDc{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetAppUpdate(ctx context.Context, req *mtp.ReqHelpGetAppUpdate) (*mtp.TypeHelpAppUpdate, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeHelpAppUpdate); ok {
		return resp, nil
	}
	return &mtp.TypeHelpAppUpdate{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpSaveAppLog(ctx context.Context, req *mtp.ReqHelpSaveAppLog) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetInviteText(ctx context.Context, req *mtp.ReqHelpGetInviteText) (*mtp.TypeHelpInviteText, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeHelpInviteText); ok {
		return resp, nil
	}
	return &mtp.TypeHelpInviteText{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhotosDeletePhotos(ctx context.Context, req *mtp.ReqPhotosDeletePhotos) (*mtp.TypeVectorLong, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int64); ok {
		return &mtp.TypeVectorLong{resp}, nil
	}
	return &mtp.TypeVectorLong{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhotosGetUserPhotos(ctx context.Context, req *mtp.ReqPhotosGetUserPhotos) (*mtp.TypePhotosPhotos, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePhotosPhotos); ok {
		return resp, nil
	}
	return &mtp.TypePhotosPhotos{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesForwardMessage(ctx context.Context, req *mtp.ReqMessagesForwardMessage) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetDhConfig(ctx context.Context, req *mtp.ReqMessagesGetDhConfig) (*mtp.TypeMessagesDhConfig, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesDhConfig); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesDhConfig{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesRequestEncryption(ctx context.Context, req *mtp.ReqMessagesRequestEncryption) (*mtp.TypeEncryptedChat, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeEncryptedChat); ok {
		return resp, nil
	}
	return &mtp.TypeEncryptedChat{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesAcceptEncryption(ctx context.Context, req *mtp.ReqMessagesAcceptEncryption) (*mtp.TypeEncryptedChat, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeEncryptedChat); ok {
		return resp, nil
	}
	return &mtp.TypeEncryptedChat{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesDiscardEncryption(ctx context.Context, req *mtp.ReqMessagesDiscardEncryption) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetEncryptedTyping(ctx context.Context, req *mtp.ReqMessagesSetEncryptedTyping) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReadEncryptedHistory(ctx context.Context, req *mtp.ReqMessagesReadEncryptedHistory) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendEncrypted(ctx context.Context, req *mtp.ReqMessagesSendEncrypted) (*mtp.TypeMessagesSentEncryptedMessage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesSentEncryptedMessage); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesSentEncryptedMessage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendEncryptedFile(ctx context.Context, req *mtp.ReqMessagesSendEncryptedFile) (*mtp.TypeMessagesSentEncryptedMessage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesSentEncryptedMessage); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesSentEncryptedMessage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendEncryptedService(ctx context.Context, req *mtp.ReqMessagesSendEncryptedService) (*mtp.TypeMessagesSentEncryptedMessage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesSentEncryptedMessage); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesSentEncryptedMessage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReceivedQueue(ctx context.Context, req *mtp.ReqMessagesReceivedQueue) (*mtp.TypeVectorLong, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int64); ok {
		return &mtp.TypeVectorLong{resp}, nil
	}
	return &mtp.TypeVectorLong{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadSaveBigFilePart(ctx context.Context, req *mtp.ReqUploadSaveBigFilePart) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) InitConnection(ctx context.Context, req *mtp.ReqInitConnection) (*any.Any, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(mtp.TL); ok {
		packed := mtp.Pack(resp)
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T: %v", resp, resp)
}
func (p *MProxy) HelpGetSupport(ctx context.Context, req *mtp.ReqHelpGetSupport) (*mtp.TypeHelpSupport, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeHelpSupport); ok {
		return resp, nil
	}
	return &mtp.TypeHelpSupport{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthBindTempAuthKey(ctx context.Context, req *mtp.ReqAuthBindTempAuthKey) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsExportCard(ctx context.Context, req *mtp.ReqContactsExportCard) (*mtp.TypeVectorInt, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int32); ok {
		return &mtp.TypeVectorInt{resp}, nil
	}
	return &mtp.TypeVectorInt{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsImportCard(ctx context.Context, req *mtp.ReqContactsImportCard) (*mtp.TypeUser, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUser); ok {
		return resp, nil
	}
	return &mtp.TypeUser{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReadMessageContents(ctx context.Context, req *mtp.ReqMessagesReadMessageContents) (*mtp.TypeMessagesAffectedMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountCheckUsername(ctx context.Context, req *mtp.ReqAccountCheckUsername) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdateUsername(ctx context.Context, req *mtp.ReqAccountUpdateUsername) (*mtp.TypeUser, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUser); ok {
		return resp, nil
	}
	return &mtp.TypeUser{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetPrivacy(ctx context.Context, req *mtp.ReqAccountGetPrivacy) (*mtp.TypeAccountPrivacyRules, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountPrivacyRules); ok {
		return resp, nil
	}
	return &mtp.TypeAccountPrivacyRules{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountSetPrivacy(ctx context.Context, req *mtp.ReqAccountSetPrivacy) (*mtp.TypeAccountPrivacyRules, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountPrivacyRules); ok {
		return resp, nil
	}
	return &mtp.TypeAccountPrivacyRules{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountDeleteAccount(ctx context.Context, req *mtp.ReqAccountDeleteAccount) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetAccountTTL(ctx context.Context, req *mtp.ReqAccountGetAccountTTL) (*mtp.TypeAccountDaysTTL, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountDaysTTL); ok {
		return resp, nil
	}
	return &mtp.TypeAccountDaysTTL{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountSetAccountTTL(ctx context.Context, req *mtp.ReqAccountSetAccountTTL) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) InvokeWithLayer(ctx context.Context, req *mtp.ReqInvokeWithLayer) (*any.Any, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(mtp.TL); ok {
		packed := mtp.Pack(resp)
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T: %v", resp, resp)
}
func (p *MProxy) ContactsResolveUsername(ctx context.Context, req *mtp.ReqContactsResolveUsername) (*mtp.TypeContactsResolvedPeer, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsResolvedPeer); ok {
		return resp, nil
	}
	return &mtp.TypeContactsResolvedPeer{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountSendChangePhoneCode(ctx context.Context, req *mtp.ReqAccountSendChangePhoneCode) (*mtp.TypeAuthSentCode, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthSentCode); ok {
		return resp, nil
	}
	return &mtp.TypeAuthSentCode{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountChangePhone(ctx context.Context, req *mtp.ReqAccountChangePhone) (*mtp.TypeUser, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUser); ok {
		return resp, nil
	}
	return &mtp.TypeUser{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetAllStickers(ctx context.Context, req *mtp.ReqMessagesGetAllStickers) (*mtp.TypeMessagesAllStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAllStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAllStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdateDeviceLocked(ctx context.Context, req *mtp.ReqAccountUpdateDeviceLocked) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetPassword(ctx context.Context, req *mtp.ReqAccountGetPassword) (*mtp.TypeAccountPassword, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountPassword); ok {
		return resp, nil
	}
	return &mtp.TypeAccountPassword{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthCheckPassword(ctx context.Context, req *mtp.ReqAuthCheckPassword) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetWebPagePreview(ctx context.Context, req *mtp.ReqMessagesGetWebPagePreview) (*mtp.TypeMessageMedia, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessageMedia); ok {
		return resp, nil
	}
	return &mtp.TypeMessageMedia{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetAuthorizations(ctx context.Context, req *mtp.ReqAccountGetAuthorizations) (*mtp.TypeAccountAuthorizations, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountAuthorizations); ok {
		return resp, nil
	}
	return &mtp.TypeAccountAuthorizations{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountResetAuthorization(ctx context.Context, req *mtp.ReqAccountResetAuthorization) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetPasswordSettings(ctx context.Context, req *mtp.ReqAccountGetPasswordSettings) (*mtp.TypeAccountPasswordSettings, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountPasswordSettings); ok {
		return resp, nil
	}
	return &mtp.TypeAccountPasswordSettings{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountUpdatePasswordSettings(ctx context.Context, req *mtp.ReqAccountUpdatePasswordSettings) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthRequestPasswordRecovery(ctx context.Context, req *mtp.ReqAuthRequestPasswordRecovery) (*mtp.TypeAuthPasswordRecovery, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthPasswordRecovery); ok {
		return resp, nil
	}
	return &mtp.TypeAuthPasswordRecovery{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthRecoverPassword(ctx context.Context, req *mtp.ReqAuthRecoverPassword) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) InvokeWithoutUpdates(ctx context.Context, req *mtp.ReqInvokeWithoutUpdates) (*any.Any, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(mtp.TL); ok {
		packed := mtp.Pack(resp)
		if packed != nil {
			return packed, nil
		}
	}
	return &any.Any{}, fmt.Errorf("unexpected return: %T: %v", resp, resp)
}
func (p *MProxy) MessagesExportChatInvite(ctx context.Context, req *mtp.ReqMessagesExportChatInvite) (*mtp.TypeExportedChatInvite, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeExportedChatInvite); ok {
		return resp, nil
	}
	return &mtp.TypeExportedChatInvite{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesCheckChatInvite(ctx context.Context, req *mtp.ReqMessagesCheckChatInvite) (*mtp.TypeChatInvite, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeChatInvite); ok {
		return resp, nil
	}
	return &mtp.TypeChatInvite{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesImportChatInvite(ctx context.Context, req *mtp.ReqMessagesImportChatInvite) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetStickerSet(ctx context.Context, req *mtp.ReqMessagesGetStickerSet) (*mtp.TypeMessagesStickerSet, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSet); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSet{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesInstallStickerSet(ctx context.Context, req *mtp.ReqMessagesInstallStickerSet) (*mtp.TypeMessagesStickerSetInstallResult, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSetInstallResult); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSetInstallResult{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesUninstallStickerSet(ctx context.Context, req *mtp.ReqMessagesUninstallStickerSet) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthImportBotAuthorization(ctx context.Context, req *mtp.ReqAuthImportBotAuthorization) (*mtp.TypeAuthAuthorization, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthAuthorization); ok {
		return resp, nil
	}
	return &mtp.TypeAuthAuthorization{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesStartBot(ctx context.Context, req *mtp.ReqMessagesStartBot) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetAppChangelog(ctx context.Context, req *mtp.ReqHelpGetAppChangelog) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReportSpam(ctx context.Context, req *mtp.ReqMessagesReportSpam) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetMessagesViews(ctx context.Context, req *mtp.ReqMessagesGetMessagesViews) (*mtp.TypeVectorInt, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]int32); ok {
		return &mtp.TypeVectorInt{resp}, nil
	}
	return &mtp.TypeVectorInt{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UpdatesGetChannelDifference(ctx context.Context, req *mtp.ReqUpdatesGetChannelDifference) (*mtp.TypeUpdatesChannelDifference, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdatesChannelDifference); ok {
		return resp, nil
	}
	return &mtp.TypeUpdatesChannelDifference{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsReadHistory(ctx context.Context, req *mtp.ReqChannelsReadHistory) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsDeleteMessages(ctx context.Context, req *mtp.ReqChannelsDeleteMessages) (*mtp.TypeMessagesAffectedMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsDeleteUserHistory(ctx context.Context, req *mtp.ReqChannelsDeleteUserHistory) (*mtp.TypeMessagesAffectedHistory, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAffectedHistory); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAffectedHistory{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsReportSpam(ctx context.Context, req *mtp.ReqChannelsReportSpam) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetMessages(ctx context.Context, req *mtp.ReqChannelsGetMessages) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetParticipants(ctx context.Context, req *mtp.ReqChannelsGetParticipants) (*mtp.TypeChannelsChannelParticipants, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeChannelsChannelParticipants); ok {
		return resp, nil
	}
	return &mtp.TypeChannelsChannelParticipants{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetParticipant(ctx context.Context, req *mtp.ReqChannelsGetParticipant) (*mtp.TypeChannelsChannelParticipant, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeChannelsChannelParticipant); ok {
		return resp, nil
	}
	return &mtp.TypeChannelsChannelParticipant{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetChannels(ctx context.Context, req *mtp.ReqChannelsGetChannels) (*mtp.TypeMessagesChats, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChats); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChats{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetFullChannel(ctx context.Context, req *mtp.ReqChannelsGetFullChannel) (*mtp.TypeMessagesChatFull, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChatFull); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChatFull{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsCreateChannel(ctx context.Context, req *mtp.ReqChannelsCreateChannel) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsEditAbout(ctx context.Context, req *mtp.ReqChannelsEditAbout) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsEditAdmin(ctx context.Context, req *mtp.ReqChannelsEditAdmin) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsEditTitle(ctx context.Context, req *mtp.ReqChannelsEditTitle) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsEditPhoto(ctx context.Context, req *mtp.ReqChannelsEditPhoto) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsCheckUsername(ctx context.Context, req *mtp.ReqChannelsCheckUsername) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsUpdateUsername(ctx context.Context, req *mtp.ReqChannelsUpdateUsername) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsJoinChannel(ctx context.Context, req *mtp.ReqChannelsJoinChannel) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsLeaveChannel(ctx context.Context, req *mtp.ReqChannelsLeaveChannel) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsInviteToChannel(ctx context.Context, req *mtp.ReqChannelsInviteToChannel) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsExportInvite(ctx context.Context, req *mtp.ReqChannelsExportInvite) (*mtp.TypeExportedChatInvite, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeExportedChatInvite); ok {
		return resp, nil
	}
	return &mtp.TypeExportedChatInvite{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsDeleteChannel(ctx context.Context, req *mtp.ReqChannelsDeleteChannel) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesToggleChatAdmins(ctx context.Context, req *mtp.ReqMessagesToggleChatAdmins) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesEditChatAdmin(ctx context.Context, req *mtp.ReqMessagesEditChatAdmin) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesMigrateChat(ctx context.Context, req *mtp.ReqMessagesMigrateChat) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSearchGlobal(ctx context.Context, req *mtp.ReqMessagesSearchGlobal) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountReportPeer(ctx context.Context, req *mtp.ReqAccountReportPeer) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReorderStickerSets(ctx context.Context, req *mtp.ReqMessagesReorderStickerSets) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetTermsOfService(ctx context.Context, req *mtp.ReqHelpGetTermsOfService) (*mtp.TypeHelpTermsOfService, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeHelpTermsOfService); ok {
		return resp, nil
	}
	return &mtp.TypeHelpTermsOfService{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetDocumentByHash(ctx context.Context, req *mtp.ReqMessagesGetDocumentByHash) (*mtp.TypeDocument, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeDocument); ok {
		return resp, nil
	}
	return &mtp.TypeDocument{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSearchGifs(ctx context.Context, req *mtp.ReqMessagesSearchGifs) (*mtp.TypeMessagesFoundGifs, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesFoundGifs); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesFoundGifs{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetSavedGifs(ctx context.Context, req *mtp.ReqMessagesGetSavedGifs) (*mtp.TypeMessagesSavedGifs, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesSavedGifs); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesSavedGifs{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSaveGif(ctx context.Context, req *mtp.ReqMessagesSaveGif) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetInlineBotResults(ctx context.Context, req *mtp.ReqMessagesGetInlineBotResults) (*mtp.TypeMessagesBotResults, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesBotResults); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesBotResults{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetInlineBotResults(ctx context.Context, req *mtp.ReqMessagesSetInlineBotResults) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendInlineBotResult(ctx context.Context, req *mtp.ReqMessagesSendInlineBotResult) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsToggleInvites(ctx context.Context, req *mtp.ReqChannelsToggleInvites) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsExportMessageLink(ctx context.Context, req *mtp.ReqChannelsExportMessageLink) (*mtp.TypeExportedMessageLink, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeExportedMessageLink); ok {
		return resp, nil
	}
	return &mtp.TypeExportedMessageLink{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsToggleSignatures(ctx context.Context, req *mtp.ReqChannelsToggleSignatures) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesHideReportSpam(ctx context.Context, req *mtp.ReqMessagesHideReportSpam) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetPeerSettings(ctx context.Context, req *mtp.ReqMessagesGetPeerSettings) (*mtp.TypePeerSettings, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePeerSettings); ok {
		return resp, nil
	}
	return &mtp.TypePeerSettings{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsUpdatePinnedMessage(ctx context.Context, req *mtp.ReqChannelsUpdatePinnedMessage) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthResendCode(ctx context.Context, req *mtp.ReqAuthResendCode) (*mtp.TypeAuthSentCode, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthSentCode); ok {
		return resp, nil
	}
	return &mtp.TypeAuthSentCode{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthCancelCode(ctx context.Context, req *mtp.ReqAuthCancelCode) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetMessageEditData(ctx context.Context, req *mtp.ReqMessagesGetMessageEditData) (*mtp.TypeMessagesMessageEditData, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessageEditData); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessageEditData{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesEditMessage(ctx context.Context, req *mtp.ReqMessagesEditMessage) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesEditInlineBotMessage(ctx context.Context, req *mtp.ReqMessagesEditInlineBotMessage) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetBotCallbackAnswer(ctx context.Context, req *mtp.ReqMessagesGetBotCallbackAnswer) (*mtp.TypeMessagesBotCallbackAnswer, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesBotCallbackAnswer); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesBotCallbackAnswer{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetBotCallbackAnswer(ctx context.Context, req *mtp.ReqMessagesSetBotCallbackAnswer) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsGetTopPeers(ctx context.Context, req *mtp.ReqContactsGetTopPeers) (*mtp.TypeContactsTopPeers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeContactsTopPeers); ok {
		return resp, nil
	}
	return &mtp.TypeContactsTopPeers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsResetTopPeerRating(ctx context.Context, req *mtp.ReqContactsResetTopPeerRating) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetPeerDialogs(ctx context.Context, req *mtp.ReqMessagesGetPeerDialogs) (*mtp.TypeMessagesPeerDialogs, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesPeerDialogs); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesPeerDialogs{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSaveDraft(ctx context.Context, req *mtp.ReqMessagesSaveDraft) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetAllDrafts(ctx context.Context, req *mtp.ReqMessagesGetAllDrafts) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountSendConfirmPhoneCode(ctx context.Context, req *mtp.ReqAccountSendConfirmPhoneCode) (*mtp.TypeAuthSentCode, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAuthSentCode); ok {
		return resp, nil
	}
	return &mtp.TypeAuthSentCode{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountConfirmPhone(ctx context.Context, req *mtp.ReqAccountConfirmPhone) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetFeaturedStickers(ctx context.Context, req *mtp.ReqMessagesGetFeaturedStickers) (*mtp.TypeMessagesFeaturedStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesFeaturedStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesFeaturedStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReadFeaturedStickers(ctx context.Context, req *mtp.ReqMessagesReadFeaturedStickers) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetRecentStickers(ctx context.Context, req *mtp.ReqMessagesGetRecentStickers) (*mtp.TypeMessagesRecentStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesRecentStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesRecentStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSaveRecentSticker(ctx context.Context, req *mtp.ReqMessagesSaveRecentSticker) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesClearRecentStickers(ctx context.Context, req *mtp.ReqMessagesClearRecentStickers) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetArchivedStickers(ctx context.Context, req *mtp.ReqMessagesGetArchivedStickers) (*mtp.TypeMessagesArchivedStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesArchivedStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesArchivedStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetAdminedPublicChannels(ctx context.Context, req *mtp.ReqChannelsGetAdminedPublicChannels) (*mtp.TypeMessagesChats, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChats); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChats{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AuthDropTempAuthKeys(ctx context.Context, req *mtp.ReqAuthDropTempAuthKeys) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetGameScore(ctx context.Context, req *mtp.ReqMessagesSetGameScore) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetInlineGameScore(ctx context.Context, req *mtp.ReqMessagesSetInlineGameScore) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetMaskStickers(ctx context.Context, req *mtp.ReqMessagesGetMaskStickers) (*mtp.TypeMessagesAllStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesAllStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesAllStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetAttachedStickers(ctx context.Context, req *mtp.ReqMessagesGetAttachedStickers) (*mtp.TypeVectorStickerSetCovered, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorStickerSetCovered{}
		v.StickerSetCovered = make([]*mtp.TypeStickerSetCovered, len(resp))
		for i, tl := range resp {
			v.StickerSetCovered[i] = tl.(*mtp.TypeStickerSetCovered)
		}
		return v, nil
	}
	return &mtp.TypeVectorStickerSetCovered{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetGameHighScores(ctx context.Context, req *mtp.ReqMessagesGetGameHighScores) (*mtp.TypeMessagesHighScores, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesHighScores); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesHighScores{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetInlineGameHighScores(ctx context.Context, req *mtp.ReqMessagesGetInlineGameHighScores) (*mtp.TypeMessagesHighScores, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesHighScores); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesHighScores{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetCommonChats(ctx context.Context, req *mtp.ReqMessagesGetCommonChats) (*mtp.TypeMessagesChats, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChats); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChats{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetAllChats(ctx context.Context, req *mtp.ReqMessagesGetAllChats) (*mtp.TypeMessagesChats, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesChats); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesChats{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpSetBotUpdatesStatus(ctx context.Context, req *mtp.ReqHelpSetBotUpdatesStatus) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetWebPage(ctx context.Context, req *mtp.ReqMessagesGetWebPage) (*mtp.TypeWebPage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeWebPage); ok {
		return resp, nil
	}
	return &mtp.TypeWebPage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesToggleDialogPin(ctx context.Context, req *mtp.ReqMessagesToggleDialogPin) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReorderPinnedDialogs(ctx context.Context, req *mtp.ReqMessagesReorderPinnedDialogs) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetPinnedDialogs(ctx context.Context, req *mtp.ReqMessagesGetPinnedDialogs) (*mtp.TypeMessagesPeerDialogs, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesPeerDialogs); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesPeerDialogs{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneRequestCall(ctx context.Context, req *mtp.ReqPhoneRequestCall) (*mtp.TypePhonePhoneCall, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePhonePhoneCall); ok {
		return resp, nil
	}
	return &mtp.TypePhonePhoneCall{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneAcceptCall(ctx context.Context, req *mtp.ReqPhoneAcceptCall) (*mtp.TypePhonePhoneCall, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePhonePhoneCall); ok {
		return resp, nil
	}
	return &mtp.TypePhonePhoneCall{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneDiscardCall(ctx context.Context, req *mtp.ReqPhoneDiscardCall) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneReceivedCall(ctx context.Context, req *mtp.ReqPhoneReceivedCall) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesReportEncryptedSpam(ctx context.Context, req *mtp.ReqMessagesReportEncryptedSpam) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsGetPaymentForm(ctx context.Context, req *mtp.ReqPaymentsGetPaymentForm) (*mtp.TypePaymentsPaymentForm, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePaymentsPaymentForm); ok {
		return resp, nil
	}
	return &mtp.TypePaymentsPaymentForm{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsSendPaymentForm(ctx context.Context, req *mtp.ReqPaymentsSendPaymentForm) (*mtp.TypePaymentsPaymentResult, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePaymentsPaymentResult); ok {
		return resp, nil
	}
	return &mtp.TypePaymentsPaymentResult{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) AccountGetTmpPassword(ctx context.Context, req *mtp.ReqAccountGetTmpPassword) (*mtp.TypeAccountTmpPassword, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeAccountTmpPassword); ok {
		return resp, nil
	}
	return &mtp.TypeAccountTmpPassword{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetBotShippingResults(ctx context.Context, req *mtp.ReqMessagesSetBotShippingResults) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSetBotPrecheckoutResults(ctx context.Context, req *mtp.ReqMessagesSetBotPrecheckoutResults) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadGetWebFile(ctx context.Context, req *mtp.ReqUploadGetWebFile) (*mtp.TypeUploadWebFile, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUploadWebFile); ok {
		return resp, nil
	}
	return &mtp.TypeUploadWebFile{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) BotsSendCustomRequest(ctx context.Context, req *mtp.ReqBotsSendCustomRequest) (*mtp.TypeDataJSON, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeDataJSON); ok {
		return resp, nil
	}
	return &mtp.TypeDataJSON{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) BotsAnswerWebhookJSONQuery(ctx context.Context, req *mtp.ReqBotsAnswerWebhookJSONQuery) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsGetPaymentReceipt(ctx context.Context, req *mtp.ReqPaymentsGetPaymentReceipt) (*mtp.TypePaymentsPaymentReceipt, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePaymentsPaymentReceipt); ok {
		return resp, nil
	}
	return &mtp.TypePaymentsPaymentReceipt{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsValidateRequestedInfo(ctx context.Context, req *mtp.ReqPaymentsValidateRequestedInfo) (*mtp.TypePaymentsValidatedRequestedInfo, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePaymentsValidatedRequestedInfo); ok {
		return resp, nil
	}
	return &mtp.TypePaymentsValidatedRequestedInfo{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsGetSavedInfo(ctx context.Context, req *mtp.ReqPaymentsGetSavedInfo) (*mtp.TypePaymentsSavedInfo, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePaymentsSavedInfo); ok {
		return resp, nil
	}
	return &mtp.TypePaymentsSavedInfo{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PaymentsClearSavedInfo(ctx context.Context, req *mtp.ReqPaymentsClearSavedInfo) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneGetCallConfig(ctx context.Context, req *mtp.ReqPhoneGetCallConfig) (*mtp.TypeDataJSON, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeDataJSON); ok {
		return resp, nil
	}
	return &mtp.TypeDataJSON{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneConfirmCall(ctx context.Context, req *mtp.ReqPhoneConfirmCall) (*mtp.TypePhonePhoneCall, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypePhonePhoneCall); ok {
		return resp, nil
	}
	return &mtp.TypePhonePhoneCall{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneSetCallRating(ctx context.Context, req *mtp.ReqPhoneSetCallRating) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) PhoneSaveCallDebug(ctx context.Context, req *mtp.ReqPhoneSaveCallDebug) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadGetCdnFile(ctx context.Context, req *mtp.ReqUploadGetCdnFile) (*mtp.TypeUploadCdnFile, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUploadCdnFile); ok {
		return resp, nil
	}
	return &mtp.TypeUploadCdnFile{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadReuploadCdnFile(ctx context.Context, req *mtp.ReqUploadReuploadCdnFile) (*mtp.TypeVectorCdnFileHash, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorCdnFileHash{}
		v.CdnFileHash = make([]*mtp.TypeCdnFileHash, len(resp))
		for i, tl := range resp {
			v.CdnFileHash[i] = tl.(*mtp.TypeCdnFileHash)
		}
		return v, nil
	}
	return &mtp.TypeVectorCdnFileHash{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) HelpGetCdnConfig(ctx context.Context, req *mtp.ReqHelpGetCdnConfig) (*mtp.TypeCdnConfig, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeCdnConfig); ok {
		return resp, nil
	}
	return &mtp.TypeCdnConfig{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesUploadMedia(ctx context.Context, req *mtp.ReqMessagesUploadMedia) (*mtp.TypeMessageMedia, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessageMedia); ok {
		return resp, nil
	}
	return &mtp.TypeMessageMedia{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) StickersCreateStickerSet(ctx context.Context, req *mtp.ReqStickersCreateStickerSet) (*mtp.TypeMessagesStickerSet, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSet); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSet{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) LangpackGetLangPack(ctx context.Context, req *mtp.ReqLangpackGetLangPack) (*mtp.TypeLangPackDifference, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeLangPackDifference); ok {
		return resp, nil
	}
	return &mtp.TypeLangPackDifference{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) LangpackGetStrings(ctx context.Context, req *mtp.ReqLangpackGetStrings) (*mtp.TypeVectorLangPackString, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorLangPackString{}
		v.LangPackString = make([]*mtp.TypeLangPackString, len(resp))
		for i, tl := range resp {
			v.LangPackString[i] = tl.(*mtp.TypeLangPackString)
		}
		return v, nil
	}
	return &mtp.TypeVectorLangPackString{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) LangpackGetDifference(ctx context.Context, req *mtp.ReqLangpackGetDifference) (*mtp.TypeLangPackDifference, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeLangPackDifference); ok {
		return resp, nil
	}
	return &mtp.TypeLangPackDifference{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) LangpackGetLanguages(ctx context.Context, req *mtp.ReqLangpackGetLanguages) (*mtp.TypeVectorLangPackLanguage, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorLangPackLanguage{}
		v.LangPackLanguage = make([]*mtp.TypeLangPackLanguage, len(resp))
		for i, tl := range resp {
			v.LangPackLanguage[i] = tl.(*mtp.TypeLangPackLanguage)
		}
		return v, nil
	}
	return &mtp.TypeVectorLangPackLanguage{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsEditBanned(ctx context.Context, req *mtp.ReqChannelsEditBanned) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsGetAdminLog(ctx context.Context, req *mtp.ReqChannelsGetAdminLog) (*mtp.TypeChannelsAdminLogResults, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeChannelsAdminLogResults); ok {
		return resp, nil
	}
	return &mtp.TypeChannelsAdminLogResults{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) StickersRemoveStickerFromSet(ctx context.Context, req *mtp.ReqStickersRemoveStickerFromSet) (*mtp.TypeMessagesStickerSet, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSet); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSet{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) StickersChangeStickerPosition(ctx context.Context, req *mtp.ReqStickersChangeStickerPosition) (*mtp.TypeMessagesStickerSet, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSet); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSet{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) StickersAddStickerToSet(ctx context.Context, req *mtp.ReqStickersAddStickerToSet) (*mtp.TypeMessagesStickerSet, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesStickerSet); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesStickerSet{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesSendScreenshotNotification(ctx context.Context, req *mtp.ReqMessagesSendScreenshotNotification) (*mtp.TypeUpdates, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeUpdates); ok {
		return resp, nil
	}
	return &mtp.TypeUpdates{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) UploadGetCdnFileHashes(ctx context.Context, req *mtp.ReqUploadGetCdnFileHashes) (*mtp.TypeVectorCdnFileHash, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.([]mtp.TL); ok {
		v := &mtp.TypeVectorCdnFileHash{}
		v.CdnFileHash = make([]*mtp.TypeCdnFileHash, len(resp))
		for i, tl := range resp {
			v.CdnFileHash[i] = tl.(*mtp.TypeCdnFileHash)
		}
		return v, nil
	}
	return &mtp.TypeVectorCdnFileHash{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetUnreadMentions(ctx context.Context, req *mtp.ReqMessagesGetUnreadMentions) (*mtp.TypeMessagesMessages, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesMessages); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesMessages{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesFaveSticker(ctx context.Context, req *mtp.ReqMessagesFaveSticker) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsSetStickers(ctx context.Context, req *mtp.ReqChannelsSetStickers) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ContactsResetSaved(ctx context.Context, req *mtp.ReqContactsResetSaved) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) MessagesGetFavedStickers(ctx context.Context, req *mtp.ReqMessagesGetFavedStickers) (*mtp.TypeMessagesFavedStickers, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeMessagesFavedStickers); ok {
		return resp, nil
	}
	return &mtp.TypeMessagesFavedStickers{}, fmt.Errorf("return type is %T", resp)
}
func (p *MProxy) ChannelsReadMessageContents(ctx context.Context, req *mtp.ReqChannelsReadMessageContents) (*mtp.TypeBool, error) {
	resp, err := p.mconn.InvokeBlocked(req)
	if err != nil {
		return nil, err
	}
	if resp, ok := resp.(*mtp.TypeBool); ok {
		return resp, nil
	}
	return &mtp.TypeBool{}, fmt.Errorf("return type is %T", resp)
}
