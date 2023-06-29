package functions

// CompareFunc returns if the two input params match some conditions.
type CompareFunc[T any] func(T, T) bool
