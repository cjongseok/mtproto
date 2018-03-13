package core

type Update interface {
	Predicate
	UpdateDate() int32
}

type UpdateCallback interface {
	OnUpdate(update Update)
}

//PredUpdatesState Value = 1;

//PredUpdatesTooLong UpdatesTooLong = 1;
//PredUpdateShortMessage UpdateShortMessage = 2;
//PredUpdateShortChatMessage UpdateShortChatMessage = 3;
//PredUpdateShort UpdateShort = 4;
//PredUpdatesCombined UpdatesCombined = 5;
//PredUpdates Updates = 6;
//PredUpdateShortSentMessage UpdateShortSentMessage = 7;

func (u *PredUpdatesState) UpdateDate() int32           { return u.Date }

func (u *PredUpdateShortMessage) UpdateDate() int32     { return u.Date }
func (u *PredUpdateShortChatMessage) UpdateDate() int32 { return u.Date }
func (u *PredUpdateShort) UpdateDate() int32            { return u.Date }
func (u *PredUpdates) UpdateDate() int32                { return u.Date }
func (u *PredUpdateShortSentMessage) UpdateDate() int32 { return u.Date }

func (u *PredUpdatesDifference) UpdateDate() int32      { return 0 }
func (u *PredUpdatesDifferenceSlice) UpdateDate() int32 { return 0 }

//func (u US_updates_difference) UpdateDate() int32         { return 0 }
func (u *PredUpdateNewMessage) UpdateDate() int32           { return 0 }
func (u *PredUpdateReadMessagesContents) UpdateDate() int32 { return 0 }
func (u *PredUpdateDeleteMessages) UpdateDate() int32       { return 0 }
func (u *PredUpdateNewEncryptedMessage) UpdateDate() int32  { return 0 }

func (u *PredUpdateChannel) UpdateDate() int32             { return 0 }
func (u *PredUpdateChannelMessageViews) UpdateDate() int32 { return 0 }
func (u *PredUpdateChannelTooLong) UpdateDate() int32      { return 0 }
func (u *PredUpdateReadChannelInbox) UpdateDate() int32    { return 0 }
func (u *PredUpdateReadChannelOutbox) UpdateDate() int32   { return 0 }
func (u *PredUpdateNewChannelMessage) UpdateDate() int32   { return 0 }
