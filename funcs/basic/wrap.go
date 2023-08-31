package basic

// Combine two functions as one.
func Combine[T, R, S any](first func(T) R, second func(R) S) func(T) S {
	return func(t T) S { return second(first(t)) }
}

// CombineResult combines func f and result r
func CombineResult[T, R any](f func(T), r R) func(T) R {
	return func(t T) R { f(t); return r }
}
