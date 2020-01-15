package lfu

import "testing"

import "github.com/stretchr/testify/assert"

func TestLFU(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(1, 1)
	lfu.Put(2, 2)
	assert.Equal(t, lfu.Get(1), 1)
	lfu.Put(3, 3)
	assert.Equal(t, lfu.Get(2), -1)
	assert.Equal(t, lfu.Get(3), 3)
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(1), -1)
	assert.Equal(t, lfu.Get(3), 3)
	assert.Equal(t, lfu.Get(4), 4)
}

func BenchmarkLFU(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
