package container

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenContainerId(t *testing.T) {
	id := GenContainerId(10)
	assert.Len(t, id, 10)
}

func BenchmarkGenContainerId(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GenContainerId(32)
	}
}
