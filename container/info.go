package container

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type ProcessStatus string

const (
	RUNNING ProcessStatus = "Running"
	EXITED  ProcessStatus = "Exited"
)

type ProcessInfo struct {
	Id          string             `json:"id"`
	Name        string             `json:"name"`
	Env         []string           `json:"env"`
	Cmd         []string           `json:"cmd"`
	State       ProcessState       `json:"state"`
	Volume      []string           `json:"volume"`
	GraphDriver ProcessGraphDriver `json:"graph_driver"`
	Network     struct{}           `json:"network"`
}

type ProcessState struct {
	Status     ProcessStatus `json:"status"`      // 容器状态
	Pid        int           `json:"pid"`         // 容器Pid
	StartedAt  time.Time     `json:"started_at"`  // 最新的启动时间
	ExitCode   int           `json:"exit_code"`   // 上一次停止状态
	FinishedAt *time.Time    `json:"finished_at"` // 上一次停止日期
}

type ProcessGraphDriver struct {
	Type        string `json:"type"`
	ReadonlyDir string `json:"readonly_dir"` // 容器只读层路径
	WriteDir    string `json:"write_dir"`    // 读写层
	WorkDir     string `json:"work_dir"`
	MergedDir   string `json:"merged_dir"`
}

func (p *ProcessInfo) Workspace() *Workspace {
	return NewWorkspace(p.Id, p.Volume)
}
func (p *ProcessInfo) Stop() {
	// 停止进程
}

func SetProcessInfo(param RunParam, w *Workspace, state ProcessState) {
	_, readonlyPath, writePath, mergedPath, workPath := w.PathMount()
	p := &ProcessInfo{
		Id:     param.Id,
		Name:   param.Name,
		Env:    param.Env,
		Cmd:    param.Cmd,
		Volume: w.volumes,
		State:  state,
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

func FindProcessInfo(id string) (pi *ProcessInfo) {
	p := ProcessPath{containerId: id}
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

func ListContainerId() []string {
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
