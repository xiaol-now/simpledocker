package cmd

import "github.com/spf13/cobra"

var RemoveImageCmd = &cobra.Command{
	Use:   "rmi IMAGE",
	Short: "Remove one or more images",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var RemoveContainerCmd = &cobra.Command{
	Use:   "rm CONTAINER",
	Short: "Remove one or more containers",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
