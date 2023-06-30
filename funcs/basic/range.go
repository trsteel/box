package basic

// ForEach apply f to each element of slice until f returns false.
func ForEach[T any](slice []T, f func(T) bool) bool {
	for i := 0; i < len(slice); i++ {
		if !f(slice[i]) {
			return false
		}
	}
	return true
}

// ForEachPtr apply f to each element of slice until f returns false.
func ForEachPtr[T any](slice []T, f func(*T) bool) bool {
	for i := 0; i < len(slice); i++ {
		if !f(&slice[i]) {
			return false
		}
	}
	return true
}

// ForEachKV apply f to each KV pair of map until f returns false.
func ForEachKV[K comparable, V any](m map[K]V, f func(K, V) bool) bool {
	for k, v := range m {
		if !f(k, v) {
			return false
		}
	}
	return true
}
