package cache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	pokecache := Cache{entries: map[string]cacheEntry{}, mu: &sync.Mutex{}}
	go pokecache.reapLoop(interval)
	return pokecache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry, ok := c.entries[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	for tick := range ticker.C {
		c.mu.Lock()
		for key, ce := range c.entries {
			if ce.createdAt.Add(interval).Before(tick) {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}
