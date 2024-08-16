package cache

import (
	"time"
)

type CacheItem struct {
	Key        string
	Value      string
	Expiration time.Time
}

type LRUCache struct {
	Capacity int
	Items    []CacheItem
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		Capacity: capacity,
		Items:    make([]CacheItem, 0, capacity),
	}
}

func (cache *LRUCache) Set(key string, value string, duration time.Duration) {
	expiration := time.Now().Add(duration)
	item := CacheItem{Key: key, Value: value, Expiration: expiration}

	// Item Allready exist check added
	for i, existingItem := range cache.Items {
		if existingItem.Key == key {
			cache.Items[i] = item
			return
		}
	}

	if len(cache.Items) < cache.Capacity {
		cache.Items = append(cache.Items, item)
	} else {
		cache.Items = append(cache.Items[1:], item)
	}
}

func (cache *LRUCache) Get(key string) (string, bool) {
	for i, item := range cache.Items {
		if item.Key == key {
			//Here using time.After
			if time.Now().After(item.Expiration) {
				cache.Items = append(cache.Items[:i], cache.Items[i+1:]...)
				return "", false
			}
			cache.Items = append(cache.Items[:i], cache.Items[i+1:]...)
			cache.Items = append(cache.Items, item)
			return item.Value, true
		}
	}
	return "", false
}
