package container

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	. "simpledocker/logger"
	"strings"
	"syscall"
)

var RunParam struct {
	Id      string
	Name    string
	TTY     bool
	Image   string
	Volumes []string
	Detach  bool     // 后台运行容器
	Memory  string   // 内存限制
	Env     []string // 环境变量
	Port    string
	Cmd     []string
}

func Run() {
	process := NewProcess()
	if err := process.Start(); err != nil {
		Logger.Fatalf("Container start:%v", err)
		return
	}
	_ = process.Wait()
}

func NewProcess() *exec.Cmd {
	args := append([]string{"InitContainer"}, RunParam.Cmd...)
	Logger.Debugf("Init container cmd: %s", strings.Join(RunParam.Cmd, " "))
	cmd := exec.Command("/proc/self/exe", args...)
	// 将容器进程跟宿主机的隔离机制
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // 隔离主机和域名
			syscall.CLONE_NEWPID | // 隔离 pid (process id)
			syscall.CLONE_NEWNS | // 隔离 mount 挂载点
			syscall.CLONE_NEWNET | // 隔离 Network
			syscall.CLONE_NEWIPC, // 隔离 System V IPC
	}
	if RunParam.TTY {
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
	} else {
		// TODO; 输出到日志文件
	}
	// 指定容器初始化后的工作目录
	w := NewWorkspace(RunParam.Id, RunParam.Volumes)
	err := w.MountFS(RunParam.Image)
	if err != nil {
		_ = w.Remove()
		Logger.Fatalln("MountFS: ", err)
	}
	cmd.Dir = w.PathMountMerged()
	cmd.Env = append(os.Environ(), RunParam.Env...)
	return cmd
}

func GenContainerId(n int) string {
	randBytes := make([]byte, n/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
