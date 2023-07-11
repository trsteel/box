package funcs

// Zero returns an zero value.
func Zero[T any]() T {
	var zero T
	return zero
}

// IsZero returns value is zero.
func IsZero[T comparable](value T) bool {
	return value == Zero[T]()
}
