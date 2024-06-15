package orderedmap

import (
	"cmp"
	"context"
)

type OrderedMap[K cmp.Ordered, V any] interface {
	Get(key K) (V, bool)
	Set(key K, value V)
	Length() int
	Delete(key K) (V, bool)
	Iter(ctx context.Context, ch chan<- OrderedMapIterEntry[K, V])
}

type OrderedMapIterEntry[K comparable, V any] struct {
	Key   K
	Value V
}
