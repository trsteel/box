package funcs

import (
	basicfuncs "github.com/trsteel/box/funcs/basic"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
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
	return basicfuncs.SortedKeys(m)
}

// ContainsKey returns whether key is present in map.
func ContainsKey[M ~map[K]V, K comparable, V any](m M, key K) bool {
	return basicfuncs.ContainsKey(m, key)
}

// ContainsAnyKey returns whether any key is present in map.
func ContainsAnyKey[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	return basicfuncs.ContainsAnyKey(m, keys)
}

// ContainsAllKeys returns whether all keys are present in map.
func ContainsAllKeys[M ~map[K]V, K comparable, V any](m M, keys []K) bool {
	return basicfuncs.ContainsAllKeys(m, keys)
}

// ForEachKV apply f to each KV pair of map until f returns false.
func ForEachKV[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) bool {
	return basicfuncs.ForEachKV(m, f)
}

// Clone returns a copy of m.
func Clone[K comparable, V any](m map[K]V) map[K]V {
	return maps.Clone(m)
}

// Merge merges multiple maps from left to right.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	out := map[K]V{}
	for _, m := range maps {
		for k, v := range m {
			out[k] = v
		}
	}
	return out
}
