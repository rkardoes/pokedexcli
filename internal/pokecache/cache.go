package pokecache

import (
	"sync"
)

type Cache struct {
	mu    sync.RWMutex
	cache map[string][]byte
}

func NewCache() *Cache {
	c := &Cache{
		cache: make(map[string][]byte),
	}
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if key == "" {
		return
	}
	c.cache[key] = val
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	val, ok := c.cache[key]
	if !ok {
		return nil, false
	}
	return val, true
}
