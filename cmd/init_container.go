package cmd

import (
	"github.com/spf13/cobra"
	"io"
	"os"
	"simpledocker/logger"
	"strings"
)

var InitContainerCmd = &cobra.Command{
	Use:    "InitContainer",
	Hidden: true,
	Short:  "Init container process run user's process in container",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("Run init container", map[string]interface{}{
			"args": args,
		})
	},
}

func ReadUserCommand() []string {
	//0 Stdin, 1 Stdout, 2 Stderr, 3是从宿主机传递过来的文件描述符
	pipe := os.NewFile(uintptr(3), "pipe")
	by, err := io.ReadAll(pipe)
	if err != nil {
		logger.Error("read pipe", map[string]interface{}{"err": err})
		return nil
	}
	return strings.Split(string(by), " ")
}
