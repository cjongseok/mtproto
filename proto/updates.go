package mtp

//func (mconn *MConn) UpdatesGetState() (TL, error) {
//	return mconn.InvokeBlocked(&ReqUpdatesGetState{})
//}
//
//func (mconn *MConn) UpdatesGetDifference(pts, ptsTotalLimit, date, qts int32) (TL, error) {
//	return mconn.InvokeBlocked(&ReqUpdatesGetDifference{
//		Pts:           pts,
//		PtsTotalLimit: ptsTotalLimit,
//		Date:          date,
//		Qts:           qts,
//	})
//}
//
//func (mconn *MConn) UpdatesGetChannelDifference(force bool, channel *TypeInputChannel, filter *TypeChannelMessagesFilter, pts, limit int32) (TL, error) {
//	return mconn.InvokeBlocked(&ReqUpdatesGetChannelDifference{
//		//Force:   force,
//		Channel: channel,
//		Filter:  filter,
//		Pts:     pts,
//		Limit:   limit,
//	})
//}
