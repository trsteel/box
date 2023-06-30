package basic

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Ternary is a 1 line if/else statement.
func Ternary[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// Keys returns all keys of map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}

// SortedKeys returns all keys of map in order.
func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	slices.Sort(keys)
	return keys
}

// ContainsKey returns whether key is present in map.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsAnyKey returns whether any key is present in map.
func ContainsAnyKey[K comparable, V any](m map[K]V, keys []K) bool {
	return !ForEach(keys, func(key K) bool { return !ContainsKey(m, key) })
}

// ContainsAllKeys returns whether all keys are present in map.
func ContainsAllKeys[K comparable, V any](m map[K]V, keys []K) bool {
	return ForEach(keys, func(key K) bool { return ContainsKey(m, key) })
}
