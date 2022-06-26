package container

import (
	"os/exec"
	"strings"
)

type Workspace struct {
	containerId string
}

func NewWorkspace(containerId string) *Workspace {
	return &Workspace{containerId: containerId}
}

func (w *Workspace) MountOverlay(image string) error {
	_, readonlyPath, writePath, mergedPath, workPath := w.PathMountOrCreate()
	err := Decompress(ImageFilePath(image), readonlyPath)
	if err != nil {
		return err
	}
	dirs := []string{
		"lowerdir=" + readonlyPath,
		"upperdir=" + writePath,
		"workdir=" + workPath,
	}
	return exec.Command("mount", "-t", "overlay", "overlay", "-o", strings.Join(dirs, ","), mergedPath).Run()
}

func (w *Workspace) UmountOverlay() error {
	return exec.Command("umount", w.PathMountMerged()).Run()
}

//func (w *Workspace) MountVolume() error {
//
//}
