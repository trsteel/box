package funcs

import (
	"github.com/trsteel/box/structs/set"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Contains reports whether elem is present in slice.
func Contains[T comparable](slice []T, elem T) bool {
	return slices.Contains(slice, elem)
}

// ContainsAny reports whether anyone is present in slice.
func ContainsAny[T comparable](slice []T, elems []T) bool {
	return set.NewSet(slice...).ContainsAny(elems)
}

// ContainsAll reports whether all elems are present in slice.
func ContainsAll[T comparable](slice []T, elems []T) bool {
	return set.NewSet(slice...).ContainsAll(elems)
}

// Sort sorts a slice of any ordered type in ascending order.
func Sort[T constraints.Ordered](slice []T) {
	slices.Sort(slice)
}

// SortBy sorts the slice by less functions.
func SortBy[T any](slice []T, less func(a, b T) bool) {
	slices.SortFunc(slice, less)
}

// Reverse the element order of slice.
func Reverse[T any](slice []T) {
	n := len(slice)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
