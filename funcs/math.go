package funcs

import (
	"golang.org/x/exp/constraints"

	basicfuncs "github.com/trsteel/box/funcs/basic"
)

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
