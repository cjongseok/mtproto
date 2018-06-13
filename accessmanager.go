package mtproto

import (
	"os"
	"encoding/json"
)

type AccessManager struct {
	accesses map[int32]Access
}

type Access struct {
	ID int32
	Required bool
	Hash int64
}

func NewAccessManager(allChats *TypeMessagesChats) (hm *AccessManager) {
	hm = &AccessManager{make(map[int32]Access)}
	if allChats == nil {
		return
	}
	hm.importTypeMessagesChats(allChats)
	return
}

func typeChatToAccess(t *TypeChat) *Access {
	if t == nil {
		return nil
	}
	if channel := t.GetChannel(); channel != nil {
		return &Access{channel.Id, true, channel.AccessHash}
	} else if chanForbidden := t.GetChannelForbidden(); chanForbidden != nil {
		return &Access{chanForbidden.Id, true, chanForbidden.AccessHash,}
	} else if chat := t.GetChat(); chat != nil {
		return &Access{chat.Id, false, 0}
	} else if chatEmpty := t.GetChatEmpty(); chatEmpty != nil {
		return &Access{chatEmpty.Id, false, 0}
	} else if chatForbidden := t.GetChatForbidden(); chatForbidden != nil {
		return &Access{chatForbidden.Id, false, 0}
	}
	return nil
}
func (am *AccessManager) importTypeMessagesChats(c *TypeMessagesChats) {
	if c == nil {
		return
	}
	if chats := c.GetMessagesChats(); chats != nil {
		for _, chat := range chats.Chats {
			access := typeChatToAccess(chat)
			if access != nil {
				am.accesses[access.ID] = *access
			}
		}
	} else if chatsSlice := c.GetMessagesChatsSlice(); chatsSlice != nil {
		for _, slice := range chatsSlice.Chats {
			access := typeChatToAccess(slice)
			if access != nil {
				am.accesses[access.ID] = *access
			}
		}
	}
}

func (am *AccessManager) ImportFromFile(credentialsFile string) error {
	f, err := os.OpenFile(credentialsFile, os.O_RDONLY, 644)
	if err != nil {
		return err
	}
	b := make([]byte, 1024*4)
	n, err := f.ReadAt(b, 0)
	if n <= 0 || (err != nil && err.Error() != "EOF") {
	}

	unmarshaled := &TypeMessagesChats{}
	err = json.Unmarshal(b[:n], unmarshaled)
	if err != nil {
		return err
	}
	am.importTypeMessagesChats(unmarshaled)
	return nil
}

func (am *AccessManager) Access(id int32) Access {
	return am.accesses[id]
}

func (am *AccessManager) IDs() []int32 {
	var ids []int32
	for id, _ := range am.accesses {
		ids = append(ids, id)
	}
	return ids
}


