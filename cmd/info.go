package cmd

import "github.com/spf13/cobra"

var InfoCmd = &cobra.Command{
	Use:   "info [CONTAINER]",
	Short: "Display container information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

	},
}
