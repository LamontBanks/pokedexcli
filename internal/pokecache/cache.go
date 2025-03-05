package pokecache

import "time"

// Use a map to cache network results
type Cache struct {
	CacheEntries    map[string]CacheEntry
	RefreshInterval time.Duration
}

type CacheEntry struct {
	createdAt time.Time // When the entry was created
	val       []byte    // The value saved
}

// Functions ---

// Creates a new cache with the given refresh interval
func NewCache(refreshInterval time.Duration) Cache {
	return Cache{
		CacheEntries:    make(map[string]CacheEntry),
		RefreshInterval: refreshInterval,
	}
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
