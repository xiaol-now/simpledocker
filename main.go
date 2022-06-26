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
	//w := container.NewWorkspace("12c8c2cb3f83f49dcb9d026e9660132b")
	//err := w.UmountFS()

	//if err := cmd.RootCmd.Execute(); err != nil {
	//	Logger.Panic(err)
	//}
}
