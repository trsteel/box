package sortedmap

import (
	"golang.org/x/exp/constraints"

	basicfuncs "github.com/trsteel/box/funcs/basic"
)

// SortedMap is a map with key's type is constraints.Ordered.
type SortedMap[K constraints.Ordered, V any] map[K]V

// Range apply f to each kv pair in map by the order of key.
func (s SortedMap[K, V]) Range(f func(K, V) bool) {
	basicfuncs.ForEach(basicfuncs.SortedKeys(s), func(key K) bool { return f(key, s[key]) })
}
