package forward

import (
	log2 "github.com/indra-labs/indra/pkg/proc/log"
	"github.com/spf13/cobra"
	"github.com/wave/ripostr/pkg/forward"
)

var (
	log   = log2.GetLogger()
	check = log.E.Chk
)

func init() {
	forward.InitFlags(forwardCommand)
}

func Init(c *cobra.Command) {

	forwardCommand.AddCommand(serveCmd)

	c.AddCommand(forwardCommand)
}

var forwardCommand = &cobra.Command{
	Use:   "forward",
	Short: "delegate posting notes to a series of relays",
	Long:  `delegate posting notes to a series of relays`,
}
