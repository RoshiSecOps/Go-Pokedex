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

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		Entries: make(map[string]cacheEntry),
	}
	go newCache.reapLoop(interval)
	return &newCache
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, entry := range c.Entries {
			if time.Since(entry.createdAt) > interval {
				delete(c.Entries, key)
			}
		}
		c.mu.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	newEntry := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.Entries[key] = newEntry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.Entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}
