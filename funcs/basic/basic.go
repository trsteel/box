package basic

// Ternary is a 1 line if/else statement.
func Ternary[T any](cond bool, t T, f T) T {
	if cond {
		return t
	}
	return f
}

// ContainsKey returns whether key is present in map.
func ContainsKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ContainsAnyKey returns whether any key is present in map.
func ContainsAnyKey[K comparable, V any](m map[K]V, keys []K) bool {
	return !ForEach(keys, func(key K) bool { return !ContainsKey(m, key) })
}

// ContainsAllKeys returns whether all keys are present in map.
func ContainsAllKeys[K comparable, V any](m map[K]V, keys []K) bool {
	return ForEach(keys, func(key K) bool { return ContainsKey(m, key) })
}
