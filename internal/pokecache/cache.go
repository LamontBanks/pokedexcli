package pokecache

import (
	"sync"
	"time"
)

// Use a map to cache network results
type Cache struct {
	CacheEntries map[string]CacheEntry
	mu           *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time // When the entry was created
	val       []byte    // The value saved
}

// Functions ---

// Creates a new cache with the given refresh interval
func NewCache(refreshInterval time.Duration) Cache {
	cache := Cache{
		CacheEntries: make(map[string]CacheEntry),
		mu:           &sync.Mutex{},
	}

	// Delete expired items from the cache
	go cache.reapLoop(refreshInterval)

	return cache
}

// Add a new entry to the cache
func (c *Cache) Add(newKey string, newVal []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	creationTime := time.Now().UTC()
	c.CacheEntries[newKey] = CacheEntry{
		val:       newVal,
		createdAt: creationTime,
	}
}

// Retrieve an entry from the cache
func (c *Cache) Get(existingKey string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cacheEntry, exists := c.CacheEntries[existingKey]
	return cacheEntry.val, exists
}

// Remove cache centries older than the refreshInterval
func (c *Cache) reapLoop(interval time.Duration) {
	// Use Ticker to get timestamps every interval
	ticker := time.NewTicker(interval)

	for tick := range ticker.C {
		c.reap(tick, interval)
	}
}

func (c *Cache) reap(latestTickerTimestamp time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	for key, cacheEntry := range c.CacheEntries {
		if cacheEntry.createdAt.Before(latestTickerTimestamp.Add(-interval)) {
			delete(c.CacheEntries, key)
		}
	}
}
