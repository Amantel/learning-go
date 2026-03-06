package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (cache *lruCache) Set(k Key, v any) bool {
	foundEl, ok := cache.items[k]
	if !ok {
		itemEl := cache.queue.PushFront(v)
		cache.items[k] = itemEl
		if len(cache.items) > cache.capacity {
			cache.queue.Remove(cache.queue.Back())
		}
		return false
	}
	// cache.queue.Remove(foundEl)
	// itemEl := cache.queue.PushFront(v)
	// cache.items[k] = itemEl

	foundEl.Value = v
	cache.queue.MoveToFront(foundEl)
	cache.items[k] = foundEl

	return true
}

func (cache *lruCache) Get(k Key) (any, bool) {
	foundEl, ok := cache.items[k]
	if ok {
		cache.queue.MoveToFront(foundEl)
		return foundEl.Value, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
	// for _, key := range cache.items {
	// 	cache.queue.Remove(key)
	// }
	cache.items = make(map[Key]*ListItem, cache.capacity)
	cache.queue = NewList()
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
