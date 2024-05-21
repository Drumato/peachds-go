package peachds

import (
	"context"
	"sync"
)

type ConcurrentOrderedMap[K comparable, V any] struct {
	Keys []K
	Map  map[K]V
	mu   sync.RWMutex
}

func NewConcurrentOrderedMap[K comparable, V any]() *ConcurrentOrderedMap[K, V] {
	return &ConcurrentOrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}
}

func ConcurrentOrderedMapFromMap[K comparable, V any](keys []K, m map[K]V) *ConcurrentOrderedMap[K, V] {
	om := &ConcurrentOrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}

	for _, k := range keys {
		om.Set(k, m[k])
	}

	return om
}

func (o *ConcurrentOrderedMap[K, V]) Length() int {
	defer o.mu.RUnlock()

	o.mu.RLock()
	return len(o.Map)
}

func (o *ConcurrentOrderedMap[K, V]) Get(key K) (V, bool) {
	defer o.mu.RUnlock()

	o.mu.RLock()
	v, ok := o.Map[key]
	return v, ok
}

func (o *ConcurrentOrderedMap[K, V]) Set(key K, value V) {
	defer o.mu.Unlock()

	o.mu.Lock()
	o.Keys = append(o.Keys, key)
	o.Map[key] = value
}

func (o *ConcurrentOrderedMap[K, V]) Iter(
	ctx context.Context,
	ch chan<- OrderedMapIterEntry[K, V],
) {
	defer o.mu.RUnlock()

	o.mu.RLock()
	defer close(ch)

	for _, k := range o.Keys {
		select {
		case <-ctx.Done():
			break
		default:
		}
		key := k

		ch <- OrderedMapIterEntry[K, V]{
			Key:   key,
			Value: o.Map[key],
		}
	}

}
