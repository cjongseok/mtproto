package mtproto

import (
	"math/rand"
)

func (mconn *MConn) MessagesGetHistory(peer TL, offsetId, offsetDate, addOffset, limit, maxId, minId int32) (*TL, error) {
	return mconn.InvokeBlocked(TL_messages_getHistory{
		Peer:        peer,
		Offset_id:   offsetId,
		Offset_date: offsetDate,
		Add_offset:  addOffset,
		Limit:       limit,
		Max_id:      maxId,
		Min_id:      minId,
	})
}

func (mconn *MConn) MessagesGetDialogs(excludePinned bool, offsetDate, offsetId int32, offsetPeer TL, limit int32) (*TL, error) {
	return mconn.InvokeBlocked(TL_messages_getDialogs{
		//Exclude_pinned: excludePinned,
		Offset_date: offsetDate,
		Offset_id:   offsetId,
		Offset_peer: offsetPeer,
		Limit:       limit,
	})
}

func (mconn *MConn) MessagesSendMessage( /*no_webpage, silent, background, clear_draft bool, */ peer TL, message string) (*TL, error) {
	return mconn.InvokeBlocked(TL_messages_sendMessage{
		//Flags: 0x00000080,
		//Flags: 0x00000002,
		//No_webpage:      no_webpage,
		//Silent:          silent,
		//Background:      background,
		//Clear_draft:     clear_draft,
		Peer: peer,
		//Reply_to_msg_id: reply_to_msg_id,
		//Reply_to_msg_id: 0,
		Message: message,
		//Random_id:       random_id,
		Random_id: rand.Int63(),
		//Reply_markup:    reply_markup,
		//Entities:        entities,
		//Reply_markup:    TL_null{},
		Reply_markup: nil,
		Entities:     nil,
	})
}
