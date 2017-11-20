package main

import (
	"log"
	"os"

	"github.com/cjongseok/mtproto"
	"fmt"
	"strings"
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

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.SetOutput(os.Stdout)
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
		mconn, sentCode, err := manager.NewAuthentication(phoneNumber, telegramAddress, false)
		handleError(err)
		log.Println("MAIN: sent code, ", sentCode)

		log.Println("MAIN: sign in")
		var code string
		fmt.Printf("Enter Code: ")
		fmt.Scanf("%s", &code)
		_, err = mconn.AuthSignIn(phoneNumber, code, sentCode.Phone_code_hash)
		handleError(err)
	}

	for {
		var cmd string
		var args []string
		fmt.Printf("$ ")
		fmt.Scanf("%s", &cmd)

		splits := strings.Split(cmd, " ")
		if len(splits) < 1 {
			continue
		} else {
			cmd = splits[0]
			if len(splits) > 1 {
				args = splits[1:]
			}
		}

		switch cmd {
		case "dialogs":
			resp, err := mconn.MessagesGetDialogs(false, 0, 0, mtproto.TL_inputPeerEmpty{}, 1)
			if err != nil {
				fmt.Println(err)
			}
			if resp != nil {
				fmt.Println(resp)
			}

		case "exit":
			return

		default:
			fmt.Printf("Wrong command (cmd: %s, args: %v)", cmd, args)
		}
	}
}
