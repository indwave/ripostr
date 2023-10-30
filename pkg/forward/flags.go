package forward

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	relayFlag   = "relay"
	npubFlag    = "allow-npub"
	timeoutFlag = "retry-timeout"
	retryFlag   = "retry-max"
)

var (
	relays  []string
	npub    []string
	timeout uint
	retries uint
)

func InitFlags(cmd *cobra.Command) {

	cmd.Flags().StringArrayP(relayFlag, "",
		[]string{"wss://relay.damus.io", "wss://nos.lol"},
		"adds a forwarding relay.",
	)

	viper.BindPFlag(relayFlag, cmd.Flags().Lookup(relayFlag))

	cmd.Flags().StringArrayP(npubFlag, "",
		[]string{},
		"allows an npub to be forward to the list of relays.",
	)

	viper.BindPFlag(npubFlag, cmd.Flags().Lookup(npubFlag))

	cmd.Flags().UintVarP(&timeout, timeoutFlag, "", 5,
		"adds a timeout in seconds before a retry.",
	)

	viper.BindPFlag(timeoutFlag, cmd.Flags().Lookup(timeoutFlag))

	cmd.Flags().UintVarP(&retries, retryFlag, "", 3,
		"the maximum number of retries before giving up.",
	)

	viper.BindPFlag(retryFlag, cmd.Flags().Lookup(retryFlag))
}
