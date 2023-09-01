package funcs

import (
	"golang.org/x/exp/constraints"
)

// Equal returns if a equals to b
func Equal[T comparable](a, b T) bool {
	return a == b
}

// LessThan returns if a is less than b
func LessThan[T constraints.Ordered](a, b T) bool {
	return a < b
}

// LargerThan returns if a is larger than b
func LargerThan[T constraints.Ordered](a, b T) bool {
	return a > b
}

// Between returns if value is in [lowerBound, upperBound].
func Between[T constraints.Ordered](value, lowerBound, upperBound T) bool {
	return lowerBound <= value && value <= upperBound
}

// Max returns the maximum value between origin and each element of slice
func Max[T constraints.Ordered](origin T, slice ...T) T {
	return Reduce(slice, origin, Max2[T])
}

// Max2 returns the maximum value between a and b
func Max2[T constraints.Ordered](a, b T) T {
	return Ternary(a >= b, a, b)
}

// Min returns the minimum value between origin and each element of slice
func Min[T constraints.Ordered](origin T, slice ...T) T {
	return Reduce(slice, origin, Min2[T])
}

// Min2 returns the minimum value between a and b
func Min2[T constraints.Ordered](a, b T) T {
	return Ternary(a <= b, a, b)
}

// Sum returns the summation between origin and each element of slice
func Sum[T constraints.Ordered](origin T, slice ...T) T {
	return Reduce(slice, origin, Sum2[T])
}

// Sum2 returns the summation between a and b
func Sum2[T constraints.Ordered](a, b T) T {
	return a + b
}
