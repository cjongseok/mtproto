package mtproto

import (
	"io/ioutil"
	"os"
	"fmt"
	"github.com/golang/protobuf/proto"
)

// AccessManager keeps channel and user accesses
// Fetching the accesses from the Telegram on every app launching can cause over rate limits,
// if the account has lots of channels and users. AccessManager helps to manage theses accesses
// as json files. Use tools/access to export them.
type AccessManager struct {
	//accesses map[int32]Access
	channels map[int32]Access
	users map[int32]Access
}

type Access struct {
	ID       int32
	//Required bool
	Hash     int64
}

// NewAccessManager instantiate a AccessManager
// If chats or contacts are given, it returns the instance fulfilled
// with the given id and hashes. Empty, otherwise.
func NewAccessManager(chats *TypeMessagesChats, contacts *TypeContactsContacts) (hm *AccessManager) {
	hm = &AccessManager{make(map[int32]Access), make(map[int32]Access)}
	if chats != nil {
		hm.importTypeMessagesChats(chats)
	}
	if contacts != nil {
		hm.importTypeContactsContacts(contacts)
	}
	return
}

func (am *AccessManager) importTypeContactsContacts(t *TypeContactsContacts) {
	p := t.GetContactsContacts()
	if p != nil {
		for _, u := range p.Users {
			predUser := u.GetUser()
			if predUser != nil {
				am.users[predUser.Id] = Access{
					predUser.Id,
					predUser.AccessHash,
				}
			}
		}
	}
}
func typeChatToAccess(t *TypeChat) *Access {
	if t == nil {
		return nil
	}
	if channel := t.GetChannel(); channel != nil {
		return &Access{channel.Id, channel.AccessHash}
	} else if chanForbidden := t.GetChannelForbidden(); chanForbidden != nil {
		return &Access{chanForbidden.Id, chanForbidden.AccessHash}
	}

	return nil
}
func (am *AccessManager) importTypeMessagesChats(c *TypeMessagesChats) {
	if c == nil {
		return
	}
	//fmt.Println(slog.StringifyIndent(c, "  "))
	if chats := c.GetMessagesChats(); chats != nil {
		for _, chat := range chats.Chats {
			access := typeChatToAccess(chat)
			if access != nil {
				am.channels[access.ID] = *access
			}
		}
	} else if chatsSlice := c.GetMessagesChatsSlice(); chatsSlice != nil {
		for _, slice := range chatsSlice.Chats {
			access := typeChatToAccess(slice)
			if access != nil {
				am.channels[access.ID] = *access
			}
		}
	}
}

func (am *AccessManager) ImportUserAccessesFromFile(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 644)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	unmarshaled := &TypeContactsContacts{}
	err = proto.Unmarshal(b, unmarshaled)
	if err != nil {
		return err
	}
	am.importTypeContactsContacts(unmarshaled)
	return nil
}
func (am *AccessManager) ImportChanAccessesFromFile(filepath string) error {
	f, err := os.OpenFile(filepath, os.O_RDONLY, 644)
	if err != nil {
		return fmt.Errorf("open file failure; %v", err)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return fmt.Errorf("read file failure; %v", err)
	}
	unmarshaled := &TypeMessagesChats{}
	err = proto.Unmarshal(b, unmarshaled)
	if err != nil {
		return fmt.Errorf("unmarshal failure; %v", err)
	}
	am.importTypeMessagesChats(unmarshaled)
	return nil
}

func (am *AccessManager) ChannelAccess(id int32) Access {
	return am.channels[id]
}
func (am *AccessManager) UserAccess(id int32) Access {
	return am.users[id]
}
func (am *AccessManager) Users() []int32 {
	var ids []int32
	for id, _ := range am.users {
		ids = append(ids, id)
	}
	return ids
}
func (am *AccessManager) Channels() []int32 {
	var ids []int32
	for id, _ := range am.channels {
		ids = append(ids, id)
	}
	return ids
}
