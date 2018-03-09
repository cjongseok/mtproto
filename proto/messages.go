package mtp

import (
	"math/rand"
)

func (mconn *MConn) MessagesGetHistory(peer *TypeInputPeer, offsetId, offsetDate, addOffset, limit, maxId, minId int32) (TL, error) {
	return mconn.InvokeBlocked(&ReqMessagesGetHistory{
		Peer:       peer,
		OffsetId:   offsetId,
		OffsetDate: offsetDate,
		AddOffset:  addOffset,
		Limit:      limit,
		MaxId:      maxId,
		MinId:      minId,
	})
}

func (mconn *MConn) MessagesGetDialogs(excludePinned bool, offsetDate, offsetId int32, offsetPeer *TypeInputPeer, limit int32) (TL, error) {
	return mconn.InvokeBlocked(&ReqMessagesGetDialogs{
		//Exclude_pinned: excludePinned,
		OffsetDate: offsetDate,
		OffsetId:   offsetId,
		OffsetPeer: offsetPeer,
		Limit:      limit,
	})
}

func (mconn *MConn) MessagesSendMessage( /*no_webpage, silent, background, clear_draft bool, */ peer *TypeInputPeer, message string) (TL, error) {
	return mconn.InvokeBlocked(&ReqMessagesSendMessage{
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
		RandomId: rand.Int63(),
		//Reply_markup:    reply_markup,
		//Entities:        entities,
		//Reply_markup:    TL_null{},
		ReplyMarkup: nil,
		Entities:    nil,
	})
}
