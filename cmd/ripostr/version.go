package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/wave/ripostr"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version number",
	Long:  `All software has versions. This is mine. Semver formatted.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ripostr.SemVer)
	},
}
