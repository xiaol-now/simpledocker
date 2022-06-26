package container

import (
	"os"
	"os/exec"
	"path"
	. "simpledocker/logger"
	"strings"
)

type Workspace struct {
	containerId string
	volumes     []string
}

func NewWorkspace(containerId string, volumes []string) *Workspace {
	return &Workspace{containerId: containerId, volumes: volumes}
}

// MountFS 挂载文件系统
func (w *Workspace) MountFS(image string) error {
	_, readonlyPath, writePath, mergedPath, workPath := w.PathMountOrCreate()
	err := Decompress(ImageFilePath(image), readonlyPath)
	if err != nil {
		return err
	}
	err = w.mountOverlay(readonlyPath, writePath, workPath, mergedPath)
	if err != nil {
		return err
	}
	for _, volume := range w.volumes {
		w.MountVolume(volume)
	}
	return nil
}

// UmountFS 卸载文件系统
func (w *Workspace) UmountFS() error {
	for _, volume := range w.volumes {
		w.UnmountVolume(volume)
	}
	return w.umountOverlay(w.PathMountMerged())
}

// MountVolume 挂载 Volume
func (w *Workspace) MountVolume(volume string) {
	volumes := strings.Split(volume, ":")
	if len(volumes) != 2 {
		return
	}
	src, dst, work := volumes[0], path.Join(w.PathMountMerged(), volumes[1]), VolumeWorkTmpPath(volumes[0]) // 宿主:容器:workdir
	TryMkdirOrFail(src)
	TryMkdirOrFail(dst)
	TryMkdirOrFail(work)
	err := w.mountOverlay(dst, src, VolumeWorkTmpPath(src), dst)
	if err != nil {
		Logger.Errorf("Volume mount: %s", err)
	}
}

//UnmountVolume 卸载 Volume
func (w *Workspace) UnmountVolume(volume string) {
	volumes := strings.Split(volume, ":")
	err := w.umountOverlay(volumes[0])
	if err != nil {
		Logger.Errorf("Volume unmount: %s", err)
	}
}

func (w *Workspace) mountOverlay(readonlyPath, writePath, workPath, mergedPath string) error {
	dirs := []string{
		"lowerdir=" + readonlyPath,
		"upperdir=" + writePath,
		"workdir=" + workPath,
	}
	return exec.Command("mount", "-t", "overlay", "overlay", "-o", strings.Join(dirs, ","), mergedPath).Run()
}

func (w *Workspace) umountOverlay(path string) error {
	return exec.Command("umount", path).Run()
}

func (w *Workspace) Delete() error {
	err := w.UmountFS()
	if err != nil {
		return err
	}
	mountPath, _, _, _, _ := w.PathMount()
	return os.RemoveAll(mountPath)
}
