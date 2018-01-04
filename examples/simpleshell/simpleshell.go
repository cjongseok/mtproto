package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	"bufio"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/slog"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const (
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
	sessionFileHome = ""
	telegramAddress = "149.154.167.50:443"
)

type subscriber struct {
	mconn *mtproto.MConn
}

func newSubscriber(mconn *mtproto.MConn) *subscriber {
	s := new(subscriber)
	s.mconn = mconn
	return s
}

func (s *subscriber) OnUpdate(u mtproto.MUpdate) {
	switch u.(type) {
	case mtproto.TL_updates:

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
		for _, chat := range u.Chats {
			fmt.Printf("update-diff chat[%d] %s\n", chat.Id, chat.Title)
		}
		for _, channel := range u.Channels {
			fmt.Printf("update-diff channel[%d] %s\n", channel.Id, channel.Title)
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
	fmt.Println("USAGE: ./simpleshell <APIID> <APIHASH> <PHONE> <ADDR>")
	fmt.Println("")
	fmt.Println("APIID 		means Telegram API id. If you do not have it yet, go https://my.telegram.org/apps")
	fmt.Println("APIHASH 	means hashcode of <APIID>. It is published together with API id.")
	fmt.Println("PHONE 		means phone number in international format w/o hyphen. e.g., +15417543010")
	fmt.Println("ADDR		means preffered Telegram server address in the form of <IP>:<PORT>. ")
	fmt.Println(" 			You can find a vaild address in your https://my.telegram.org/apps page.")
}

func help() {
	//fmt.Println("$ <CMD> [ARGS]")
	//fmt.Println("")
	fmt.Println("COMMANDS:")
	fmt.Println("  dialogs\n  contacts\n  toppeers\n  user <ID> <HASH>\n  channel <ID> <HASH>")
	fmt.Println("  allchats\n  help\n  exit")
}

func main() {
	//slog.EnableLogging()
	logf, err := os.OpenFile("ss.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
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

	configuration, err := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, sessionFileHome, 0, 0)
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
		reader := bufio.NewReader(os.Stdin)
		cmd, err := reader.ReadString('\n')
		if err != nil {
			help()
			continue
		}
		cmd = strings.Trim(cmd, "\n")
		args := strings.Split(cmd, " ")

		switch args[0] {

		case "dialogs": // $ dialogs
			if len(args) != 1 {
				help()
				break
			}
			resp, err := mconn.MessagesGetDialogs(false, 0, 0, mtproto.TL_inputPeerEmpty{}, 1)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_dialogsSlice).Unstrip(), "  "))

		case "contacts": // $ contacts
			if len(args) != 1 {
				help()
				break
			}
			resp, err := mconn.ContactsGetContacts(0)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_contacts_contacts).Unstrip(), "  "))
			handleError(err)

		case "toppeers": // $ toppeers
			if len(args) != 1 {
				help()
				break
			}
			resp, err := mconn.ContactsGetTopPeers(true, false, false, true, true, 0, 0, 0)
			log.Println(*resp)
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_contacts_topPeers).Unstrip(), "  "))

		case "user": // $ user <ID> <HASH>
			if len(args) != 3 {
				help()
				break
			}
			id64, err := strconv.ParseInt(args[1], 10, 0)
			if err != nil {
				fmt.Printf("Invalid ID, %s: ID should be integer\n", args[1])
				break
			}
			hash, err := strconv.ParseInt(args[2], 10, 0)
			if err != nil {
				fmt.Printf("Invalid HASH, %s: HASH should be integer\n", args[1])
				break
			}
			user, err := mconn.UsersGetFullUsers(mtproto.TL_inputUser{int32(id64), hash})
			handleError(err)
			fmt.Println(slog.StringifyIndent(user, "  "))

		case "channel": // $ channel <ID> <HASH>
			if len(args) != 3 {
				help()
				break
			}
			inputchannel, err := parseInputChannel(args[1], args[2])
			ids := []mtproto.TL{inputchannel}
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getChannels{ids})
			handleError(err)
			log.Println(*resp)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))

		case "fullchannel": // $ fullchan <ID> <HASH>
			if len(args) != 3 {
				help()
				break
			}
			inputchannel, err := parseInputChannel(args[1], args[2])
			handleError(err)
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getFullChannel{inputchannel})
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chatFull).Unstrip(), "  "))

		case "chanmsg": // $ chanmsg <CHAN_ID> <CHAN_HASH>
			if len(args) != 3 {
				help()
				break
			}
			inputchannel, err := parseInputChannel(args[1], args[2])
			handleError(err)
			resp, err := mconn.InvokeBlocked(mtproto.TL_channels_getMessages{inputchannel, []int32{}})
			handleError(err)
			fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_messages).Unstrip(), "  "))

		case "msg": // $ msg <MSG_ID>
			if len(args) != 2 {
				help()
				break
			}
			id, err := strconv.Atoi(args[1])
			handleError(err)
			resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getMessages{[]int32{int32(id)}})
			handleError(err)
			fmt.Println(slog.StringifyIndent(resp, "  "))

		case "chanhistory": // chanhistory <CHAN_ID> <CHAN_HASH> <LIMIT>
			if len(args) != 4 {
				help()
				break
			}
			peer, err := parseInputPeerChannel(args[1], args[2])
			handleError(err)
			limit, err := strconv.Atoi(args[3])
			handleError(err)
			resp, err := mconn.MessagesGetHistory(*peer, 0, 0, 0, int32(limit), 0, 0)
			handleError(err)
			fmt.Println(slog.StringifyIndent(resp, "  "))

		case "send2c": // send2c <CHAN_ID> <CHAN_HASH> <MSG>
			if len(args) != 4 {
				help()
				return
			}
			chanId, err := strconv.ParseInt(args[1], 0, 32)
			handleError(err)
			chanHash, err := strconv.ParseInt(args[2], 0, 64)
			handleError(err)
			//resp, err := mconn.MessagesSendMessage(mtproto.TL_peerChannel{int32(chanId)}, args[2])

			resp, err := mconn.InvokeBlocked(mtproto.TL_messages_sendMessage{
				Peer:      mtproto.TL_inputPeerChannel{int32(chanId), int64(chanHash)},
				Message:   args[3],
				Random_id: rand.Int63(),
			})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(slog.StringifyIndent(*resp, " "))
			//fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_dialogsSlice).Unstrip(), "  "))

		case "send2u": // send2u <USER_ID> <MSG>
			if len(args) != 3 {
				help()
				return
			}
			userId, err := strconv.ParseInt(args[1], 0, 32)
			handleError(err)
			resp, err := mconn.MessagesSendMessage(mtproto.TL_peerUser{int32(userId)}, args[2])
			handleError(err)
			fmt.Println(slog.StringifyIndent(*resp, " "))

			//case "rcvdmsg": // $ rcvdmsg
			//	if len(args) != 1 {
			//		help()
			//		break
			//	}
			//	resp, err := mconn.InvokeBlocked(mtproto.TL_messages_receivedMessages{0})
			//	handleError(err)
			//	fmt.Println(slog.StringifyIndent(resp, "  "))

			//case "chats":		// $ chats <ID>...
			//	if len(args) < 2 {
			//		help()
			//		break
			//	}
			//	//TODO: parse IDs
			//	resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getChats{make([]int32, 0)})
			//	log.Println(*resp)
			//	handleError(err)
			//	fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))

		case "allchats": // $ allchats [-f <filename>]
			switch len(args) {
			case 1:
				resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getAllChats{make([]int32, 0)})
				handleError(err)
				log.Println(*resp)
				fmt.Println(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))
			case 3:
				if args[1] == "-f" {
					resp, err := mconn.InvokeBlocked(mtproto.TL_messages_getAllChats{make([]int32, 0)})
					handleError(err)
					log.Println(*resp)
					f, err := os.OpenFile(args[2], os.O_CREATE|os.O_WRONLY, 0666)
					if err != nil {
						fmt.Printf("Cannot create/open file (%s): %v\n", args[2], err)
					} else {
						f.WriteString(slog.StringifyIndent((*resp).(mtproto.TL_messages_chats).Unstrip(), "  "))
					}
				} else {
					help()
				}
			default:
				help()
			}
			//case "search":
			//	resp, err := mconn.InvokeBlocked(mtproto.TL_contacts_search{args[0], 0})

		case "help":
			help()

		case "exit":
			if len(args) != 1 {
				help()
				break
			}
			return

		case "": // new line
			// ignore it

		default:
			fmt.Println("Wrong command")
			help()
		}
	}
}

func parseInputPeerChannel(id string, hash string) (*mtproto.TL_inputPeerChannel, error) {
	id64, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("Invalid ID, %s: ID should be integer\n", id)
	}
	hash64, err := strconv.ParseInt(hash, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("Invalid HASH, %s: HASH should be integer\n", hash)
	}
	peer := new(mtproto.TL_inputPeerChannel)
	peer.Channel_id = int32(id64)
	peer.Access_hash = hash64
	return peer, nil
}

func parseInputChannel(id string, hash string) (*mtproto.TL_inputChannel, error) {
	id64, err := strconv.ParseInt(id, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("Invalid ID, %s: ID should be integer\n", id)
	}
	hash64, err := strconv.ParseInt(hash, 10, 0)
	if err != nil {
		return nil, fmt.Errorf("Invalid HASH, %s: HASH should be integer\n", hash)
	}
	inputchannel := new(mtproto.TL_inputChannel)
	inputchannel.Channel_id = int32(id64)
	inputchannel.Access_hash = hash64
	return inputchannel, nil
}
