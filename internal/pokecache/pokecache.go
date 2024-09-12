package pokecache

import(
	"time"
	"sync"
)

type cacheEntry struct{
	createdAt time.Time
	val []byte
}

type Cache struct{
	data map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) Cache{
	result := Cache{
		data: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
		interval: interval,
	}
	go result.reapLoop()
	return result
}

func (cache *Cache)Add(key string, value []byte){
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry := cacheEntry{
		createdAt: time.Now(),
		val: value,
	}
	cache.data[key] = entry
}

func (cache *Cache)Get(key string)([]byte, bool){
	cache.mu.Lock()
	defer cache.mu.Unlock()
	entry, ok := cache.data[key]
	if !ok{
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache)reapLoop(){
	ticker := time.NewTicker(cache.interval)
	for range ticker.C {
		currentTime := time.Now()
		cache.mu.Lock()
		for k, v := range cache.data{
			if currentTime.After(v.createdAt.Add(cache.interval)){
				delete(cache.data, k)
			}
		}
		cache.mu.Unlock()
	}
}