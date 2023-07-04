package basic

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

// Ternary is a 1 line if/else statement.
func Ternary[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
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
