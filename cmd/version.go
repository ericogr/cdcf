package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Version information",
	Long:  "Get version information.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Create docker config file v0.0.1 (by EricoGR) - https://github.com/ericogr/cdcf")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
