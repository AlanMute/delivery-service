package cache

import "sync"

type Cache struct {
	data map[string]interface{}
	mu   sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]interface{}),
	}
}

func (c *Cache) Add(key string, data interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = data
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	d, ok := c.data[key]

	return d, ok
}
