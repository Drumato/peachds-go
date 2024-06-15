package orderedmap_test

import (
	"context"
	"testing"

	"github.com/Drumato/peachds-go/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestConcurrentOrderedMap_Set(t *testing.T) {
	m := orderedmap.NewConcurrentOrderedMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	assert.Equal(t, 3, len(m.Map))
}

func TestConcurrentOrderedMap_Length(t *testing.T) {
	m := orderedmap.ConcurrentOrderedMapFromMap([]string{"one", "two", "three"}, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert.Equal(t, 3, m.Length())
}

func TestConcurrentOrderedMap_Iter(t *testing.T) {
	m := orderedmap.ConcurrentOrderedMapFromMap([]string{"one", "two", "three"}, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	ch := make(chan orderedmap.OrderedMapIterEntry[string, int], 3)
	go m.Iter(context.Background(), ch)
	assert.Equal(t, 1, (<-ch).Value)
	assert.Equal(t, 2, (<-ch).Value)
	assert.Equal(t, 3, (<-ch).Value)
}

func BenchmarkConcurrentOrderedMap_Set(b *testing.B) {
	m := orderedmap.NewConcurrentOrderedMap[int, int]()

	b.ResetTimer()
	for i := range b.N {
		m.Set(i, i)
	}
}

func BenchmarkConcurrentOrderedMap_Get(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewConcurrentOrderedMap[int, int]()
	for i := range b.N {
		m.Set(i, i)
	}

	b.StartTimer()

	for i := range b.N {
		m.Get(i)
	}
}

func BenchmarkConcurrentOrderedMap_Iter(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewConcurrentOrderedMap[int, int]()
	for i := range b.N {
		m.Set(i, i)
	}

	b.StartTimer()

	ch := make(chan orderedmap.OrderedMapIterEntry[int, int], 1024)
	go m.Iter(context.Background(), ch)

	for range b.N {
		<-ch
	}
}
