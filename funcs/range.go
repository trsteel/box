package funcs

import (
	"golang.org/x/exp/constraints"
)

// ForRange apply f to each number in a sequence construct with step in range [start, end).
func ForRange[T constraints.Integer | constraints.Float](start, end, step T, f func(T) bool) bool {
	for i := start; i < end; i += step {
		if !f(i) {
			return false
		}
	}
	return true
}

// ForEach apply f to each element of slice until f returns false.
func ForEach[T any](slice []T, f func(T) bool) bool {
	return ForRange(0, len(slice), 1, Combine(IndexFunc(slice), f))
}

// ForEachPtr apply f to each element of slice until f returns false.
func ForEachPtr[T any](slice []T, f func(*T) bool) bool {
	return ForRange(0, len(slice), 1, Combine(Combine(IndexFunc(slice), ToPtr[T]), f))
}

// Reduce reduces slice to an accumulated result of running each element in slice by accumulator.
func Reduce[T any, R any](slice []T, initial R, accumulator func(agg R, item T) R) R {
	for i := 0; i < len(slice); i++ {
		initial = accumulator(initial, slice[i])
	}
	return initial
}
