package syncmap

import (
	"sync"

	"github.com/trsteel/box/funcs"
)

// SyncMap is a wrap of sync.Map with certain K/V type.
type SyncMap[K, V any] struct {
	sync.Map
}

// Load is a wrap of sync.Map's Load func with certain K/V type.
func (s *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	return funcs.ReflectIF[V](s.Map.Load(key))
}

// Store is a wrap of sync.Map's Store func with certain K/V type.
func (s *SyncMap[K, V]) Store(key K, value V) {
	s.Map.Store(key, value)
}

// LoadOrStore is a wrap of sync.Map's LoadOrStore func with certain K/V type.
func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	return funcs.ReflectIF[V](s.Map.LoadOrStore(key, value))
}

// LoadAndDelete is a wrap of sync.Map's LoadAndDelete func with certain K/V type.
func (s *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	return funcs.ReflectIF[V](s.Map.LoadAndDelete(key))
}

// Delete is a wrap of sync.Map's Delete func with certain K/V type.
func (s *SyncMap[K, V]) Delete(key K) {
	s.Map.Delete(key)
}

// Range is a wrap of sync.Map's Range func with certain K/V type.
func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
