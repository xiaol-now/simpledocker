package container

import (
	"os"
	"path"
	"simpledocker/logger"
)

const (
	// RuntimeContainerPath 容器运行时的 cgroup 路径
	RuntimeContainerPath = "/var/run/simpledocker"
	LibraryPath          = "/var/lib/simpledocker"
	ImagePath            = "/var/lib/simpledocker/image"
)

func (w *Workspace) PathRuntime() string {
	return path.Join(RuntimeContainerPath, w.containerId)
}

func (w *Workspace) RuntimeInfo() string {
	return path.Join(RuntimeContainerPath, w.containerId, "info.json")
}
func (w *Workspace) RuntimeLog() string {
	return path.Join(RuntimeContainerPath, w.containerId, "container.log")
}

func (w *Workspace) PathMount() string {
	return path.Join(LibraryPath, w.containerId)
}

func (w *Workspace) PathMountReadonly() string {
	return path.Join(w.PathMount(), "lower")
}

func (w *Workspace) PathMountWrite() string {
	return path.Join(w.PathMount(), "upper")
}

func (w *Workspace) PathMountMerged() string {
	return path.Join(w.PathMount(), "merged")
}

func (w *Workspace) PathMountWork() string {
	return path.Join(w.PathMount(), "work")
}

func tryMkdir(path string) {
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Fatal("mkdir"+path, map[string]interface{}{"err": err})
		}
	}
}
