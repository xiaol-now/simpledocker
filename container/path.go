package container

import (
	"os"
	"path"
	. "simpledocker/logger"
	"strings"
)

const (
	// RuntimeContainerPath 容器运行时的 cgroup 路径
	RuntimeContainerPath = "/var/run/simpledocker"
	LibraryPath          = "/var/lib/simpledocker"
	ImagePath            = "/var/lib/simpledocker/image"
)

type ProcessPath struct {
}

func (w *Workspace) PathRuntime() string {
	return path.Join(RuntimeContainerPath, w.containerId)
}

func (w *Workspace) RuntimeInfo() string {
	return path.Join(RuntimeContainerPath, w.containerId, "info.json")
}
func (w *Workspace) RuntimeLog() string {
	return path.Join(RuntimeContainerPath, w.containerId, "container.log")
}

func (w *Workspace) PathMountMerged() string {
	_, _, _, mergedPath, _ := w.PathMount()
	return mergedPath
}

func (w *Workspace) PathMount() (mountPath, readonlyPath, writePath, mergedPath, workPath string) {
	dir := path.Join(LibraryPath, w.containerId)
	return dir,
		path.Join(dir, "lower"),
		path.Join(dir, "upper"),
		path.Join(dir, "merged"),
		path.Join(dir, "work")
}

func (w *Workspace) PathMountOrCreate() (mountPath, readonlyPath, writePath, mergedPath, workPath string) {
	mountPath, readonlyPath, writePath, mergedPath, workPath = w.PathMount()
	TryMkdirOrFail(mountPath)
	TryMkdirOrFail(readonlyPath)
	TryMkdirOrFail(writePath)
	TryMkdirOrFail(mergedPath)
	TryMkdirOrFail(workPath)
	return
}

func VolumeWorkTmpPath(volumePath string) string {
	return path.Join("/tmp", strings.ReplaceAll(volumePath, string(os.PathSeparator), "_"))
}

func ImageFilePath(image string) string {
	return path.Join(ImagePath, image)
}

func TryMkdirOrFail(path string) {
	if err := TryMkdir(path); err != nil {
		Logger.Fatalf("mkdir %#v err: %#v", path, err)
	}
}

func TryMkdir(path string) error {
	if !ExistDir(path) {
		return os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// ExistDir 判断目录是否存在
func ExistDir(dirname string) bool {
	fi, err := os.Stat(dirname)
	return (err == nil || os.IsExist(err)) && fi.IsDir()
}

func ExistFile(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}
