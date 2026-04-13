package pokecache

import (
	"sync"
	"time"
)

// Internal package to handle caching

type Cache struct {
	Entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

/*
func NewCache(interval time.Duration) {

}*/

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.Entries[key] = newEntry
}

func (c Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
