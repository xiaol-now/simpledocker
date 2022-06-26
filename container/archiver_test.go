package container

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDecompress(t *testing.T) {
	tests := []struct {
		name      string
		filename  string
		untarpath string
		tarfile   string
	}{
		{"Decompress 1", "./testdata/busybox.tar", "./testdata/untarpath", "./testdata/busybox_copy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.NoError(t, TryMkdir(tt.untarpath))
			assert.NoError(t, Decompress(tt.filename, tt.untarpath))
			assert.NoError(t, Compress(tt.untarpath, tt.tarfile))
			_ = os.RemoveAll(tt.untarpath)
			_ = os.RemoveAll(tt.tarfile + ".tar")
		})
	}
}
