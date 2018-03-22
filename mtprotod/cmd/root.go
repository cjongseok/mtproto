package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
	"os"
	"github.com/spf13/viper"
)

// Daemon exit codes
const (
	normal      = 0
	failure     = 1
	invalidArgs = 128
)

// Cobra command root
var rootCmd = &cobra.Command{
	Use: "mtprotod",
	Short: "Telegram MTProto proxy",
	Long: `Telegram Proxy. It has its own Telegram
MTProto implementation, proxy clients can calls Telegram
RPC procedures through it over gRPC.`,
	Version: "71.0.0 (MTProto Layer 71)",
	}

// Execute runs the root Cobra command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.AddCommand(startCmd)
}

func initConfig() {
	viper.AutomaticEnv()
}
