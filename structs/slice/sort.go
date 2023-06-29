package slice

import (
	"github.com/trsteel/box/functions"
	"github.com/trsteel/box/structs/combination"
)

// Sorter implements sort.Interface.
type Sorter[T any] struct {
	*combination.Pair[Slice[T], functions.CompareFunc[T]]
}

// NewSorter returns a struct pointer.
func NewSorter[T any](slice Slice[T], less functions.CompareFunc[T]) *Sorter[T] {
	return &Sorter[T]{Pair: combination.NewPair(slice, less)}
}

// Len is the number of elements in the slice.
func (s *Sorter[T]) Len() int {
	return s.First.Len()
}

// Swap swaps the elements with indexes i and j.
func (s *Sorter[T]) Swap(i, j int) {
	s.First.Swap(i, j)
}

// Less reports whether the element with index i must sort before the element with index j.
func (s *Sorter[T]) Less(i, j int) bool {
	return s.Second(s.First[i], s.First[j])
}
