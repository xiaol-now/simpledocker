package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"simpledocker/container"
	"text/tabwriter"
	"time"
)

var ImageListCmd = &cobra.Command{
	Use:   "image",
	Short: "List images",
	Run: func(cmd *cobra.Command, args []string) {
		_ = container.TryMkdir(container.ImagePath)
		w := tabwriter.NewWriter(os.Stdout, 15, 1, 2, ' ', 0)
		files, _ := ioutil.ReadDir(container.ImagePath)
		_, _ = fmt.Fprintln(w, "IMAGE\tCREATED\tSIZE")
		for _, f := range files {
			_, _ = fmt.Fprintf(w, "%s\t%s\t%.2fMB\n", f.Name(), timeFormat(f.ModTime()), sizeFormat(f.Size()))
		}
		_ = w.Flush()
	},
}

func timeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func sizeFormat(size int64) float32 {
	return float32(size) / 1024 / 1024
}

var ProcessListCmd = &cobra.Command{
	Use:   "ps",
	Short: "List containers",
	Run: func(cmd *cobra.Command, args []string) {
		for _, p := range container.ListContainerId() {
			println(p)
		}
	},
}
