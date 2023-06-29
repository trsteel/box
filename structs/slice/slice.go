package slice

import "sort"

// Slice is an alise of []T
type Slice[T any] []T

// New returns a slice built by elements
func New[T any](elements ...T) Slice[T] {
	return elements
}

// Clone returns a deepcopy object.
func (s Slice[T]) Clone() Slice[T] {
	dst := make([]T, len(s))
	copy(dst, s)
	return dst
}

// Len is the number of elements in the slice.
func (s Slice[T]) Len() int {
	return len(s)
}

// Swap swaps the elements with indexes i and j.
func (s Slice[T]) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Reverse the element order of slice.
func (s Slice[T]) Reverse() {
	n := s.Len()
	for i := 0; i < n/2; i++ {
		s.Swap(i, n-i-1)
	}
}

// SortBy sorts the slice by less functions.
func (s Slice[T]) SortBy(less func(a, b T) bool) {
	sort.Sort(NewSorter(s, less))
}
