
# LRU Cache Package

The `lru_cache` package provides an in-memory Least Recently Used (LRU) Cache implementation in Go, using a combination of a doubly linked list and a `sync.Map` for efficient caching operations. This package is thread-safe and can be easily integrated into your Go applications.

---

## Installation

To install the package, use:

```bash
go get github.com/yingwei123/lru_cache
```

---

## Usage

### Create an LRU Cache

You can create an LRU cache by specifying the maximum capacity:

```go
import (
	"lru_cache"
)

func main() {
	cache := lru_cache.CreateLRUCache(3) // Capacity of 3

	// Add some items to the cache, note here the second parameter of Put can be any type
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Access a value
	value := cache.Get("key1")
	fmt.Println(value.(string)) // Output: value1, to access the value, cast it to the type of the value
}
```

### Methods

#### `Get(key string) any`

Retrieves the value associated with the given key. If the key exists, the associated node is marked as the most recently used (moved to the head). If the key does not exist, it returns `nil`.

Example:

```go
value := cache.Get("key1")
if value != nil {
	fmt.Println("Cache hit:", value)
} else {
	fmt.Println("Cache miss")
}
```

#### `Put(key string, value any)`

Adds a key-value pair to the cache. If the cache is full, the least recently used item is evicted.

Example:

```go
cache.Put("key3", "value3")
```

#### `GetCacheLength() int`

Returns the current size of the cache.

Example:

```go
length := cache.GetCacheLength()
fmt.Println("Cache size:", length)
```

---

## Example

```go
package main

import (
	"fmt"
	"lru_cache"
)

func main() {
	cache := lru_cache.CreateLRUCache(2)

	// Add some items to the cache
	cache.Put("key1", "value1")
	cache.Put("key2", "value2")

	// Access a value
	fmt.Println(cache.Get("key1")) // Output: value1

	// Add a new key (evicts the least recently used key)
	cache.Put("key3", "value3")

	// Check evicted key
	fmt.Println(cache.Get("key2")) // Output: nil (evicted)
	fmt.Println(cache.Get("key3")) // Output: value3
}
```

---

## Internals

### Doubly Linked List

The doubly linked list is used to track the usage order of cached items:
- **Head**: Most recently used item.
- **Tail**: Least recently used item.

### sync.Map

The `sync.Map` provides thread-safe operations for fast key-value lookups.

---

## License

This project is licensed under the MIT License. Feel free to use it as you see fit!

---

## Contributions

Contributions are welcome! Feel free to submit issues or pull requests to enhance this package.

---

## Author

Developed by Yingwei Li.
