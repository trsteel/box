package funcs

// Empty returns an empty value.
func Empty[T any]() T {
	var zero T
	return zero
}
