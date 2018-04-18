package main

import (
	"fmt"
	"github.com/cjongseok/mtproto"
	"os"
	"strconv"
	"github.com/cjongseok/slog"
)

func usage() {
	fmt.Println(`Usage: keygen <APIID> <APIHASH> <PHONE> [KEY_NAME]

Params:
  APIID    means Telegram API id. If you do not have it yet, go "https://my.telegram.org/apps".
  APIHASH  means hashcode of <APIID>. It is published together with API id.
  PHONE    means phone number in international format w/o hyphen. e.g., "+15417543010".

Options:
  KEY_NAME means key file name to be generated. Default is "key.mtproto".
`)
}

const (
	defaultNewKeyFile = "key.mtproto"
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
	telegramAddress = "149.154.167.50:443"
)

func main() {
	slog.DisableLogging()

	// parse arguments
	apiId, apiHash, phoneNumber, key, err := parseArgs()
	if err != nil {
		usage()
		handleError(err)
	}
	if key == "" {
		key = defaultNewKeyFile
	}
	config, err := mtproto.NewConfiguration(apiId, apiHash, appVersion, deviceModel, systemVersion, language, 0, 0, key)
	handleError(err)

	// check if the file exists
	_, err = os.Stat(key)
	sessionExists := !os.IsNotExist(err)
	if sessionExists {
		handleError(fmt.Errorf("\"%s\" already exists\n", key))
	}

	// request to send authentication code to the phone
	var manager *mtproto.Manager
	var mconn *mtproto.Conn
	var sentCode *mtproto.TypeAuthSentCode
	manager, err = mtproto.NewManager(config)
	handleError(err)
	mconn, sentCode, err = manager.NewAuthentication(phoneNumber, telegramAddress, false)
	handleError(err)

	// sign-in with the code from the user input
	var code string
	fmt.Printf("Enter Code: ")
	fmt.Scanf("%s", &code)
	auth, err := mconn.SignIn(phoneNumber, code, sentCode.GetValue().PhoneCodeHash)
	handleError(err)
	if auth.Value.GetUser().GetUser() != nil {
		user := auth.Value.GetUser().GetUser()
		fmt.Printf("Signed in as \"%s\"\n", user.FirstName)
	} else if auth.Value.GetUser().GetUserEmpty() != nil {
		fmt.Println("Signed in with empty user")
	} else {
		fmt.Println("Signed in without user response: neither user nor user empty")
	}
	fmt.Println("Key is generated.")
}

func handleError(err error) {
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
}

func parseArgs() (apiId int32, apiHash, phoneNumber, key string, err error){
	switch len(os.Args) {
	case 4:
	case 5:
		key = os.Args[4]
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
	return
}
