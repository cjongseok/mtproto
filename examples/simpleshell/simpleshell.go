package main

import (
	"log"
	"os"
	"fmt"
	"regexp"

	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/slog"
	"time"
	"strings"
	"strconv"
)

const (
	appVersion = "0.0.1"
	deviceModel = ""
	systemVersion = ""
	language = ""
	sessionFileHome = ""
	telegramAddress = "149.154.167.50:443"
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
	//case mtproto.TL_updateShort:
	//	u := u.(mtproto.TL_updateShort)
	case mtproto.TL_updateShortMessage:
		u := u.(mtproto.TL_updateShortMessage)
		fmt.Printf("msg[%d] %d: %s\n", u.Id, u.User_id, u.Message)
	//case mtproto.TL_updateShortChatMessage:
	//	fmt.Printf("Shell: %T: %v\n", u, u)
	//case mtproto.TL_updateShortSentMessage:
	//	fmt.Printf("Shell: %T: %v\n", u, u)
	case mtproto.US_updates_difference:
		u := u.(mtproto.US_updates_difference)
		for _, m := range u.New_messages {
			fmt.Printf("update-diff msg[%d] %d: %s\n", m.Id, m.From_id, m.Message)
		}
		for _, user := range u.Users {
			fmt.Printf("update-diff user[%d] %s\n", user.Id, user.Username)
		}
		for _, channel := range u.Channels {
			fmt.Printf("update-diff chat[%d] %s\n", channel.Id, channel.Title)
		}
		fmt.Printf("full-update-diff(%s):\n%s\n", time.Now(), slog.StringifyIndent(u, "  "))
	default:
		fmt.Printf("update(%s):\n%s\n", time.Now(), slog.StringifyIndent(u, "  "))
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Println("USAGE: ./simpleshell <APIID> <APIHASH> <PHONE> [ADDR]")
	fmt.Println("")
	fmt.Println("APIID 	means Telegram API id. If you do not have it yet, go https://my.telegram.org/apps")
	fmt.Println("APIHASH 	means hashcode of <APIID>. It is published together with API id.")
	fmt.Println("PHONE 	means phone number in international format w/o hyphen. e.g., +15417543010")
	fmt.Println("ADDR		means preffered Telegram server address in the form of <IP>:<PORT>.")
	fmt.Println("			e.g., 149.154.167.50:443")
}

func help() {
	//fmt.Println("$ <CMD> [ARGS]")
	//fmt.Println("")
	fmt.Println("COMMANDS:")
	fmt.Println("  dialogs\n  contacts\n  toppeers\n  user <ID> <HASH>\n  channel <ID> <HASH>")
	fmt.Println("  allchats\n  help\n  exit")
}

func main() {
	slog.EnableLogging()
	logf, err := os.OpenFile("ss.log", os.O_RDWR | os.O_CREATE | os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logf.Close()
	slog.SetLogOutput(logf)

	if len(os.Args) != 5 {
		usage()
		return
	}
	apiId64, err := strconv.ParseInt(os.Args[1], 10, 32)
	apiId := int32(apiId64)
	apiHash := os.Args[2]
	phoneNumber := os.Args[3]
	if err != nil {
		usage()
		return
	}

	var preferredAddr string
	switch len(os.Args) {
	case 4:
		preferredAddr = ""
	case 5:
		// validate the addr
		matched, err := regexp.MatchString("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{1,5}$", os.Args[4])
		if err != nil || !matched {
			fmt.Println("Invalid ADDR")
			usage()
			return
		}
		preferredAddr = os.Args[1]
	default:
		usage()
		return
	}

	configuration, err := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, sessionFileHome)
	handleError(err)
	manager, err := mtproto.NewManager(configuration)
	handleError(err)


	// Connect by phonenumber
	var mconn *mtproto.MConn
	if manager.IsAuthenticated(phoneNumber) {
		log.Println("MAIN: load authentication")
		mconn, err = manager.LoadAuthentication(phoneNumber, preferredAddr)
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
		args := strings.Split(cmd, " ")

		switch args[0] {

		case "dialogs":		// $ dialogs
			if len(args) != 1 {
				help()
				continue
			}
			resp, err := mconn.MessagesGetDialogs(false, 0, 0, mtproto.TL_inputPeerEmpty{}, 1)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_dialogsSlice).Unstrip(), "  "))

		case "contacts":	// $ contacts
			if len(args) != 1 {
				help()
				continue
			}
			resp, err := mconn.ContactsGetContacts(0)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_contacts_contacts).Unstrip(), "  "))
			handleError(err)

		case "toppeers":	// $ toppeers
			if len(args) != 1 {
				help()
				continue
			}
			resp, err := mconn.ContactsGetTopPeers(true, false, false, true, true, 0, 0, 0)
			log.Println(*resp)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_contacts_topPeers).Unstrip(), "  "))

		case "user":	// $ user <ID> <HASH>
			if len(args) != 3 {
				help()
				continue
			}
			id64, err := strconv.ParseInt(args[1], 10, 0)
			if err != nil {
				fmt.Printf("Invalid ID, %s: ID should be integer\n", args[1])
				continue
			}
			hash, err := strconv.ParseInt(args[2], 10, 0)
			if err != nil {
				fmt.Printf("Invalid HASH, %s: HASH should be integer\n", args[1])
				continue
			}
			user, err := mconn.UsersGetFullUsers(mtproto.TL_inputUser{int32(id64), hash})
			handleError(err)
			fmt.Println(slog.StringifyIndent(user, "  "))


		case "channel":	// $ channel <ID> <HASH>
			if len(args) != 3 {
				help()
				continue
			}
			id64, err := strconv.ParseInt(args[1], 10, 0)
			if err != nil {
				fmt.Printf("Invalid ID, %s: ID should be integer\n", args[1])
				continue
			}
			hash64, err := strconv.ParseInt(args[2], 10, 0)
			if err != nil {
				fmt.Printf("Invalid HASH, %s: HASH should be integer\n", args[1])
				continue
			}
			ids := []mtproto.TL{
				mtproto.TL_inputChannel{int32(id64), hash64},
			}
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getChannels{ids})
			log.Println(*resp)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))

		//case "chats":		// $ chats <ID>...
		//	if len(args) < 2 {
		//		help()
		//		continue
		//	}
		//	//TODO: parse IDs
		//	resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getChats{make([]int32, 0)})
		//	log.Println(*resp)
		//	handleError(err)
		//	fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))

		case "allchats": 	// $ allchats
			if len(args) != 1 {
				help()
				continue
			}
			resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getAllChats{make([]int32, 0)})
			log.Println(*resp)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))

		//case "search":
		//	resp, err := mconn.InvokeBlocked(mtproto.TL_contacts_search{args[0], 0})

		case "help":
			if len(args) != 1 {
				help()
				continue
			}
			help()

		case "exit":
			if len(args) != 1 {
				help()
				continue
			}
			return

		default:
			fmt.Println("Wrong command")
			help()
		}
	}
}
