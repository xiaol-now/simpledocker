package cmd

import (
	"github.com/spf13/cobra"
	"simpledocker/container"
	. "simpledocker/logger"
)

var ImportCmd = &cobra.Command{
	Use:   "import image",
	Short: "Import image",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_ = container.TryMkdir(container.ImagePath)
		if err := container.ImportImage(args[0], container.ImagePath); err != nil {
			Logger.Fatal(err)
		}
	},
}
