package funcs

import (
	"golang.org/x/exp/constraints"

	basicfuncs "github.com/trsteel/box/funcs/basic"
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
	return basicfuncs.Reduce(slice, origin, func(agg T, item T) T {
		return basicfuncs.Ternary(agg >= item, agg, item)
	})
}

// Min returns the minimum value between origin and each element of slice
func Min[T constraints.Ordered](origin T, slice ...T) T {
	return basicfuncs.Reduce(slice, origin, func(agg T, item T) T {
		return basicfuncs.Ternary(agg <= item, agg, item)
	})
}
