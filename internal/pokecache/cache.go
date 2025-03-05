package pokecache

import (
	"sync"
	"time"
)

// Use a map to cache network results
type Cache struct {
	CacheEntries    map[string]CacheEntry
	RefreshInterval time.Duration
	mu              *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time // When the entry was created
	val       []byte    // The value saved
}

// Functions ---

// Creates a new cache with the given refresh interval
func NewCache(refreshInterval time.Duration) Cache {
	cache := Cache{
		CacheEntries:    make(map[string]CacheEntry),
		RefreshInterval: refreshInterval,
		mu:              &sync.Mutex{},
	}

	// Delete expired items from the cache
	go cache.reapLoop()

	return cache
}

// Add a new entry to the cache
func (c *Cache) Add(newKey string, newVal []byte) {
	_, exists := c.CacheEntries[newKey]
	if !exists {
		creationTime := time.Now()
		c.CacheEntries[newKey] = CacheEntry{
			val:       newVal,
			createdAt: creationTime,
		}
	}
}

// Retrieve an entry from the cache
func (c *Cache) Get(existingKey string) ([]byte, bool) {
	cacheEntry, exists := c.CacheEntries[existingKey]

	if !exists {
		return []byte{}, exists
	}

	return cacheEntry.val, exists
}

// Remove cache centries older than the refreshInterval
func (c *Cache) reapLoop() {
	// Use Ticker to get timestamps every interval
	ticker := time.NewTicker(c.RefreshInterval)

	for latestTimestamp := range ticker.C {
		// For now, print all entries older than current time - interval
		for key, cacheEntry := range c.CacheEntries {
			if cacheEntry.createdAt.Before(latestTimestamp.Add(-c.RefreshInterval)) {
				c.mu.Lock()

				delete(c.CacheEntries, key)

				c.mu.Unlock()
			}
		}
	}
}
