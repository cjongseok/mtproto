package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/mtproto/proxy"
	"os"
	"fmt"
	"strings"
	"github.com/cjongseok/slog"
	"golang.org/x/net/context"
)

// Proxy client parameters
var (
	proxyAddr string
	//clientCfgFile string
)

// Cobra command
var testCmd = &cobra.Command{
	Use: "test",
	Short: "Test proxy",
	Long: `Test the proxy by getting recent Telegram dialogs.
If it is normal, you can see the dialogs in JSON.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// check required flags
		required := []string{"proxy"}
		var notset string
		for _, flagStr := range required {
			f := viper.Get(flagStr)
			switch casted := f.(type) {
			case string:
				if casted == "" {
					notset = fmt.Sprintf("%s, --%s", notset, flagStr)
				}
			case int:
				if casted == 0 {
					notset = fmt.Sprintf("%s, --%s", notset, flagStr)
				}
			}
		}
		if notset != "" {
			notset = strings.Replace(notset, ", ", "", 1)
			return fmt.Errorf("%s are required", notset)

		}

		// get flags
		proxyAddr = viper.GetString("proxy")
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(dialogs(proxyAddr))
	},
}

func init() {
	viper.BindEnv("proxy")
	flags := testCmd.PersistentFlags()
	//flags.StringVar(&clientCfgFile, "client_config", "", "config file (default is $HOME/.cobra.yaml)")

	flags.String("proxy", "", "Proxy address")
	viper.BindPFlag("proxy", flags.Lookup("proxy"))
}

func dialogs(addr string) int {
	client, err := proxy.NewClient(fmt.Sprintf("localhost:%d", port))
	emptyPeer := &mtproto.TypeInputPeer{Value: &mtproto.TypeInputPeer_InputPeerEmpty{&mtproto.PredInputPeerEmpty{}}}
	dialogs, err := client.MessagesGetDialogs(context.Background(), &mtproto.ReqMessagesGetDialogs{
		OffsetDate: 0,
		OffsetId:   0,
		OffsetPeer: emptyPeer,
		Limit:      1,
	})
	if err  != nil {
		fmt.Println("test failure:", err)
		return failure
	}
	switch casted := dialogs.Value.(type) {
	case *mtproto.TypeMessagesDialogs_MessagesDialogs:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogs, "  "))
	case *mtproto.TypeMessagesDialogs_MessagesDialogsSlice:
		fmt.Println(slog.StringifyIndent(casted.MessagesDialogsSlice, "  "))
	default:
		fmt.Println("got unknown types of dialogs:")
		fmt.Println(slog.StringifyIndent(dialogs.Value, "  "))
	}
	return normal
}
