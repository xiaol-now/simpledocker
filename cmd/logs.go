package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"os"
	"simpledocker/container"
)

var LogsContainerCmd = &cobra.Command{
	Use:   "logs [CONTAINER]",
	Short: "",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		info := container.FindProcessInfo(args[0])
		if info == nil {
			return
		}
		f, err := os.Open(info.Workspace().PathRuntimeLog())
		if err != nil {
			return
		}
		_, _ = io.Copy(os.Stdout, f)
	},
}
