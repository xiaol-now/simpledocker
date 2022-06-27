package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWorkspace_Mount(t *testing.T) {
	// 挂载目录
	assert.NoError(t, TryMkdir(ImagePath))
	assert.NoError(t, ImportImage("container/testdata/busybox.tar", ImagePath))
	volumnes := []string{
		"/tmp/volume1:/root",
		"/tmp/volume2:/tmp/volume2",
	}
	tests := []struct {
		name string
		Id   string
	}{
		{"mount container 1", GenContainerId(32)},
		{"mount container 2", GenContainerId(32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorkspace(tt.Id, volumnes)
			assert.NoError(t, w.MountFS("busybox"))
			assert.NoError(t, w.Remove())
		})
	}
}
