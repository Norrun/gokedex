package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	eitries map[string]cacheEntry
	mux     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		make(map[string]cacheEntry),
		&sync.Mutex{},
	}

	c.reapLoop(interval)
	return c
}

func (c Cache) Add(key string, val []byte) {
	c.mux.Lock()
	c.eitries[key] = cacheEntry{time.Now(), val}
	c.mux.Unlock()
}
func (c Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	v, exists := c.eitries[key]
	c.mux.Unlock()

	return v.val, exists
}

func (c Cache) reapLoop(interval time.Duration) {
	clock := time.NewTicker(interval)
	go func() {
		for t := range clock.C {
			c.mux.Lock()
			for k, v := range c.eitries {
				if v.createdAt.Add(interval).Before(t) {
					delete(c.eitries, k)
				}
			}
			c.mux.Unlock()
		}
	}()
}
