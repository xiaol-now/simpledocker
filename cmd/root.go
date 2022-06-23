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
}
