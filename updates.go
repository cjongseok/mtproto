package mtproto

func (mconn *MConn) UpdatesGetState() (*TL, error) {
	return mconn.InvokeBlocked(TL_updates_getState{})
}

func (mconn *MConn) UpdatesGetDifference(pts, ptsTotalLimit, date, qts int32) (*TL, error) {
	return mconn.InvokeBlocked(TL_updates_getDifference{
		Pts:             pts,
		Pts_total_limit: ptsTotalLimit,
		Date:            date,
		Qts:             qts,
	})
}

func (mconn *MConn) UpdatesGetChannelDifference(force bool, channel, filter TL, pts, limit int32) (*TL, error) {
	return mconn.InvokeBlocked(TL_updates_getChannelDifference{
		Force:   force,
		Channel: channel,
		Filter:  filter,
		Pts:     pts,
		Limit:   limit,
	})
}
