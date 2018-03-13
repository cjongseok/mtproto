package proxy

import (
	"flag"
	"fmt"
	"github.com/cjongseok/mtproto/mtp"
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
	apiId   = flag.Int("apiid", 0, "Telegram API id")
	apiHash = flag.String("apihash", "", "Telegram API hash")
	phone   = flag.String("phone", "", "Phone number including nation code")
	addr    = flag.String("addr", "", "Preferred Telegram server address")

	proxy  *Server
	client *Client
)

func beforeTest(t *testing.T) {
	flag.Parse()
	if apiId == nil || apiHash == nil || phone == nil {
		t.SkipNow()
	}
	fmt.Printf("apiid: %v\napihash: %v\nphone: %s\naddr: %s\n", *apiId, *apiHash, *phone, *addr)

	if proxy == nil {
		configuration, err := mtp.NewConfiguration(int32(*apiId), *apiHash, appVersion, deviceModel, systemVersion, language, sessionFileHome, 0, 0)
		handleError(t, err)
		proxy = NewServer(port)
		err = proxy.ConnectToTelegram(configuration, *phone, *addr)
		handleError(t, err)
		err = proxy.ServeRPC()
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
	emptyPeer := &mtp.TypeInputPeer{&mtp.TypeInputPeer_InputPeerEmpty{&mtp.PredInputPeerEmpty{}}}
	dialogs, err := client.MessagesGetDialogs(context.Background(), &mtp.ReqMessagesGetDialogs{
		OffsetDate: 0,
		OffsetId:   0,
		OffsetPeer: emptyPeer,
		Limit:      1,
	})
	handleError(t, err)
	switch casted := dialogs.Value.(type) {
	case *mtp.TypeMessagesDialogs_MessagesDialogs:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogs, "  "))
	case *mtp.TypeMessagesDialogs_MessagesDialogsSlice:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogsSlice, "  "))
	default:
		handleError(t, fmt.Errorf("unknown return: %T: %v", dialogs.Value, dialogs.Value))
	}
}

func TestUpdateStream(t *testing.T) {
	beforeTest(t)
	defer afterTest(t)
	stream, err := client.ListenOnUpdates(context.Background(), &ListenRequest{})
	handleError(t, err)
	//handleError(t, fmt.Errorf("stream init failure: %v", err))
	limit := 5
	for updateCounter := 0; updateCounter < limit; {
		u, err := stream.Recv()
		handleError(t, err)
		//handleError(t, fmt.Errorf("stream recv failure: %v", err))
		fmt.Println("got an update:", u)
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
