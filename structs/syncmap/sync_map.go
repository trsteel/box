package syncmap

import "sync"

// SyncMap is a wrap of sync.Map with certain K/V type.
type SyncMap[K, V any] struct {
	sync.Map
}

// Load is a wrap of sync.Map's Load func with certain K/V type.
func (s *SyncMap[K, V]) Load(key K) (value V, ok bool) {
	if _value, _ok := s.Map.Load(key); _ok {
		return _value.(V), _ok
	}
	var zero V
	return zero, false
}

// Store is a wrap of sync.Map's Store func with certain K/V type.
func (s *SyncMap[K, V]) Store(key K, value V) {
	s.Map.Store(key, value)
}

// LoadOrStore is a wrap of sync.Map's LoadOrStore func with certain K/V type.
func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (actual V, loaded bool) {
	_actual, _loaded := s.Map.LoadOrStore(key, value)
	return _actual.(V), _loaded
}

// LoadAndDelete is a wrap of sync.Map's LoadAndDelete func with certain K/V type.
func (s *SyncMap[K, V]) LoadAndDelete(key K) (value V, loaded bool) {
	if _value, _loaded := s.Map.LoadAndDelete(key); _loaded {
		return _value.(V), _loaded
	}
	var zero V
	return zero, false
}

// Delete is a wrap of sync.Map's Delete func with certain K/V type.
func (s *SyncMap[K, V]) Delete(key K) {
	s.Map.Delete(key)
}

// Swap is a wrap of sync.Map's Swap func with certain K/V type.
func (s *SyncMap[K, V]) Swap(key K, value V) (previous V, loaded bool) {
	if _previous, _loaded := s.Map.Swap(key, value); _loaded {
		return _previous.(V), _loaded
	}
	var zero V
	return zero, false
}

// CompareAndSwap is a wrap of sync.Map's CompareAndSwap func with certain K/V type.
func (s *SyncMap[K, V]) CompareAndSwap(key K, old, new V) (swapped bool) {
	return s.Map.CompareAndSwap(key, old, new)
}

// CompareAndDelete is a wrap of sync.Map's CompareAndDelete func with certain K/V type.
func (s *SyncMap[K, V]) CompareAndDelete(key K, old V) (deleted bool) {
	return s.Map.CompareAndDelete(key, old)
}

// Range is a wrap of sync.Map's Range func with certain K/V type.
func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.Map.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}
