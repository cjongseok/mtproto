package mtproto

import "errors"

func (mconn *MConn) ContactsGetContacts(hash string) (*TL, error) {
	return mconn.InvokeBlocked(TL_contacts_getContacts{
		Hash: hash,
	})
}

func (mconn *MConn) ContactsGetTopPeers(correspondents, botsPM, botsInline, groups, channels bool, offset, limit, hash int32) (*TL, error) {
	x := <-mconn.InvokeNonBlocked(TL_contacts_getTopPeers{
		Correspondents: correspondents,
		Bots_pm:        botsPM,
		Bots_inline:    botsInline,
		Groups:         groups,
		Channels:       channels,
		Offset:         offset,
		Limit:          limit,
		Hash:           hash,
	})

	if x.err != nil {
		return nil, x.err
	}

	switch x.data.(type) {
	case TL_contacts_topPeersNotModified:
	case TL_contacts_topPeers:
	default:
		return nil, errors.New("MTProto::ContactsGetTopPeers error: Unknown type")
	}

	return &x.data, nil
}
