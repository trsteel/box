package funcs

// Ternary is a 1 line if/else statement.
func Ternary[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// ReflectIF returns value after reflect or zero if reflect failed.
func ReflectIF[T any](value any, cond bool) (T, bool) {
	if cond {
		return value.(T), true
	}
	return Zero[T](), false
}

// ReflectOR returns value after reflect or the fallback value if reflect failed.
func ReflectOR[T any](value any, fallback T) T {
	if v, ok := value.(T); ok {
		return v
	}
	return fallback
}

// Swap two args.
func Swap[T any, R any](first T, second R) (R, T) {
	return second, first
}
