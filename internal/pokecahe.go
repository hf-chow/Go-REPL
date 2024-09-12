package pokecache

import (
	"log"
	"sync"
	"time"
)

const INTERVAL = 5 

type Cache struct {
	Entry		map[string]cacheEntry 	
	Lock		sync.Mutex
}

type cacheEntry struct {
	createdTime	time.Time
	value		[]byte
}

func NewCache(interval time.Duration) Cache {
	newCache :=  Cache{
		Entry: 		make(map[string]cacheEntry),
	}
	newCache.reapLoop(interval)
	return newCache
}

func (cache *Cache) Add(key string, val []byte) {
	var entry cacheEntry
	entry.value = val 
	entry.createdTime = time.Now()

	cache.Entry[key] = entry
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.Lock.Lock()
	defer cache.Lock.Unlock()

	entry, found := cache.Entry[key]
	if found {
		return entry.value, true
	} else {
		log.Printf("no such cache found: %v", key)
		return []byte{}, false
	}
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	done := make(chan bool)

	go func() {
		time.Sleep(INTERVAL * time.Second)
		for key, entry := range cache.Entry{
			cacheDuration := time.Since(entry.createdTime)
			if INTERVAL < (cacheDuration) {
				delete(cache.Entry, key)
				log.Printf("cache etntry %v deleted", key)
			}
		}
		done <- true
	}()

}
