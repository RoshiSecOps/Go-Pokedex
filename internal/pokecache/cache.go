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
