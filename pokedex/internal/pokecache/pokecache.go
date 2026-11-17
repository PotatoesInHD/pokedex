package pokeapi
import "sync"
import "time"


type Cache struct {
	data	map[string]CacheEntry
	mu		sync.Mutex
}

type CacheEntry struct{
	createdAt 	time.Time
	val			[]byte
}


func NewCache(interval time.Duration) *Cache {
	pokeCache := Cache{data: make(map[string]CacheEntry)}
	go reapLoop(interval, &pokeCache)
	return &pokeCache
}


func reapLoop(interval time.Duration, cache *Cache) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
		for range ticker.C {
			cache.mu.Lock()
			for key, val := range cache.data {
				if interval < time.Since(val.createdAt) {
					delete(cache.data, key)
			}
		}
		cache.mu.Unlock()
	}
}


func (c *Cache) AddCache(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	elem := CacheEntry{}
	elem.val = val
	elem.createdAt = time.Now()
	c.data[key] = elem
}


func (c *Cache) GetCache(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if elem, ok := c.data[key]; ok {
		return elem.val, true
	}
	return nil, false
}
