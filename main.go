package main

import (
	"simpledocker/cmd"
	. "simpledocker/logger"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		Logger.Panic(err)
	}
}

/*const (
	// 挂载了 memory subsystem的hierarchy的根目录位置
	cgroupMemoryHierarchyMount = "/sys/fs/cgroup/memory"
)

func main2() {
	println(os.Args[0])
	if os.Args[0] == "/proc/self/exe" {
		//容器进程
		fmt.Printf("current pid %d\t ppid:%d \n", syscall.Getpid(), syscall.Getppid())

		cmd := exec.Command("sh", "-c", "sleep infinity")
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			panic(err)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		panic(err)
	}
	// 得到 fork出来进程映射在外部命名空间的pid
	fmt.Printf("fork pid: %+v \n\n", cmd.Process.Pid)

	// 创建子cgroup
	newCgroup := path.Join(cgroupMemoryHierarchyMount, "cgroup-demo-memory")
	// 将容器进程放到子cgroup中
	if err := ioutil.WriteFile(path.Join(newCgroup, "tasks"), []byte(strconv.Itoa(cmd.Process.Pid)), 0644); err != nil {
		panic(err)
	}
	// 限制cgroup的内存使用
	if err := ioutil.WriteFile(path.Join(newCgroup, "memory.limit_in_bytes"), []byte("100m"), 0644); err != nil {
		panic(err)
	}
	cmd.Process.Wait()
}
*/
