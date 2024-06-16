package lrucache

import (
	"cmp"
	"container/list"
)

type FastSingleThreadLRUCache[K cmp.Ordered, V any] struct {
	capacity int
	// the element has the value that is typed as V
	cache map[K]*list.Element
	// the element of the list has the value that is typed as CacheEntry[K, V]
	dll *list.List
}

type CacheEntry[K cmp.Ordered, V any] struct {
	Key   K
	Value V
}

func NewFastSingleThreadLRUCache[K cmp.Ordered, V any](capacity int) LRUCache[K, V] {
	dll := list.New()

	return &FastSingleThreadLRUCache[K, V]{
		capacity: capacity,
		cache:    make(map[K]*list.Element),
		dll:      dll,
	}
}

func (c *FastSingleThreadLRUCache[K, V]) Capacity() int {
	return c.capacity
}

func (c *FastSingleThreadLRUCache[K, V]) Get(key K) (V, bool) {
	element, ok := c.cache[key]
	if !ok {
		var v V
		return v, false
	}

	c.dll.MoveToFront(element)
	return element.Value.(CacheEntry[K, V]).Value, true
}

func (c *FastSingleThreadLRUCache[K, V]) Put(key K, value V) {
	element, ok := c.cache[key]
	if ok {
		element.Value = CacheEntry[K, V]{
			Key:   key,
			Value: value,
		}
		c.dll.MoveToFront(element)
		c.cache[key] = element
		return
	}

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
}
