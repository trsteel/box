package funcs

import (
	basicfuncs "github.com/trsteel/box/funcs/basic"
	"golang.org/x/exp/constraints"
)

// Keys returns all keys of map.
func Keys[K comparable, V any](m map[K]V) []K {
	return basicfuncs.Keys(m)
}

// SortedKeys returns all keys of map in order.
func SortedKeys[K constraints.Ordered, V any](m map[K]V) []K {
	return basicfuncs.SortedKeys(m)
}

// ContainsKey returns whether key is present in map.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	return basicfuncs.ContainsKey(m, key)
}

// ContainsAnyKey returns whether any key is present in map.
func ContainsAnyKey[K comparable, V any](m map[K]V, keys []K) bool {
	return basicfuncs.ContainsAnyKey(m, keys)
}

// ContainsAllKeys returns whether all keys are present in map.
func ContainsAllKeys[K comparable, V any](m map[K]V, keys []K) bool {
	return basicfuncs.ContainsAllKeys(m, keys)
}

// ForEachKV apply f to each KV pair of map until f returns false.
func ForEachKV[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	return basicfuncs.ForEachKV(m, f)
}
