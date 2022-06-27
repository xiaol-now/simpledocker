package container

type ProcessInfo struct {
	Id          string
	Name        string
	Env         string
	Cmd         []string
	State       ProcessState
	Mount       []string // 挂载目录 /path:path
	GraphDriver ProcessGraphDriver
	Network     struct{}
}

type ProcessState struct {
	Status     string // 容器状态
	Pid        int    // 容器Pid
	StartedAt  string // 最新的启动时间
	ExitCode   int    // 上一次停止状态
	FinishedAt string // 上一次停止日期
}

type ProcessGraphDriver struct {
	Type        string
	ReadonlyDir string // 容器只读层路径
	WriteDir    string // 读写层
	WorkDir     string // 挂载层
}

func FindProcessInfo(containerId string) *ProcessInfo {
	return nil
}
