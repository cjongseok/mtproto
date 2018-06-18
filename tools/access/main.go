package main

import (
	"flag"
	"github.com/cjongseok/mtproto"
	"os"
	"fmt"
	"context"
	"github.com/cjongseok/slog"
	"github.com/golang/protobuf/proto"
)

const (
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
)

var (
	credentials = flag.String("credentials", "credentials.json", "Credentials file. Default is credentials.json")
	chans       = flag.String("chans", "chans.acc", "Output file. Default is chans.json")
	users       = flag.String("users", "users.acc", "Output file. Default is users.json")
)

const logfile = "access.log"

func main() {
	flag.Parse()
	if *credentials == "" || *chans== "" || *users == "" {
		flag.Usage()
		os.Exit(1)
	}

	slog.SetLogOutputAsFile(logfile)

	chatsFp, err := os.OpenFile(*chans, os.O_WRONLY | os.O_CREATE, 0644)
	handleError(err)
	usersFp, err := os.OpenFile(*users, os.O_WRONLY | os.O_CREATE, 0644)
	handleError(err)
	config, err := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, *credentials)
	handleError(err)
	manager, err := mtproto.NewManager(config)
	handleError(err)
	conn, err := manager.LoadAuthentication()
	handleError(err)
	caller := mtproto.RPCaller{conn}

	// chats
	allchats, err := caller.MessagesGetAllChats(context.Background(), &mtproto.ReqMessagesGetAllChats{})
	handleError(err)
	marshaledChats, err := proto.Marshal(allchats)
	handleError(err)
	err = chatsFp.Truncate(0)
	handleError(err)
	_, err = chatsFp.WriteAt(marshaledChats, 0)
	handleError(err)
	fmt.Println(*chans, "is generated")

	// users
	contacts, err := caller.ContactsGetContacts(context.Background(), &mtproto.ReqContactsGetContacts{})
	handleError(err)
	marshaledContacts, err := proto.Marshal(contacts)
	handleError(err)
	err = usersFp.Truncate(0)
	handleError(err)
	_, err = usersFp.WriteAt(marshaledContacts, 0)
	handleError(err)
	fmt.Println(*users, "is generated")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("see", logfile)
		os.Exit(1)
	}
}
