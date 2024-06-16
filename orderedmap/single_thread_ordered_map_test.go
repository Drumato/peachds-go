package orderedmap_test

import (
	"context"
	"testing"

	"github.com/Drumato/peachds-go/orderedmap"
	"github.com/stretchr/testify/assert"
)

func TestSingleThreadOrderedMap_Set(t *testing.T) {
	m := &orderedmap.SingleThreadOrderedMap[string, int]{
		Keys: make([]string, 0),
		Map:  make(map[string]int),
	}
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	assert.Equal(t, 3, len(m.Map))
}

func TestSingleThreadOrderedMap_Length(t *testing.T) {
	m := orderedmap.SingleThreadOrderedMapFromMap([]string{"one", "two", "three"}, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert.Equal(t, 3, m.Length())
}

func TestSingleThreadOrderedMap_Iter(t *testing.T) {
	m := orderedmap.SingleThreadOrderedMapFromMap([]string{"one", "two", "three"}, map[string]int{
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

func TestSingleThreadOrderedMap_Delete(t *testing.T) {
	m := orderedmap.SingleThreadOrderedMapFromMap([]string{"one", "two", "three"}, map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	m.Delete("two")
	assert.Equal(t, 2, m.Length())
	v, ok := m.Get("two")
	assert.False(t, ok)
	assert.Equal(t, 0, v)
}

func BenchmarkSingleThreadOrderedMap_Set(b *testing.B) {
	m := orderedmap.NewSingleThreadOrderedMap[int, int]()

	b.ResetTimer()
	for i := range b.N {
		m.Set(i, i)
	}
}

func BenchmarkSingleThreadOrderedMap_Get(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewSingleThreadOrderedMap[int, int]()
	for i := range b.N {
		m.Set(i, i)
	}

	b.StartTimer()

	for i := range b.N {
		m.Get(i)
	}
}

func BenchmarkSingleThreadOrderedMap_Iter(b *testing.B) {
	b.StopTimer()
	m := orderedmap.NewSingleThreadOrderedMap[int, int]()
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
