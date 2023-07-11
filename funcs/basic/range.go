package basic

import "golang.org/x/exp/constraints"

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
	return ForRange(0, len(slice), 1, func(i int) bool { return f(slice[i]) })
}

// ForEachPtr apply f to each element of slice until f returns false.
func ForEachPtr[T any](slice []T, f func(*T) bool) bool {
	return ForRange(0, len(slice), 1, func(i int) bool { return f(&slice[i]) })
}

// ForEachKV apply f to each KV pair of map until f returns false.
func ForEachKV[M ~map[K]V, K comparable, V any](m M, f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}

// Reduce reduces slice to an accumulated result of running each element in slice by accumulator.
func Reduce[T any, R any](slice []T, initial R, accumulator func(agg R, item T) R) R {
	ForRange(0, len(slice), 1, func(i int) bool { initial = accumulator(initial, slice[i]); return true })
	return initial
}
