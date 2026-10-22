package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	mapCache      map[string]cacheEntry
	mu            sync.Mutex
	checkInterval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		mapCache:      map[string]cacheEntry{},
		checkInterval: interval,
	}
	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := c.mapCache[key]
	entry.createdAt = time.Now()
	entry.val = val
	c.mapCache[key] = entry
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if entry, ok := c.mapCache[key]; ok {
		return entry.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.checkInterval)
	defer ticker.Stop()
	for range ticker.C {
		c.mu.Lock()
		for key, val := range c.mapCache {
			if time.Since(val.createdAt) > c.checkInterval {
				delete(c.mapCache, key)
			}
		}
		c.mu.Unlock()
	}
}
