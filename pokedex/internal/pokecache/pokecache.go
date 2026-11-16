package pokeapi

type Cache struct {
	data map[string][]byte
}


func NewCache() *Cache {
	pokeCache := Cache{data: make(map[string][]byte)}
	return &pokeCache
}


func (c *Cache) AddCache(key string, val []byte) {
	c.data[key] = val
}


func (c *Cache) GetCache(key string) ([]byte, bool) {
	if mapData, ok := c.data[key]; ok {
		return mapData, true
	}
	return nil, false
}
