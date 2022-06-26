package container

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestImportImage(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		imagePath string
	}{
		{"import 1", "./testdata/busybox.tar", "./testdata/image"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, TryMkdir(tt.imagePath))
			assert.NoError(t, ImportImage(tt.filename, tt.imagePath))
			_ = os.RemoveAll(tt.imagePath)
		})
	}
}
