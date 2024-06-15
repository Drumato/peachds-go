package orderedmap

import (
	"cmp"
	"context"
)

type SerialOrderedMap[K cmp.Ordered, V any] struct {
	Keys []K
	Map  map[K]V
}

func NewSerialOrderedMap[K cmp.Ordered, V any]() OrderedMap[K, V] {
	return &SerialOrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}
}

func SerialOrderedMapFromMap[K cmp.Ordered, V any](keys []K, m map[K]V) *SerialOrderedMap[K, V] {
	om := &SerialOrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}

	for _, k := range keys {
		om.Set(k, m[k])
	}

	return om
}

func (o *SerialOrderedMap[K, V]) Length() int {
	return len(o.Map)
}

func (o *SerialOrderedMap[K, V]) Get(key K) (V, bool) {
	v, ok := o.Map[key]
	return v, ok
}

func (o *SerialOrderedMap[K, V]) Set(key K, value V) {
	o.Keys = append(o.Keys, key)
	o.Map[key] = value
}

func (o *SerialOrderedMap[K, V]) Delete(key K) (V, bool) {
	foundIdx := -1
	for i, k := range o.Keys {
		if k == key {
			foundIdx = i
			break
		}
	}

	var deleted V
	if foundIdx == -1 {
		return deleted, false
	}

	o.Keys = append(o.Keys[:foundIdx], o.Keys[foundIdx+1:]...)
	deleted = o.Map[key]
	delete(o.Map, key)
	return deleted, true
}

func (o *SerialOrderedMap[K, V]) Iter(
	ctx context.Context,
	ch chan<- OrderedMapIterEntry[K, V],
) {
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
