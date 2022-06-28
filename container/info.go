package container

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

type ProcessStatus string

const (
	RUNNING ProcessStatus = "Running"
	EXITED  ProcessStatus = "Exited"
)

type ProcessInfo struct {
	Id          string
	Name        string
	Env         []string
	Cmd         []string
	State       ProcessState
	Mount       []string // 挂载目录 /path:path
	GraphDriver ProcessGraphDriver
	Network     struct{}
}

type ProcessState struct {
	Status     ProcessStatus // 容器状态
	Pid        int           // 容器Pid
	StartedAt  string        // 最新的启动时间
	ExitCode   int           // 上一次停止状态
	FinishedAt string        // 上一次停止日期
}

type ProcessGraphDriver struct {
	Type        string
	ReadonlyDir string // 容器只读层路径
	WriteDir    string // 读写层
	WorkDir     string
	MergedDir   string
}

func SetProcessInfo(param RunParam, state ProcessState, w Workspace) {
	_, readonlyPath, writePath, mergedPath, workPath := w.PathMount()
	p := &ProcessInfo{
		Id:    param.Id,
		Name:  param.Name,
		Env:   param.Env,
		Cmd:   param.Cmd,
		State: state,
		GraphDriver: ProcessGraphDriver{
			Type:        "overlay",
			ReadonlyDir: readonlyPath,
			WriteDir:    writePath,
			WorkDir:     workPath,
			MergedDir:   mergedPath,
		},
	}
	encoder := json.NewEncoder(OpenProcessInfo(w.PathRuntimeInfo()))
	defer CloseProcessInfo()
	_ = encoder.Encode(p)
}

func FindProcessInfo(p ProcessPath) (pi *ProcessInfo) {
	_ = TryMkdir(p.PathRuntime())
	dirs, err := ioutil.ReadDir(p.PathRuntime())
	if err != nil {
		return nil
	}
	doesntExist := true
	for _, info := range dirs {
		if info.IsDir() && strings.HasPrefix(info.Name(), p.containerId) {
			p.containerId = info.Name()
			doesntExist = false
			break
		}
	}
	if doesntExist {
		return nil
	}
	decoder := json.NewDecoder(OpenProcessInfo(p.PathRuntimeInfo()))
	defer CloseProcessInfo()
	_ = decoder.Decode(pi)
	return nil
}

func ListProcessInfo() []string {
	dirs, err := ioutil.ReadDir(RuntimeContainerPath)
	if err != nil {
		return nil
	}
	var containers []string
	for _, info := range dirs {
		if info.IsDir() {
			containers = append(containers, info.Name())
		}
	}
	return containers
}

var _config *os.File

func OpenProcessInfo(filename string) *os.File {
	if _config != nil {
		return _config
	}
	_config, _ = os.OpenFile(filename, os.O_CREATE|os.O_RDWR, os.ModePerm)
	return _config
}

func CloseProcessInfo() {
	if _config != nil {
		_config.Close()
		_config = nil
	}
}
