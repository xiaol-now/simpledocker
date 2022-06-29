package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Short: "这是一个用golang实现的docker",
}

func init() {
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(InitContainerCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(ImageListCmd)
	RootCmd.AddCommand(ProcessListCmd)
	RootCmd.AddCommand(RemoveContainerCmd)
	RootCmd.AddCommand(RemoveImageCmd)
	RootCmd.AddCommand(StopContainerCmd)
}
