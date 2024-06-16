package lrucache

import (
	"cmp"
	"container/list"
	"sync"
)

type FastConcurrentLRUCache[K cmp.Ordered, V any] struct {
	capacity int
	// the element has the value that is typed as V
	cache map[K]*list.Element
	// the element of the list has the value that is typed as CacheEntry[K, V]
	dll *list.List
	mu  *sync.RWMutex
}

func NewFastConcurrentLRUCache[K cmp.Ordered, V any](capacity int) LRUCache[K, V] {
	dll := list.New()

	return &FastConcurrentLRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*list.Element),
		dll:      dll,
		mu:       new(sync.RWMutex),
	}
}

func (c *FastConcurrentLRUCache[K, V]) Capacity() int {
	defer c.mu.RUnlock()
	c.mu.RLock()
	return c.capacity
}

func (c *FastConcurrentLRUCache[K, V]) Get(key K) (V, bool) {
	c.mu.RLock()
	element, ok := c.cache[key]
	c.mu.RUnlock()
	if !ok {
		var v V
		return v, false
	}

	c.mu.Lock()
	c.dll.MoveToFront(element)
	c.mu.Unlock()

	return element.Value.(CacheEntry[K, V]).Value, true
}

func (c *FastConcurrentLRUCache[K, V]) Put(key K, value V) {
	c.mu.RLock()
	element, ok := c.cache[key]
	c.mu.RUnlock()
	if ok {
		element.Value = CacheEntry[K, V]{
			Key:   key,
			Value: value,
		}
		c.mu.Lock()
		c.dll.MoveToFront(element)
		c.cache[key] = element
		c.mu.Unlock()
		return
	}

	c.mu.Lock()
	newE := c.dll.PushFront(CacheEntry[K, V]{
		Key:   key,
		Value: value,
	})
	c.cache[key] = newE
	if len(c.cache) > c.capacity {
		oldest := c.dll.Back()
		c.dll.Remove(oldest)
		delete(c.cache, oldest.Value.(CacheEntry[K, V]).Key)
	}
	c.mu.Unlock()
}
