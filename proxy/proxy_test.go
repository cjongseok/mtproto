package proxy

import (
	"flag"
	"fmt"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/slog"
	"golang.org/x/net/context"
	"testing"
)

const (
	port            = 11011
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
	sessionFileHome = ""
	telegramAddress = "149.154.167.50:443"
)

var (
	//apiId   = flag.Int("apiid", 0, "Telegram API id")
	//apiHash = flag.String("apihash", "", "Telegram API hash")
	//phone   = flag.String("phone", "", "Phone number including nation code")
	//addr    = flag.String("addr", "", "Preferred Telegram server address")
	secrets = flag.String("secrets", "", "MTProto secrets file")

	proxy  *Server
	client *Client
)

func beforeTest(t *testing.T) {
	flag.Parse()

	if proxy == nil {
		configuration, err := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, *secrets)
		handleError(t, err)
		proxy = NewServer(port)
		err = proxy.Start(configuration)
		handleError(t, err)
	}

	var err error
	client, err = NewClient(fmt.Sprintf("localhost:%d", port))
	handleError(t, err)
}

func afterTest(t *testing.T) {
	//proxy.Stop()
	//proxy = nil
	client = nil
}
func TestDialogs(t *testing.T) {
	beforeTest(t)
	defer afterTest(t)
	emptyPeer := &mtproto.TypeInputPeer{Value: &mtproto.TypeInputPeer_InputPeerEmpty{&mtproto.PredInputPeerEmpty{}}}
	dialogs, err := client.MessagesGetDialogs(context.Background(), &mtproto.ReqMessagesGetDialogs{
		OffsetDate: 0,
		OffsetId:   0,
		OffsetPeer: emptyPeer,
		Limit:      1,
	})
	handleError(t, err)
	switch casted := dialogs.Value.(type) {
	case *mtproto.TypeMessagesDialogs_MessagesDialogs:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogs, "  "))
	case *mtproto.TypeMessagesDialogs_MessagesDialogsSlice:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogsSlice, "  "))
	default:
		handleError(t, fmt.Errorf("unknown return: %T: %v", dialogs.Value, dialogs.Value))
	}
}

func TestUpdateMessage(t *testing.T) {
	// send a message

	// update the message
}

func TestListenOnUpdates(t *testing.T) {
	beforeTest(t)
	defer afterTest(t)
	stream, err := client.ListenOnUpdates(context.Background(), &ListenRequest{})
	handleError(t, err)
	limit := 5
	for updateCounter := 0; updateCounter < limit; {
		u, err := stream.Recv()
		handleError(t, err)
		fmt.Println("[client]: got an update:", u)
		updateCounter++
	}
}

func handleError(t *testing.T, err error) {
	if err != nil {
		fmt.Println("error:", err)
		t.FailNow()
		t.SkipNow()
	}
}
