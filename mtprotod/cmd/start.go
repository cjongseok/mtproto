package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/cjongseok/mtproto"
	"github.com/cjongseok/mtproto/proxy"
	"os"
	"fmt"
	"strings"
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
	port int
	secrets string
)

// Cobra command
var startCmd = &cobra.Command{
	Use: "start",
	Short: "Start proxy",
	Long: `Start a proxy at the given port.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {

		// check required flags
		required := []string{"port", "secrets"}
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
		secrets = viper.GetString("secrets")
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		os.Exit(startProxy(port, secrets))
	},
}

func init() {
	viper.BindEnv("port")
	flags := startCmd.PersistentFlags()
	//flags.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")

	flags.Int("port", 0, "port number the proxy runs at")
	flags.String("secrets", "", "MTProto secrets")

	viper.BindPFlag("port", flags.Lookup("port"))
	viper.BindPFlag("secrets", flags.Lookup("secrets"))
}

func startProxy(port int, secrets string) int {
	//slog.DisableLogging()
	config, err := mtproto.NewConfiguration(appVersion, deviceModel, systemVersion, language, 0, 0, secrets)
	if err != nil {
		return invalidArgs
	}
	server := proxy.NewServer(port)
	err = server.Start(config)
	if err != nil {
		return failure
	}
	server.Wait()
	return normal
}
