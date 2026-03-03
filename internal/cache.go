package internal

import (
	"fmt"
	"time"
	"errors"
	"sync"
)

type Cache struct {
	entries map[string]cacheEntry
	duration time.Duration
	mu sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

func(c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	entry := cacheEntry{
		val : val,
		createdAt : time.Now(),
	}

	c.entries[key] = entry
	c.mu.Unlock()
}

func(c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if entry, ok := c.entries[key]; ok != false {
		return entry.val, true
	}
	return nil, false
}

func(c *Cache) remove(key string) {
	delete(c.entries, key)
}

func NewCache(duration int) (*Cache, error) {
	
	tick, _ := time.ParseDuration("10s")
	d, err := time.ParseDuration(fmt.Sprintf("%ds", duration))
	
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Could not parse duration: %v", err))
	}
	
	cache := Cache{
		entries : make(map[string]cacheEntry),
		duration : d,
	}

	ticker := time.NewTicker(tick)
	go cache.validateCache(ticker.C)
	

	return &cache, nil
}

func(c *Cache) validateCache(ch <-chan(time.Time)) {

	for {
		select {
			case <-ch:
			c.mu.Lock()
			for key, entry := range c.entries {
				if time.Since(entry.createdAt) > c.duration {
					c.remove(key)
				}
			}
			c.mu.Unlock()
		}
	}

}

