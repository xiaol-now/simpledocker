package cmd

import (
	"github.com/spf13/cobra"
	"io/ioutil"
	"simpledocker/container"
)

var ImageCmd = &cobra.Command{
	Use: "image",
	Run: func(cmd *cobra.Command, args []string) {
		_ = container.TryMkdir(container.ImagePath)
		files, _ := ioutil.ReadDir(container.ImagePath)
		for _, f := range files {
			println(f.Name(), f.ModTime(), f.Size())
		}
	},
}
