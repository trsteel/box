package funcs

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// Keys returns all keys of map.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	return maps.Keys(m)
}

// Values returns all values of map.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	return maps.Values(m)
}

// SortedKeys returns all keys of map in order.
func SortedKeys[M ~map[K]V, K constraints.Ordered, V any](m M) []K {
	keys := maps.Keys(m)
	slices.Sort(keys)
	return keys
}

// ContainsKey returns whether key is present in map.
func ContainsKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsAnyKey returns whether any key is present in map.
func ContainsAnyKey[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	return !ForEach(keys, func(key K) bool { return !ContainsKey(m, key) })
}

// ContainsAllKeys returns whether all keys are present in map.
func ContainsAllKeys[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	return ForEach(keys, func(key K) bool { return ContainsKey(m, key) })
}

// ForEachKV apply f to each KV pair of map until f returns false.
func ForEachKV[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Clone returns a copy of m.
func Clone[M ~map[K]V, K comparable, V any](m M) M {
	return maps.Clone(m)
}

// Merge merges multiple maps from left to right.
func Merge[M ~map[K]V, K comparable, V any](maps ...M) M {
	out := M{}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}
