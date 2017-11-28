package mtproto

type MUpdate interface {
	UpdateDate() int32
}

type MUpdateCallback interface {
	OnUpdate(update MUpdate)
}

func (u TL_updateShort) UpdateDate() int32 {return u.Date}
func (u TL_updateShortMessage) UpdateDate() int32 {return u.Date}
func (u TL_updateShortChatMessage) UpdateDate() int32 {return u.Date}
func (u TL_updateShortSentMessage) UpdateDate() int32 {return u.Date}
func (u US_updates_difference) UpdateDate() int32 {return 0}
func (u TL_updates) UpdateDate() int32 {return u.Date}
func (u TL_updateNewMessage) UpdateDate() int32 {return 0}
func (u TL_updateReadMessagesContents) UpdateDate() int32 {return 0}
func (u TL_updateDeleteMessages) UpdateDate() int32 {return 0}
func (u TL_updateNewEncryptedMessage) UpdateDate() int32 {return 0}
