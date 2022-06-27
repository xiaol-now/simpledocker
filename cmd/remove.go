package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"simpledocker/container"
)

var RemoveImageCmd = &cobra.Command{
	Use:   "rmi [IMAGE...]",
	Short: "Remove one or more images",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		files, _ := ioutil.ReadDir(container.ImagePath)
		for _, f := range files {
			if InArray(f.Name(), args) {
				_ = os.Remove(f.Name())
			}
		}
	},
}

var RemoveContainerCmd = &cobra.Command{
	Use:   "rm [CONTAINER...]",
	Short: "Remove one or more containers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//files, _ := ioutil.ReadDir(container.RuntimeContainerPath)
	},
}

func InArray(needle string, haystack []string) bool {
	for _, s := range haystack {
		if needle == s {
			return true
		}
	}
	return false
}
