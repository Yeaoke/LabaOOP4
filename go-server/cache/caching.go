package cache

import (
	"sync"
)

type Cache struct {
	data map[string]interface{}
	mtx  sync.Mutex
}

func (c *Cache) Get(key string) (interface{}, bool) {
	c.mtx.Lock()

	defer c.mtx.Unlock()

	val, found := c.data[key]
	return val, found
}
