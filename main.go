package main

import (
	"simpledocker/container"
	. "simpledocker/logger"
)

func main() {
	// 挂载目录
	_ = container.TryMkdir(container.ImagePath)
	if err := container.ImportImage("container/testdata/busybox.tar", container.ImagePath); err != nil {
		Logger.Fatal(err)
	}
	volumnes := []string{
		"/tmp/volume1:/root",
		"/tmp/volume2:/tmp/volume2",
	}
	w := container.NewWorkspace(container.GenContainerId(32), volumnes)
	err := w.MountFS("busybox")
	if err != nil {
		Logger.Panic(err)
	}

	// 解除挂载
	volumnes = []string{
		"/tmp/volume1:/root",
		"/tmp/volume2:/tmp/volume2",
	}
	w = container.NewWorkspace("3ac1da1f871cc491f90cefcf80811d01", volumnes)
	err = w.UmountFS()
	if err != nil {
		Logger.Panic(err)
	}

	//if err := cmd.RootCmd.Execute(); err != nil {
	//	Logger.Panic(err)
	//}
}
