package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	mu       sync.Mutex
}

type item struct {
	k Key
	v any
}

func (cache *lruCache) Set(k Key, v any) bool {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	foundEl, ok := cache.items[k]

	if ok {
		foundEl.Value.(*item).v = v
		cache.queue.MoveToFront(foundEl)
		return true
	}

	newItem := &item{k: k, v: v}
	itemEl := cache.queue.PushFront(newItem)
	cache.items[k] = itemEl
	if len(cache.items) > cache.capacity {
		lastEl := cache.queue.Back()
		if lastEl != nil {
			cache.queue.Remove(lastEl)
			delete(cache.items, lastEl.Value.(*item).k)
		}
	}
	return false
}

func (cache *lruCache) Get(k Key) (any, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	foundEl, ok := cache.items[k]
	if ok {
		cache.queue.MoveToFront(foundEl)
		return foundEl.Value.(*item).v, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}

func NewCache(capacity int) Cache {
	if capacity <= 0 {
		panic("capacity must be positive")
	}

	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
