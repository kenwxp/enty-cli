package cache

import (
	"entysquare/filer-backend/util"
	"sync"
	"time"
)

type ICache interface {
	//size support: 1KB,100KB,1MB,2MB,1GB
	SetMaxMemory(size string) bool
	//key expire after expire time
	Set(key string, val interface{}, expire time.Duration)
	//get one key
	Get(key string) (interface{}, bool)
	//delete one key
	Del(key string) bool
	//exists one key
	Exists(key string) bool
	//delete all key
	Flush() bool
	//get all key
	Keys() []string
	//garbage collection every second
	GcLoop()
	//delete expired key
	DeleteExpired()
}

var CacheAll = NewCache()

type Item struct {
	Object     interface{}
	Expiration int64
}

type cache struct {
	size     string
	items    map[string]Item
	mu       sync.RWMutex
	interval time.Duration
}

func (c *cache) SetMaxMemory(size string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	c.size = size
	return true
}

func (c *cache) Set(k string, x interface{}, d time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	e := util.TimeNow().Add(d * time.Second).Unix()

	c.items[k] = Item{
		Object:     x,
		Expiration: e,
	}
}

func (c *cache) Get(k string) (interface{}, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	item, found := c.items[k]
	if !found {
		return nil, false
	}

	if util.TimeNow().Unix() > item.Expiration {
		return nil, false
	}

	return item.Object, true
}

func (c *cache) Del(k string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, found := c.items[k]; found {
		delete(c.items, k)
		return true
	}

	return false
}

func (c *cache) Exists(k string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	if _, found := c.items[k]; found {
		return true
	}

	return false
}

func (c *cache) Flush() bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items = map[string]Item{}
	return true
}

func (c *cache) Keys() []string {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var keys []string
	for k := range c.items {
		keys = append(keys, k)
	}

	return keys
}

func (c *cache) GcLoop() {
	ticker := time.NewTicker(c.interval)

	for {
		select {
		case <-ticker.C:
			c.DeleteExpired()
		}
	}
}

func (c *cache) DeleteExpired() {
	now := util.TimeNow().Unix()

	for k, v := range c.items {
		if now > v.Expiration {
			c.Del(k)
		}
	}
}

func NewCache() ICache {
	c := &cache{
		size:     "1024",
		items:    map[string]Item{},
		interval: time.Second,
	}

	//go c.GcLoop()
	return c
}

//func TestCacheExample(t *testing.T) {
//	c := NewCache()
//
//	c.Set("foo", "bar", 5)
//
//	time.Sleep(10 * time.Second)
//	r, b := c.Get("foo")
//
//	t.Log(r, b)
//	t.Log("test cache finished.")
//
//}
