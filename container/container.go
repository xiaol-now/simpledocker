package container

import (
	"crypto/rand"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

type Process struct {
	conf   ProcessConfig
	reader *os.File
	writer *os.File
	proc   *exec.Cmd
}

func NewProcess(options ...Option) *Process {
	var conf ProcessConfig
	for _, opt := range options {
		opt(&conf)
	}
	readPipe, writePipe, _ := os.Pipe()
	return &Process{
		conf:   conf,
		reader: readPipe,
		writer: writePipe,
	}
}

func (p *Process) Create() error {
	cmd := exec.Command("/proc/self/exe", "InitContainer")
	// 将容器进程跟宿主机的隔离机制
	cmd.SysProcAttr = &syscall.SysProcAttr{
		//Cloneflags: syscall.CLONE_NEWUTS | // 隔离主机和域名
		//	syscall.CLONE_NEWPID | // 隔离 pid (process id)
		//	syscall.CLONE_NEWNS | // 隔离 mount 挂载点
		//	syscall.CLONE_NEWNET | // 隔离 Network
		//	syscall.CLONE_NEWIPC, // 隔离 System V IPC
	}
	if p.conf.TTY {
		cmd.Stdout = os.Stdout
		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
	} else {
		// TODO; 输出到日志文件
	}
	// 传递给容器进程的文件描述符 fd
	// 描述符id是3开始，0 Stdin, 1 Stdout, 2 Stderr
	cmd.ExtraFiles = []*os.File{p.reader}
	// 容器环境变量
	cmd.Env = append(os.Environ(), p.conf.Envs...)
	p.proc = cmd
	return cmd.Start()
}

func (p *Process) Limit() {

}

func (p *Process) Run() error {
	return p.proc.Wait()
}

func GenContainerId(n int) string {
	randBytes := make([]byte, n/2)
	_, _ = rand.Read(randBytes)
	return fmt.Sprintf("%x", randBytes)
}
