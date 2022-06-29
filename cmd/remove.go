package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"simpledocker/container"
	. "simpledocker/logger"
	"strings"
)

var RemoveImageCmd = &cobra.Command{
	Use:   "rmi [IMAGE...]",
	Short: "Remove one or more images",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		files, _ := ioutil.ReadDir(container.ImagePath)
		for _, f := range files {
			if InPrefixArray(f.Name(), args) {
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
		Ids := container.ListContainerId()
		for _, arg := range args {
			if !InPrefixArray(arg, Ids) {
				continue
			}
			info := container.FindProcessInfo(arg)
			if info != nil {
				err := info.Workspace().Remove()
				if err != nil {
					Logger.Errorf("Remove Container fail: %s", info.Id)
				}
			}
		}
	},
}

func InPrefixArray(needle string, haystack []string) bool {
	for _, s := range haystack {
		if strings.HasPrefix(s, needle) {
			return true
		}
	}
	return false
}
