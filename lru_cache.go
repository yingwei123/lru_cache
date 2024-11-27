package lru_cache

import (
	dll "lru_cache/doubly_linked_list"
	"sync"
)

type LRUCache struct {
	Capacity int
	Cache    sync.Map
	DLL      *dll.DoublyLinkedList
}

// CreateLRUCache creates a new LRUCache with the given capacity.
func CreateLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Cache:    sync.Map{},
		DLL:      dll.CreateDoublyLinkedList(),
	}
}

// Get retrieves the value associated with the given key.
// If the key exists, it moves the node to the head of the DLL and returns the value.
// If the key does not exist, it returns nil.
func (lru *LRUCache) Get(key string) any {
	node, ok := lru.Cache.Load(key)
	if !ok {
		return nil
	}

	lru.DLL.RemoveNode(node.(*dll.Node))
	lru.DLL.AddToHead(node.(*dll.Node).Value, key)
	lru.Cache.Store(key, lru.DLL.Head)

	return node.(*dll.Node).Value
}

// GetCacheLength returns the number of items in the cache.
func (lru *LRUCache) GetCacheLength() int {
	count := 0

	lru.Cache.Range(func(key, value any) bool {
		count++
		return true
	})

	return count
}

// Put adds a key-value pair to the cache.
// If the cache is full, it removes the least recently used item.
func (lru *LRUCache) Put(key string, value any) {
	if lru.GetCacheLength() >= lru.Capacity {
		node := lru.DLL.RemoveTail()
		if node != nil {
			lru.Cache.Delete(node.Key)
		}
	}

	lru.DLL.AddToHead(value, key)
	lru.Cache.Store(key, lru.DLL.Head)
}
