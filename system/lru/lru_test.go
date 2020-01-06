package lru

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLru(t *testing.T) {

	lru := Constructor(2)
	lru.Put(2, 1)
	lru.Put(2, 2)
	assert.Equal(t, lru.Get(2), 2)
	lru.Put(1, 1) // evicts key 2
	lru.Put(4, 1)
	assert.Equal(t, lru.Get(2), -1)
}
