package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path"
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
		for _, arg := range args {
			for _, f := range files {
				Logger.Infoln(f.Name(), arg, strings.HasPrefix(f.Name(), arg))
				if strings.HasPrefix(f.Name(), arg) {
					_ = os.Remove(path.Join(container.ImagePath, f.Name()))
				}
			}
		}
	},
}

var RemoveContainerCmd = &cobra.Command{
	Use:   "rm [CONTAINER...]",
	Short: "Remove one or more containers",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
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
