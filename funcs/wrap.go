package funcs

// Combine two functions as one.
func Combine[T, R, S any](first func(T) R, second func(R) S) func(T) S {
	return func(t T) S { return second(first(t)) }
}
