package funcs

// Must is a helper function to panic if an error occurs.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// MustTrue is a helper function to panic if ok is false.
func MustTrue[T any](value T, ok bool) T {
	if !ok {
		panic("should be true")
	}
	return value
}
