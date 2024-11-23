package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache() Cache {
	return Cache{
		cache: make(map[string]cacheEntry),
	}
}

func (c *Cache) Add(key string, value []byte) {
	c.cache[key] = cacheEntry{
		val:       value,
		createdAt: time.Now().UTC(),
	}
}