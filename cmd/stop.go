package cmd

import (
	"github.com/spf13/cobra"
	"simpledocker/container"
	. "simpledocker/logger"
)

var StopContainerCmd = &cobra.Command{
	Use:   "stop [CONTAINER]",
	Short: "Stop container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		info := container.FindProcessInfo(args[0])
		if info == nil {
			Logger.Fatalln("Container not found")
		}
		if info.State.Status == container.RUNNING {
			info.Stop()
		}
	},
}
