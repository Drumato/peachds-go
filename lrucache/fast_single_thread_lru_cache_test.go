package lrucache_test

import (
	"testing"

	"github.com/Drumato/peachds-go/lrucache"
	"github.com/stretchr/testify/assert"
)

func TestFastSingleThreadLRUCache(t *testing.T) {
	cache := lrucache.NewFastSingleThreadLRUCache[int, int](2)

	val, ok := cache.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	// キャッシュに要素を追加
	cache.Put(1, 1)

	val, ok = cache.Get(1)
	assert.True(t, ok)
	assert.Equal(t, 1, val)

	// キャッシュに要素を追加
	cache.Put(2, 2)
	if val, ok := cache.Get(2); !ok || val != 2 {
		t.Fatalf("expected (2, true), got (%d, %v)", val, ok)
	}

	// キャッシュの容量を超えたときに古い要素が削除されることを確認
	cache.Put(3, 3) // ここでキー1が削除される

	val, ok = cache.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)

	val, ok = cache.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 2, val)

	val, ok = cache.Get(3)
	assert.True(t, ok)
	assert.Equal(t, 3, val)

	// 既存のキーの値を更新して、順序が正しいか確認
	cache.Put(2, 20)

	val, ok = cache.Get(2)
	assert.True(t, ok)
	assert.Equal(t, 20, val)
}

func TestFastSingleThreadLRUCache_EdgeCases(t *testing.T) {
	cache := lrucache.NewFastSingleThreadLRUCache[int, int](0)

	// 容量0のキャッシュに要素を追加しても何も保持されないことを確認
	cache.Put(1, 1)

	val, ok := cache.Get(1)
	assert.False(t, ok)
	assert.Equal(t, 0, val)
}

func BenchmarkFastSingleThreadLRUCache_Put(b *testing.B) {
	cache := lrucache.NewFastSingleThreadLRUCache[int, int](b.N)
	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}
}

func BenchmarkFastSingleThreadLRUCache_Get(b *testing.B) {
	cache := lrucache.NewFastSingleThreadLRUCache[int, int](b.N)
	for i := 0; i < b.N; i++ {
		cache.Put(i, i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		cache.Get(i)
	}
}
