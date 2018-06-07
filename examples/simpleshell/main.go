package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"bufio"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/slog"
	"strconv"
	"strings"
	"time"
	"math/rand"
	"golang.org/x/net/context"
	"net"
)

const (
	defaultNewKeyFile = "key.mtproto"

	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
	telegramAddress = "149.154.167.50:443"
)

func usage() {
	fmt.Println(`USAGE: ./simpleshell <APIID> <APIHASH> <PHONE> <ADDR|KEY>
Params:
  APIID     means Telegram API id. If you do not have it yet, go https://my.telegram.org/apps
  APIHASH   means hashcode of <APIID>. It is published together with API id.
  PHONE     means phone number in international format w/o hyphen. e.g., +15417543010
  ADDR      means preffered Telegram server address in the form of <IP>:<PORT>.
            You can find a vaild address in your https://my.telegram.org/apps page. Sign-in
            without key would require the user to input a verification code sent to the user's
            Telegram account, and generate the key file, 'key.mtproto'.
  KEY       means MTProto key file. It sign-in to the stored Telegram server endpoint with the
            phone number using its credentials.
`)
}

type subscriber struct {
	mconn *mtproto.Conn
}

func newSubscriber(mconn *mtproto.Conn) *subscriber {
	s := new(subscriber)
	s.mconn = mconn
	return s
}

func isServerEndpoint(addr string) (err error) {
	var tcpAddr *net.TCPAddr
	tcpAddr, err = net.ResolveTCPAddr("tcp", addr)
	if err == nil && tcpAddr.IP.To4() == nil && tcpAddr.IP.To16() == nil {
		err = fmt.Errorf("invalid ip address")
	}
	return
}

func parseArgs() (apiId int32, apiHash, phoneNumber, preferredAddr, key string, err error){
	switch len(os.Args) {
	case 5:
	default:
		err = fmt.Errorf("invalid number of arguments")
		return
	}

	var apiId64 int64
	apiId64, err = strconv.ParseInt(os.Args[1], 10, 32)
	if err != nil {
		return
	}
	apiId = int32(apiId64)
	apiHash = os.Args[2]
	phoneNumber = os.Args[3]

	var matched bool
	matched, err = regexp.MatchString("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}:\\d{1,5}$", os.Args[4])
	if matched && err == nil {
		preferredAddr = os.Args[4]
		return
	}
	if _, err = os.Stat(os.Args[4]); err == nil {
		key = os.Args[4]
		return
	}
	err = fmt.Errorf("invalid address or non-existent file")
	return
}

func (s *subscriber) OnUpdate(u mtproto.Update) {
	fmt.Printf("update(%s):\n%s\n", time.Now(), slog.StringifyIndent(u, "  "))
}

func handleError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func help() {
	fmt.Println("Commands:")
	fmt.Printf("\tdialogs\n\tsend2c <CHAN_ID> <CHAN_HASH> <MSG>\n\thelp\n\texit\n")
}

func main() {
	// set up logging
	logf, err := os.OpenFile("ss.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("error opening file: %v", err)
	}
	defer logf.Close()
	slog.SetLogOutput(logf)

	// parse args
	apiId, apiHash, phoneNumber, addr, key, err := parseArgs()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		usage()
		handleError(err)
	}

	config, err := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, key)
	handleError(err)

	// Connect by phone number
	var manager *mtproto.Manager
	var mconn *mtproto.Conn
	if config.KeyPath == "" {
		config.KeyPath = defaultNewKeyFile
		log.Println("MAIN: new authentication")

		// request to send authentication code to the phone
		var sentCode *mtproto.TypeAuthSentCode
		manager, err = mtproto.NewManager(config)
		handleError(err)
		mconn, sentCode, err = manager.NewAuthentication(phoneNumber, addr, false)
		handleError(err)

		// sign-in with the code from the user input
		var code string
		fmt.Printf("Enter Code: ")
		fmt.Scanf("%s", &code)
		log.Println("entered code = ", code)
		_, err = mconn.SignIn(phoneNumber, code, sentCode.GetValue().PhoneCodeHash)
		handleError(err)
	} else {
		log.Println("MAIN: load authentication")
		manager, err = mtproto.NewManager(config)
		handleError(err)
		mconn, err = manager.LoadAuthentication(phoneNumber)
		handleError(err)
	}

	mconn.AddUpdateCallback(newSubscriber(mconn))

	caller := mtproto.RPCaller{mconn}
	for {
		// ready for the user command
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

		// parse & execute the command
		switch args[0] {
		case "dialogs": // $ dialogs
			if len(args) != 1 {
				help()
				break
			}
			emptyPeer := &mtproto.TypeInputPeer{&mtproto.TypeInputPeer_InputPeerEmpty{&mtproto.PredInputPeerEmpty{}}}
			resp, err := caller.MessagesGetDialogs(context.Background(), &mtproto.ReqMessagesGetDialogs{
				OffsetDate: 0, OffsetId: 0, OffsetPeer: emptyPeer, Limit: 1,
			})
			handleError(err)
			switch dialogs := resp.GetValue().(type) {
			case *mtproto.TypeMessagesDialogs_MessagesDialogs:
				fmt.Println(slog.StringifyIndent(dialogs.MessagesDialogs, "  "))
			case *mtproto.TypeMessagesDialogs_MessagesDialogsSlice:
				fmt.Println(slog.StringifyIndent(dialogs.MessagesDialogsSlice, "  "))
			}
		case "send2c": // send2c <CHAN_ID> <CHAN_HASH> <MSG>
			if len(args) != 4 {
				help()
				return
			}
			chanId, err := strconv.ParseInt(args[1], 0, 32)
			handleError(err)
			chanHash, err := strconv.ParseInt(args[2], 0, 64)
			handleError(err)
			peer := &mtproto.TypeInputPeer{&mtproto.TypeInputPeer_InputPeerChannel{
				&mtproto.PredInputPeerChannel{
					int32(chanId), int64(chanHash),
				}}}
			resp, err := caller.MessagesSendMessage(context.Background(), &mtproto.ReqMessagesSendMessage{
				Peer:      peer,
				Message:   args[3],
				RandomId: rand.Int63(),
			})
			handleError(err)
			fmt.Println("send response:", slog.StringifyIndent(resp, "  "))
		case "contacts": // $ contacts
			if len(args) != 1 {
				help()
				return
			}
			resp, err := caller.ContactsGetContacts(context.Background(), &mtproto.ReqContactsGetContacts{})
			handleError(err)
			fmt.Println("contacts response:", slog.StringifyIndent(resp, "  " ))
		case "help":
			help()
		case "exit":
			if len(args) != 1 {
				help()
				break
			}
			return
		case "": // new line
		default:
			fmt.Println("Wrong command")
			help()
		}
	}
}
