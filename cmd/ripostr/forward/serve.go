package forward

import (
	"context"
	"github.com/indra-labs/indra/pkg/interrupt"
	"github.com/wave/ripostr"

	log2 "github.com/indra-labs/indra/pkg/proc/log"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serves an instance of the relay",
	Long:  `Serves an instance of the relay.`,
	Run: func(cmd *cobra.Command, args []string) {

		log.I.Ln("-- ", log2.App.Load(), ripostr.SemVer, "- Nostr Privacy. --")

		_, cancel := context.WithCancel(context.Background())
		interrupt.AddHandler(cancel)

		//// Seed //
		//
		//go seed.Run(ctx)
		//
		//select {
		//case <-seed.WhenStartFailed():
		//	log.I.Ln("stopped")
		//case <-seed.WhenShutdown():
		//	log.I.Ln("shutdown complete")
		//}

		log.I.Ln("-- fin --")
	},
}
