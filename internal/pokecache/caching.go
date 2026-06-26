package pokecache

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Mut      sync.RWMutex
	Registry map[string]CacheEntry
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		interval: interval,
		Registry: make(map[string]CacheEntry),
	}
	go c.reapLoop()
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.Mut.Lock()
	defer c.Mut.Unlock()
	c.Registry[key] = CacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mut.RLock()
	defer c.Mut.RUnlock()

	val, exists := c.Registry[key]
	return val.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		t := time.Now()
		c.Mut.Lock()
		for key, entry := range c.Registry {
			age := t.Sub(entry.createdAt)
			if age > c.interval {
				delete(c.Registry, key)
			}
		}
		c.Mut.Unlock()
	}
}
