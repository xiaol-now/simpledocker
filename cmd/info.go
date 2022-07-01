package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"simpledocker/container"
)

var InfoCmd = &cobra.Command{
	Use:   "info [CONTAINER]",
	Short: "Display container information",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		info := container.FindProcessInfo(args[0])
		if info != nil {
			b, _ := json.MarshalIndent(info, "", "\t")
			println(string(b))
		}
	},
}
