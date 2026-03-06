package hw04lrucache

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
}

type item struct {
	k Key
	v any
}

func (cache *lruCache) Set(k Key, v any) bool {
	foundEl, ok := cache.items[k]

	newItem := item{k: k, v: v}
	if !ok {
		itemEl := cache.queue.PushFront(newItem)
		cache.items[k] = itemEl
		if len(cache.items) > cache.capacity {
			lastEl := cache.queue.Back()
			cache.queue.Remove(lastEl)
			delete(cache.items, lastEl.Value.(item).k)
		}
		return false
	}
	foundEl.Value = newItem
	cache.queue.MoveToFront(foundEl)
	cache.items[k] = foundEl

	return true
}

func (cache *lruCache) Get(k Key) (any, bool) {
	foundEl, ok := cache.items[k]
	if ok {
		cache.queue.MoveToFront(foundEl)
		return foundEl.Value.(item).v, true
	}
	return nil, false
}

func (cache *lruCache) Clear() {
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
