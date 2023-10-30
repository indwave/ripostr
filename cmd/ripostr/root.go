package main

import (
	"errors"
	log2 "github.com/indra-labs/indra/pkg/proc/log"
	"github.com/indra-labs/indra/pkg/util/appdata"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/wave/ripostr"
	"github.com/wave/ripostr/cmd/ripostr/forward"
	"github.com/wave/ripostr/cmd/ripostr/relay"
	"os"
	"strings"
)

var headerTxt = `ripostr (` + ripostr.SemVer + `) - Nostr Privacy.

The nostr swiss army knife.
`

var (
	cfgFile   string
	cfgSave   bool
	logsDir   string
	logsLevel string
	dataDir   string
	//network   string

	rootCmd = &cobra.Command{
		Use:   "ripostr",
		Short: "Nostr Privacy.",
		Long:  headerTxt,
	}
)

func init() {

	viper.SetEnvPrefix("RIPOSTR")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()

	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initData)

	cobra.OnFinalize(persistConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config-file", "C", "", "config file (default is $HOME/.ripostr/config.toml)")
	rootCmd.PersistentFlags().BoolVarP(&cfgSave, "config-save", "", false, "saves the config file with any eligible envs/flags passed")
	rootCmd.PersistentFlags().StringVarP(&logsDir, "logs-dir", "L", "", "logging directory (default is $HOME/.ripostr/logs)")
	rootCmd.PersistentFlags().StringVarP(&logsLevel, "logs-level", "", "info", "set logging level  off|fatal|error|warn|info|check|debug|trace")
	rootCmd.PersistentFlags().StringVarP(&dataDir, "data-dir", "D", appdata.Dir("ripostr", false), "data directory (default is $HOME/.ripostr/data)")
	//rootCmd.PersistentFlags().StringVarP(&network, "network", "N", "mainnet", "selects the network  mainnet|testnet|simnet")

	viper.BindPFlag("logs-dir", rootCmd.PersistentFlags().Lookup("logs-dir"))
	viper.BindPFlag("logs-level", rootCmd.PersistentFlags().Lookup("logs-level"))
	viper.BindPFlag("data-dir", rootCmd.PersistentFlags().Lookup("data-dir"))
	//viper.BindPFlag("network", rootCmd.PersistentFlags().Lookup("network"))

	relay.Init(rootCmd)
	forward.Init(rootCmd)
}

func initData() {

	if viper.GetString("data-dir") == "" {
		viper.Set("data-dir", appdata.Dir("ripostr", false))
	}
}

func initLogging() {

	if logsDir == "" {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		logsDir = home + "/.ripostr/logs"
	}

	log2.SetLogLevel(log2.GetLevelByString(viper.GetString("logs-level"), log2.Debug))
}

func initConfig() {

	if cfgFile == "" {

		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		cfgFile = home + "/.ripostr/config.toml"
	}

	viper.SetConfigFile(cfgFile)

	if _, err := os.Stat(cfgFile); errors.Is(err, os.ErrNotExist) {
		return
	}

	if err := viper.ReadInConfig(); err != nil {
		log.E.Ln("failed to read config file:", err)
		os.Exit(1)
	}

}

func persistConfig() {

	if !cfgSave {
		return
	}

	if err := viper.WriteConfig(); err != nil {
		log.E.Ln("failed to save config file:", err)
		os.Exit(1)
	}
}
