package peachds

import "context"

type OrderedMap[K comparable, V any] struct {
	Keys []K
	Map  map[K]V
}

type OrderedMapIterEntry[K comparable, V any] struct {
	Key   K
	Value V
}

func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}
}

func OrderedMapFromMap[K comparable, V any](m map[K]V) *OrderedMap[K, V] {
	om := &OrderedMap[K, V]{
		Keys: make([]K, 0),
		Map:  make(map[K]V),
	}

	for k, v := range m {
		om.Set(k, v)
	}

	return om
}

func (o *OrderedMap[K, V]) Length() int {
	return len(o.Map)
}

func (o *OrderedMap[K, V]) Get(key K) (V, bool) {
	v, ok := o.Map[key]
	return v, ok
}

func (o *OrderedMap[K, V]) Set(key K, value V) {
	o.Keys = append(o.Keys, key)
	o.Map[key] = value
}

func (o *OrderedMap[K, V]) Iter(
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
