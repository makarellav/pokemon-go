package cache

import (
	"sync"
	"time"
)

type RequestCache struct {
	entries map[string]cacheEntry
	mu      sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	data      []byte
}

func NewRequestCache(cacheTTL time.Duration) *RequestCache {
	newCache := &RequestCache{
		entries: make(map[string]cacheEntry),
		mu:      sync.Mutex{},
	}

	go newCache.reap(cacheTTL)

	return newCache
}

func (rq *RequestCache) Add(key string, data []byte) {
	rq.mu.Lock()
	defer rq.mu.Unlock()

	rq.entries[key] = cacheEntry{
		createdAt: time.Now(),
		data:      data,
	}
}

func (rq *RequestCache) Get(key string) ([]byte, bool) {
	rq.mu.Lock()
	defer rq.mu.Unlock()

	entry, ok := rq.entries[key]

	return entry.data, ok
}

func (rq *RequestCache) reap(cacheTTL time.Duration) {
	ticker := time.NewTicker(cacheTTL)

	for range ticker.C {
		rq.reapLoop(cacheTTL)
	}
}

func (rq *RequestCache) reapLoop(cacheTTL time.Duration) {
	rq.mu.Lock()
	defer rq.mu.Unlock()

	for k, v := range rq.entries {
		if time.Since(v.createdAt) >= cacheTTL {
			delete(rq.entries, k)
		}
	}
}
