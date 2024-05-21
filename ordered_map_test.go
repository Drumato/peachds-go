package peachds_test

import (
	"context"
	"testing"

	"github.com/Drumato/peachds-go"
	"github.com/stretchr/testify/assert"
)

func TestOrderedMap_Set(t *testing.T) {
	m := peachds.NewOrderedMap[string, int]()
	m.Set("one", 1)
	m.Set("two", 2)
	m.Set("three", 3)

	assert.Equal(t, 3, len(m.Map))
}

func TestOrderedMap_Length(t *testing.T) {
	m := peachds.OrderedMapFromMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	assert.Equal(t, 3, m.Length())
}

func TestOrderedMap_Iter(t *testing.T) {
	m := peachds.OrderedMapFromMap(map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	})

	ch := make(chan peachds.OrderedMapIterEntry[string, int], 3)
	go m.Iter(context.Background(), ch)
	assert.Equal(t, 1, (<-ch).Value)
	assert.Equal(t, 2, (<-ch).Value)
	assert.Equal(t, 3, (<-ch).Value)
}

func BenchmarkOrderedMap_Set(b *testing.B) {
	m := peachds.NewOrderedMap[int, int]()

	b.ResetTimer()
	for i := range b.N {
		m.Set(i, i)
	}
}

func BenchmarkOrderedMap_Get(b *testing.B) {
	b.StopTimer()
	m := peachds.NewOrderedMap[int, int]()
	for i := range b.N {
		m.Set(i, i)
	}

	b.StartTimer()

	for i := range b.N {
		m.Get(i)
	}
}

func BenchmarkOrderedMap_Iter(b *testing.B) {
	b.StopTimer()
	m := peachds.NewOrderedMap[int, int]()
	for i := range b.N {
		m.Set(i, i)
	}

	b.StartTimer()

	ch := make(chan peachds.OrderedMapIterEntry[int, int], 1024)
	go m.Iter(context.Background(), ch)

	for range b.N {
		<-ch
	}
}
