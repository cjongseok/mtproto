package main

import (
	"log"
	"os"

	"github.com/cjongseok/mtproto"
	"fmt"
	"reflect"
	"encoding/json"
)

const (
	apiId = 12345	// YOUR API_ID
	apiHash = "abcdefghijk1234567890" // YOUR API_HASH
	appVersion = "0.0.1"
	deviceModel = ""
	systemVersion = ""
	language = ""
	sessionFileHome = ""

	telegramAddress = "149.154.167.50:443"
	phoneNumber = "+0177778888"	// YOUR phone number
)

type subscriber struct{
	mconn *mtproto.MConn
}

func newSubscriber(mconn *mtproto.MConn) *subscriber {
	s := new(subscriber)
	s.mconn = mconn
	return s
}

func (s *subscriber) OnUpdate(u mtproto.MUpdate) {
	switch u.(type) {
	case mtproto.TL_updateShort:
		u := u.(mtproto.TL_updateShort)
		switch u.Update.(type) {
		case mtproto.TL_updateUserStatus:
			status := u.Update.(mtproto.TL_updateUserStatus)
			fmt.Printf("StatusChange: %v ", reflect.TypeOf(status.Status))
			switch status.Status.(type) {
			case mtproto.TL_userStatusOnline:
				online := status.Status.(mtproto.TL_userStatusOnline)
				fmt.Printf("Expires: %d\n", online.Expires)
			case mtproto.TL_userStatusOffline:
				offline := status.Status.(mtproto.TL_userStatusOffline)
				fmt.Printf("WasOnline: %d\n", offline.Was_online)
			}
		}
	case mtproto.TL_updateShortMessage:
		u := u.(mtproto.TL_updateShortMessage)
		fmt.Printf("msg[%d] %d: %s\n", u.Id, u.User_id, u.Message)
	case mtproto.TL_updateShortChatMessage:
		fmt.Printf("Shell: %T: %v\n", u, u)
	case mtproto.TL_updateShortSentMessage:
		fmt.Printf("Shell: %T: %v\n", u, u)
	case mtproto.US_updates_difference:
		u := u.(mtproto.US_updates_difference)
		for _, m := range u.New_messages {
			fmt.Printf("update-diff msg[%d] %d: %s\n", m.Id, m.From_id, m.Message)
		}
		for _, user := range u.Users {
			fmt.Printf("update-diff user[%d] %s\n", user.Id, user.Username)
		}
		for _, chat := range u.Chats {
			fmt.Printf("update-diff chat[%d] %s\n", chat.Id, chat.Title)
		}
		fmt.Printf("update-diff emsgs{%v} state{%v} others{%v}\n", u.New_encrypted_messages, u.State, u.Other_updates)
	default:
		marshaled, err := json.Marshal(u)
		if err == nil {
			fmt.Printf("uncaught update: %T%s\n", u, marshaled)
		} else {
			fmt.Printf("uncaught update: %T{%v}\n", u, u)
		}
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("simpleshell <CMD> [ARGS]")
	fmt.Println("")
	fmt.Println("COMMANDS:")
	fmt.Println("  dialogs, contacts, toppeers, channels, chats, allchats, help, exit")
}

func main() {
	logf, err := os.OpenFile("ss.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logf.Close()

	//log.SetOutput(os.Stdout)
	log.SetOutput(logf)
	configuration, err := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, sessionFileHome)
	handleError(err)
	manager, err := mtproto.NewManager(configuration)
	handleError(err)


	// Connect by phonenumber
	var mconn *mtproto.MConn
	if manager.IsAuthenticated(phoneNumber) {
		log.Println("MAIN: load authentication")
		mconn, err = manager.LoadAuthentication(phoneNumber)
		handleError(err)

	} else {
		log.Println("MAIN: new authentication")
		var sentCode *mtproto.TL_auth_sentCode
		mconn, sentCode, err = manager.NewAuthentication(phoneNumber, telegramAddress, false)
		handleError(err)
		log.Println("MAIN: sent code, ", sentCode)

		log.Println("MAIN: sign in")
		var code string
		fmt.Printf("Enter Code: ")
		fmt.Scanf("%s", &code)
		log.Println("entered code = ", code)
		_, err = mconn.AuthSignIn(phoneNumber, code, sentCode.Phone_code_hash)
		handleError(err)
	}

	mconn.AddUpdateCallback(newSubscriber(mconn))

	for {
		var cmd string
		fmt.Printf("$ ")
		fmt.Scanf("%s", &cmd)
		//cmd := "dialogs"
		switch cmd {

		case "dialogs":
			resp, err := mconn.MessagesGetDialogs(false, 0, 0, mtproto.TL_inputPeerEmpty{}, 1)
			handleError(err)
			fmt.Println((*resp).(mtproto.TL_messages_dialogsSlice).Unstrip())

		case "contacts":
			resp, err := mconn.ContactsGetContacts("")
			fmt.Println((*resp).(mtproto.TL_contacts_contacts).Unstrip())
			handleError(err)

		case "toppeers":
			resp, err := mconn.ContactsGetTopPeers(true, false, false, true, true, 0, 0, 0)
			log.Println(*resp)
			handleError(err)
			fmt.Println((*resp).(mtproto.TL_contacts_topPeers).Unstrip())

		case "channels":
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getChannels{make([]mtproto.TL, 0)})
			log.Println(*resp)
			handleError(err)
			fmt.Println((*resp).(mtproto.TL_messages_chats).Unstrip())

		case "chats":
			resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getChats{make([]int32, 0)})
			log.Println(*resp)
			handleError(err)
			fmt.Println((*resp).(mtproto.TL_messages_chats).Unstrip())

		case "allchats":
			resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getAllChats{make([]int32, 0)})
			log.Println(*resp)
			handleError(err)
			fmt.Println((*resp).(mtproto.TL_messages_chats).Unstrip())

		//case "search":
		//	resp, err := mconn.InvokeBlocked(mtproto.TL_contacts_search{args[0], 0})

		case "help":
			usage()

		case "exit":
			return

		default:
			fmt.Println("Wrong command")
			usage()
		}
	}
}
