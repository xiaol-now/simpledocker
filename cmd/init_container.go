package cmd

import (
	"github.com/spf13/cobra"
	. "simpledocker/logger"
)

var InitContainerCmd = &cobra.Command{
	Use:    "InitContainer",
	Hidden: true,
	Short:  "Init container process run user's process in container",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Infof("Run init container: %#v", args)
	},
}
