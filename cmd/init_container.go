package cmd

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"path/filepath"
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

		err := SetUpMount()
		if err != nil {
			Logger.Fatalf("SetUpMount: %s", err)
		}
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

func SetUpMount() error {
	err := pivotRoot()
	if err != nil {
		logrus.Errorf("pivot root, err: %v", err)
		return err
	}

	// systemd 加入linux之后, mount namespace 就变成 shared by default, 所以你必须显示
	//声明你要这个新的mount namespace独立。
	err = syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err != nil {
		return err
	}
	//mount proc
	defaultMountFlags := syscall.MS_NOEXEC | syscall.MS_NOSUID | syscall.MS_NODEV
	err = syscall.Mount("proc", "/proc", "proc", uintptr(defaultMountFlags), "")
	if err != nil {
		logrus.Errorf("mount proc, err: %v", err)
		return err
	}
	// mount temfs, temfs是一种基于内存的文件系统
	err = syscall.Mount("tmpfs", "/dev", "tmpfs", syscall.MS_NOSUID|syscall.MS_STRICTATIME, "mode=755")
	if err != nil {
		logrus.Errorf("mount tempfs, err: %v", err)
		return err
	}

	return nil
}

// 改变当前root的文件系统
func pivotRoot() error {
	root, err := os.Getwd()
	if err != nil {
		return err
	}
	logrus.Infof("current location is %s", root)

	// systemd 加入linux之后, mount namespace 就变成 shared by default, 所以你必须显示
	//声明你要这个新的mount namespace独立。
	err = syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
	if err != nil {
		return err
	}
	// 为了使当前root的老 root 和新 root 不在同一个文件系统下，我们把root重新mount了一次
	// bind mount是把相同的内容换了一个挂载点的挂载方法
	if err := syscall.Mount(root, root, "bind", syscall.MS_BIND|syscall.MS_REC, ""); err != nil {
		return fmt.Errorf("mount rootfs to itself error: %v", err)
	}
	// 创建 rootfs/.pivot_root 存储 old_root
	pivotDir := filepath.Join(root, ".pivot_root")
	_, err = os.Stat(pivotDir)
	if err != nil && os.IsNotExist(err) {
		if err := os.Mkdir(pivotDir, 0777); err != nil {
			return err
		}
	}
	// pivot_root 到新的rootfs, 现在老的 old_root 是挂载在rootfs/.pivot_root
	// 挂载点现在依然可以在mount命令中看到
	if err := syscall.PivotRoot(root, pivotDir); err != nil {
		return fmt.Errorf("pivot_root %v", err)
	}
	// 修改当前的工作目录到根目录
	if err := syscall.Chdir("/"); err != nil {
		return fmt.Errorf("chdir / %v", err)
	}

	pivotDir = filepath.Join("/", ".pivot_root")
	// unmount rootfs/.pivot_root
	if err := syscall.Unmount(pivotDir, syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount pivot_root dir %v", err)
	}
	// 删除临时文件夹
	return os.Remove(pivotDir)
}
