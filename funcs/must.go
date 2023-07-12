package funcs

// Must is a helper function to panic if an error occurs.
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}
