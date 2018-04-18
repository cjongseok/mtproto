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
)

// MTProto config
const (
	appVersion      = "0.0.1"
	deviceModel     = ""
	systemVersion   = ""
	language        = ""
)

// Proxy parameters
var (
	port, apiid int
	apihash, phone, addr, secrets string
	cfgFile string
)

// Cobra command
var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start proxy",
	Long: `Start a proxy at the given port.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// check required flags
		required := []string{"port", "apiid", "apihash", "phone", "addr", "secrets"}
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
		port = viper.GetInt("port")
		apiid = viper.GetInt("apiid")
		apihash = viper.GetString("apihash")
		phone = viper.GetString("phone")
		addr = viper.GetString("addr")
		secrets = viper.GetString("secrets")
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(startProxy(port, int32(apiid), apihash, phone, addr, secrets))
	},
}

func init() {
	viper.BindEnv("port")
	flags := startCmd.PersistentFlags()
	flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	flags.Int("port", 0, "port number the proxy runs at")
	flags.Int("apiid", 0, "Telegram API ID")
	flags.String("apihash", "", "Telegram API hash")
	flags.String("phone", "", "phone number registered to Telegram")
	flags.String("addr", "", "Telegram server address the proxy connects to")
	flags.String("secrets", "", "MTProto secrets")

	viper.BindPFlag("port", flags.Lookup("port"))
	viper.BindPFlag("apiid", flags.Lookup("apiid"))
	viper.BindPFlag("apihash", flags.Lookup("apihash"))
	viper.BindPFlag("phone", flags.Lookup("phone"))
	viper.BindPFlag("addr", flags.Lookup("addr"))
	viper.BindPFlag("secrets", flags.Lookup("secrets"))
}

func startProxy(port int, apiid int32, apihash, phone, addr, secrets string) int {
	slog.DisableLogging()
	config, err := mtproto.NewConfiguration(int32(apiid), apihash,
		appVersion, deviceModel, systemVersion, language, 0, 0, secrets)
	if err != nil {
		return invalidArgs
	}
	server := proxy.NewServer(port)
	err = server.Start(config, phone, addr)
	if err != nil {
		return failure
	}
	server.Wait()
	return normal
}
