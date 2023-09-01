package set

import "github.com/trsteel/box/funcs"

// Set implements a set.
type Set[T comparable] map[T]any

// NewSet initializes a new Set instance with all elements added.
func NewSet[T comparable](elems ...T) Set[T] {
	s := make(Set[T], len(elems))
	s.Add(elems...)
	return s
}

// Add values into this set.
func (s Set[T]) Add(elems ...T) {
	for _, elem := range elems {
		s[elem] = nil
	}
}

// Len returns number of elements stored in this set.
func (s Set[T]) Len() int {
	return len(s)
}

// Contains checks whether this set contains the element.
func (s Set[T]) Contains(elem T) bool {
	return funcs.ContainsKey(s, elem)
}

// ContainsAny checks whether anyone is present in set.
func (s Set[T]) ContainsAny(elems []T) bool {
	return funcs.ContainsAnyKey(s, elems)
}

// ContainsAll checks whether all elems are present in set.
func (s Set[T]) ContainsAll(elems []T) bool {
	return funcs.ContainsAllKeys(s, elems)
}
