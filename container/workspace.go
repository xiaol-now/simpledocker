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

func (w *Workspace) MountOverlay() error {
	tryMkdir(w.PathMountReadonly())
	tryMkdir(w.PathMountWrite())
	tryMkdir(w.PathMountMerged())
	dirs := []string{
		"lowerdir=" + w.PathMountReadonly(),
		"upperdir=" + w.PathMountWrite(),
		"workdir=" + w.PathMountWork(),
	}
	return exec.Command("mount", "-t", "overlay", "overlay", "-o", strings.Join(dirs, ","), w.PathMountMerged()).Run()
}

func (w *Workspace) UmountOverlay() error {
	return exec.Command("umount", w.PathMountMerged()).Run()
}

//func (w *Workspace) MountVolume() error {
//
//}
