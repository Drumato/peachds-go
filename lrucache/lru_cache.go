package lrucache

import "cmp"

type LRUCache[K cmp.Ordered, V any] interface {
	Capacity() int
	Get(key K) (V, bool)
	Put(key K, value V)
}
