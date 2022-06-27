package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Short: "This is a simple container",
}

func init() {
	RootCmd.AddCommand(RunCmd)
	RootCmd.AddCommand(InitContainerCmd)
	RootCmd.AddCommand(ImportCmd)
	RootCmd.AddCommand(InfoCmd)
	RootCmd.AddCommand(ImageCmd)
	RootCmd.AddCommand(RemoveContainerCmd)
	RootCmd.AddCommand(RemoveImageCmd)
	RootCmd.AddCommand(RemoveImageCmd)
}
