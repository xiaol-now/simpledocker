package container

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWorkspace_Mount(t *testing.T) {
	tests := []struct {
		name string
		Id   string
	}{
		{"mount container 1", GenContainerId(32)},
		{"mount container 2", GenContainerId(32)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := NewWorkspace(tt.Id, nil)
			assert.NoError(t, w.MountFS("busybox"))
			f, err := os.Stat(w.PathMountMerged())
			assert.NoError(t, err)
			assert.True(t, f.IsDir())
		})
	}
}
