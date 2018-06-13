package main

import (
	"flag"
	"github.com/cjongseok/mtproto"
	"os"
	"fmt"
	"context"
	"github.com/cjongseok/slog"
	"encoding/json"
)

const (
	defaultNewKeyFile = "credentials.json"
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
)


var (
	credentials = flag.String("credentials", "credentials.json", "Credentials file. Default is credentials.json")
	out = flag.String("out", "out.json", "Output file. Default is out.json")
)


const logfile = "allchats.log"

func main() {
	flag.Parse()
	if *credentials == "" || *out == "" {
		flag.Usage()
		os.Exit(1)
	}

	slog.SetLogOutputAsFile(logfile)

	f, err := os.OpenFile(*out, os.O_WRONLY | os.O_CREATE, 644)
	handleError(err)
	config, err := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, *credentials)
	handleError(err)
	manager, err := mtproto.NewManager(config)
	handleError(err)
	conn, err := manager.LoadAuthentication()
	handleError(err)
	caller := mtproto.RPCaller{conn}
	resp, err := caller.MessagesGetAllChats(context.Background(), &mtproto.ReqMessagesGetAllChats{})
	handleError(err)
	//slog.StringifyIndent(resp, "  ")
	marshaled, err := json.Marshal(resp)
	handleError(err)
	err = f.Truncate(0)
	handleError(err)
	_, err = f.WriteAt(marshaled, 0)
	handleError(err)
	fmt.Println(*out, "is generated")
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		fmt.Println("see log file:", logfile)
		os.Exit(1)
	}
}
