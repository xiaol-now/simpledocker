package container

import (
	"os"
	"os/exec"
	"path"
	. "simpledocker/logger"
	"strings"
)

type Workspace struct {
	ProcessPath
	containerId string
	volumes     []string
}

func NewWorkspace(containerId string, volumes []string) *Workspace {
	return &Workspace{
		volumes:     volumes,
		ProcessPath: ProcessPath{containerId: containerId},
	}
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
		Logger.Fatalf("Volume mount: %s", err)
	}
}

//UnmountVolume 卸载 Volume
func (w *Workspace) UnmountVolume(volume string) {
	volumes := strings.Split(volume, ":")
	dst, work := path.Join(w.PathMountMerged(), volumes[1]), VolumeWorkTmpPath(volumes[0]) // 容器:workdir
	err := w.umountOverlay(dst)
	if err != nil {
		Logger.Fatalf("Volume unmount: %s", err)
	}
	_ = os.RemoveAll(work)
}

func (w *Workspace) mountOverlay(readonlyPath, writePath, workPath, mergedPath string) error {
	dirs := []string{
		"lowerdir=" + readonlyPath,
		"upperdir=" + writePath,
		"workdir=" + workPath,
	}
	args := []string{"-t", "overlay", "overlay", "-o", strings.Join(dirs, ","), mergedPath}
	Logger.Debugf("Workspace mount: \n\tmount %s", strings.Join(args, " "))
	cmd := exec.Command("mount", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (w *Workspace) umountOverlay(path string) error {
	Logger.Debugf("Workspace umount: \n\tumount %s", path)
	return exec.Command("umount", path).Run()
}

func (w *Workspace) Remove() error {
	err := w.UmountFS()
	if err != nil {
		return err
	}
	mountPath, _, _, _, _ := w.PathMount()
	_ = os.RemoveAll(mountPath)
	_ = os.RemoveAll(w.PathRuntimeInfo())
	_ = os.RemoveAll(w.PathRuntimeLog())
	return nil
}
