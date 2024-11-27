package lru_cache

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache(t *testing.T) {
	t.Run("test put", testPut)
	t.Run("test put with capacity", testPutWithCapacity)
}

func testPut(t *testing.T) {
	lru := CreateLRUCache(2)
	lru.Put("1", 1)
	lru.Put("2", 2)

	assert.Equal(t, 2, lru.GetCacheLength())
}

func testPutWithCapacity(t *testing.T) {
	lru := CreateLRUCache(1)
	lru.Put("1", 1)
	lru.Put("2", 2)

	assert.Equal(t, 1, lru.GetCacheLength())
	assert.Equal(t, 2, lru.Get("2"))
	assert.Nil(t, lru.Get("1"))
}
