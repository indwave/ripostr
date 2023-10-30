package relay

import (
	log2 "github.com/indra-labs/indra/pkg/proc/log"
	"github.com/spf13/cobra"
)

var (
	log   = log2.GetLogger()
	check = log.E.Chk
)

func init() {
	//storage.InitFlags(serveCmd)
	//p2p.InitFlags(serveCmd)
	//rpc.InitFlags(serveCmd)
}

func Init(c *cobra.Command) {

	relayCommand.AddCommand(serveCmd)

	c.AddCommand(relayCommand)
}

var relayCommand = &cobra.Command{
	Use:   "relay",
	Short: "run and manage your relay",
	Long:  `run and manage your relay`,
}
