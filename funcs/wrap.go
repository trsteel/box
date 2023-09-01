package funcs

// Combine two functions as one.
func Combine[T, R, S any](first func(T) R, second func(R) S) func(T) S {
	return func(t T) S { return second(first(t)) }
}

// WithResult combines func f and result r.
func WithResult[T, R any](f func(T), r R) func(T) R {
	return func(t T) R { f(t); return r }
}

// WithArg injects one arg to func f.
func WithArg[T, R any](f func() R) func(T) R {
	return func(_ T) R { return f() }
}

// CopyFunc transfer value to a function.
func CopyFunc[T any](val T) func() T {
	return func() T { return val }
}
