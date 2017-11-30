package mtproto

type unstripped interface{
	Strip() unstrip
}

type unstrip interface {
	Unstrip() unstripped	// CAVEAT: it make a copy of slice fields
	UnstrippedString() string
}

// unstripped types
type US_messages_dialogsSlice struct {
	Count 		int32
	Dialogs 	[]TL_dialog
	Messages 	[]TL_message
	//Chats 		[]TL_chat
	Channels 	[]TL_channel
	Users 		[]TL_user
}

type US_contacts_contacts struct {
	Contacts 	[]TL_contact
	Users 		[]TL_user
}

type US_contacts_topPeers struct {
	Categories 	[]TL_topPeerCategoryPeers
	Channels	[]TL_channel
	Users 		[]TL_user
}

type US_topPeerCategoryPeers struct {
	Cetegory 	TL
	Count 		int32
	Peers 		[]TL_topPeer
}

type US_messages_chats struct {
	//Chats 		[]TL_chat	// XXX: TL_chat is DEPRECATED?
	Channels 	[]TL_channel
}

type US_updates_difference struct {
	New_messages 			[]TL_message
	New_encrypted_messages 	[]TL_encryptedMessage
	Other_updates 			[]TL
	//Chats 					[]TL_chat 	// XXX: TL_chat is DEPRECATED
	Channels 				[]TL_channel
	Users 					[]TL_user
	State 					TL_updates_state
}

// interface implementations
func Unstrip(tl TL) (unstripped, bool) {
	switch tl.(type) {
	case TL_messages_dialogsSlice:
		return tl.(TL_messages_dialogsSlice).Unstrip(), true
	default:
		return nil, false
	}
}

func (us US_messages_dialogsSlice) Strip() unstrip {return nil}
func (tl TL_messages_dialogsSlice) Unstrip() unstripped {
	return US_messages_dialogsSlice{
		tl.Count,
		unstripTLdialogs(tl.Dialogs),
		unstripTLmsgs(tl.Messages),
		//unstripTLchats(tl.Chats),
		unstripTLchannels(tl.Chats),
		unstripTLusers(tl.Users),
	}
}


func (us US_contacts_contacts) Strip() unstrip {return nil}
func (tl TL_contacts_contacts) Unstrip() unstripped {
	return US_contacts_contacts{
		unstripTLcontacts(tl.Contacts),
		unstripTLusers(tl.Users),
	}
}

func (us US_contacts_topPeers) Strip() unstrip {return nil}
func (tl TL_contacts_topPeers) Unstrip() unstripped {
	return US_contacts_topPeers{
		unstripTLtopPeerCategoryPeerses(tl.Categories)	,
		//unstripTLchats(tl.Chats),		// XXX: TL_chat is DEPRECATED
		unstripTLchannels(tl.Chats),
		unstripTLusers(tl.Users),
	}
}

func (us US_topPeerCategoryPeers) Strip() unstrip {return nil}
func (tl TL_topPeerCategoryPeers) Unstrip() unstripped {
	return US_topPeerCategoryPeers{
		tl.Category,
		tl.Count,
		unstripTLtopPeers(tl.Peers),
	}
}

func (us US_messages_chats) Strip() unstrip {return nil}
func (tl TL_messages_chats) Unstrip() unstripped {
	return US_messages_chats{
		//unstripTLchats(tl.Chats),		// XXX: TL_chat is DEPRECATED?
		unstripTLchannels(tl.Chats),
	}
}

func (us US_updates_difference) Strip() unstrip {return nil}
func (tl TL_updates_difference) Unstrip() unstripped {
	return US_updates_difference{
		unstripTLmsgs(tl.New_messages),
		unstripTLencryptedMsg(tl.New_encrypted_messages),
		tl.Other_updates,
		//unstripTLchats(tl.Chats), 	// XXX: TL_chat is DEPRECATED
		unstripTLchannels(tl.Chats),
		unstripTLusers(tl.Users),
		tl.State.(TL_updates_state),
	}
}

// utility functions
func unstripTLdialogs(tls []TL) []TL_dialog {
	dialogs := make([]TL_dialog, len(tls))
	for i, tl := range tls {
		dialogs[i] = tl.(TL_dialog)
	}
	return dialogs
}
func unstripTLmsgs(tls []TL) []TL_message {
	msgs := make([]TL_message, len(tls))
	for i, tl := range tls {
		msgs[i] = tl.(TL_message)
	}
	return msgs
}
func unstripTLchannels(tls []TL) []TL_channel {
	channels := make([]TL_channel, len(tls))
	for i, tl := range tls {
		channels[i] = tl.(TL_channel)
	}
	return channels
}
func unstripTLchats(tls []TL) []TL_chat {
	chats := make([]TL_chat, len(tls))
	for i, tl := range tls {
		chats[i] = tl.(TL_chat)
	}
	return chats
}
func unstripTLusers(tls []TL) []TL_user {
	users := make([]TL_user, len(tls))
	for i, tl := range tls {
		users[i] = tl.(TL_user)
	}
	return users;
}
func unstripTLcontacts(tls []TL) []TL_contact {
	contacts := make([]TL_contact, len(tls))
	for i, tl := range tls {
		contacts[i] = tl.(TL_contact)
	}
	return contacts
}
func unstripTLtopPeerCategoryPeerses(tls []TL) []TL_topPeerCategoryPeers {
	topPeerCategoryPeerses := make([]TL_topPeerCategoryPeers, len(tls))
	for i, tl := range tls {
		topPeerCategoryPeerses[i] = tl.(TL_topPeerCategoryPeers)
	}
	return topPeerCategoryPeerses
}
func unstripTLtopPeers(tls []TL) []TL_topPeer {
	topPeers := make([]TL_topPeer, len(tls))
	for i, tl := range tls {
		topPeers[i] = tl.(TL_topPeer)
	}
	return topPeers
}
func unstripTLencryptedMsg(tls []TL) []TL_encryptedMessage {
	ems := make([]TL_encryptedMessage, len(tls))
	for i, tl := range tls {
		ems[i] = tl.(TL_encryptedMessage)
	}
	return ems
}

