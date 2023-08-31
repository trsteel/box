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

// ReflectIF returns value after reflect or zero if reflect failed.
func ReflectIF[T any](value any, cond bool) (T, bool) {
	if cond {
		return value.(T), true
	}
	var zero T
	return zero, false
}

// ReflectOR returns value after reflect or the fallback value if reflect failed.
func ReflectOR[T any](value any, fallback T) T {
	if v, ok := value.(T); ok {
		return v
	}
	return fallback
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
