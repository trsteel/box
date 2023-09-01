package funcs

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

// Contains reports whether elem is present in slice.
func Contains[T comparable, S ~[]T](slice S, elem T) bool {
	return slices.Contains(slice, elem)
}

// ContainsAny reports whether anyone is present in slice.
func ContainsAny[T comparable, S ~[]T](slice S, elems []T) bool {
	return ContainsAnyKey(SliceToMap(slice, func(_ T) any { return nil }), elems)
}

// ContainsAll reports whether all elems are present in slice.
func ContainsAll[T comparable, S ~[]T](slice S, elems []T) bool {
	return ContainsAllKeys(SliceToMap(slice, func(_ T) any { return nil }), elems)
}

// SliceToMap convert slice to map.
func SliceToMap[K comparable, V any, S ~[]K](slice S, f func(K) V) map[K]V {
	result := make(map[K]V, len(slice))
	for i := 0; i < len(slice); i++ {
		result[slice[i]] = f(slice[i])
	}
	return result
}

// IndexFunc convert slice to a function which input is index and output is slice[index].
func IndexFunc[T any, S ~[]T](slice S) func(int) T {
	return func(i int) T { return slice[i] }
}

// Sort sorts a slice of any ordered type in ascending order.
func Sort[T constraints.Ordered, S ~[]T](slice S) {
	slices.Sort(slice)
}

// SortBy sorts the slice by less functions.
func SortBy[T any, S ~[]T](slice S, less func(a, b T) bool) {
	slices.SortFunc(slice, less)
}

// Reverse the element order of slice.
func Reverse[T any, S ~[]T](slice S) {
	n := len(slice)
	for i, j := 0, n-1; i < n/2; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// Pack returns a slice built by elems.
func Pack[T any](elems ...T) []T {
	return elems
}

// Fill initial to each element of slice.
func Fill[T any, S ~[]T](slice S, initial T) []T {
	return FillBy(slice, WithArg[int, T](CopyFunc(initial)))
}

// FillBy f to each element of slice.
func FillBy[T any, S ~[]T](slice S, f func(index int) T) []T {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(i)
	}
	return slice
}

// Repeat returns a slice that each element equals to initial
func Repeat[T any](len int, initial T) []T {
	return Fill(make([]T, len), initial)
}

// RepeatBy returns a slice that each element generated by f.
func RepeatBy[T any](len int, f func(index int) T) []T {
	return FillBy(make([]T, len), f)
}

// Map manipulates a slice and transforms it to a slice of another type.
func Map[T any, R any, S ~[]T](slice S, iteratee func(item T) R) []R {
	return FillBy(make([]R, len(slice)), Combine(IndexFunc(slice), iteratee))
}

// MapPtr manipulates a slice and transforms each element's pointer of it to a slice of another type.
func MapPtr[T any, R any, S ~[]T](slice S, iteratee func(item *T) R) []R {
	return FillBy(make([]R, len(slice)), Combine(Combine(IndexFunc(slice), ToPtr[T]), iteratee))
}
