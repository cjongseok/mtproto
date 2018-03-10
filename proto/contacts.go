package mtp

import "errors"

//func (mconn *MConn) ContactsGetContacts(hash int32) (TL, error) {
//	return mconn.InvokeBlocked(&ReqContactsGetContacts{
//		Hash: hash,
//	})
//}

func (mconn *MConn) ContactsGetTopPeers(correspondents, botsPM, botsInline, groups, channels bool, offset, limit, hash int32) (*TypeContactsTopPeers, error) {
	x := <-mconn.InvokeNonBlocked(&ReqContactsGetTopPeers{
		//Correspondents: correspondents,
		//Bots_pm:        botsPM,
		//Bots_inline:    botsInline,
		//Groups:         groups,
		//Channels:       channels,
		Offset: offset,
		Limit:  limit,
		Hash:   hash,
	})

	if x.err != nil {
		return nil, x.err
	}

	//	*TypeContactsTopPeers_ContactsTopPeersNotModified
	//	*TypeContactsTopPeers_ContactsTopPeers
	switch peers := x.data.(type) {
	case *PredContactsTopPeersNotModified:
		return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeersNotModified{peers}}, nil
	case *PredContactsTopPeers:
		return &TypeContactsTopPeers{&TypeContactsTopPeers_ContactsTopPeers{peers}}, nil
		//case *TypeContactsTopPeers:
		//	return x.data.(*TypeContactsTopPeers), nil
	}
	return nil, errors.New("MTProto::ContactsGetTopPeers error: Unknown type")
}
