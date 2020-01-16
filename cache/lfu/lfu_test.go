package lfu

import "testing"

import "github.com/stretchr/testify/assert"

func TestLFU(t *testing.T) {
	lfu := Constructor(2)
	lfu.Put(3, 1)
	lfu.print()
	lfu.Put(2, 1)
	lfu.print()
	lfu.Put(2, 2)
	lfu.print()
	lfu.Put(4, 4)
	assert.Equal(t, lfu.Get(2), 2)
}

func BenchmarkLFU(b *testing.B) {

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
	}
}
