package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	. "simpledocker/logger"
	"strings"
	"syscall"
)

var InitContainerCmd = &cobra.Command{
	Use:    "InitContainer",
	Hidden: true,
	Short:  "Init container process run user's process in container",
	Run: func(cmd *cobra.Command, args []string) {
		Logger.Debugf("Run init container: %#v", strings.Join(args, " "))

		SetUpMount()
		// 在系统环境 PATH中寻找命令的绝对路径
		path, err := exec.LookPath(args[0])
		if err != nil {
			Logger.Fatalf("cmd %s not found: %s", args[0], err)
		}
		err = syscall.Exec(path, args[0:], os.Environ())
		if err != nil {
			Logger.Fatalf("cmd exec fail: %+v\t%+v\t%+v", path, args[0:], os.Environ())
		}
	},
}

func SetUpMount() {
	err := syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err != nil {
		Logger.Fatalf("")
	}
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err = syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		Logger.Fatalf("mount proc, err: %v", err)
	}
}
