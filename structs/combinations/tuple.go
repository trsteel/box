package combinations

// Tuple is a combination of three elements.
type Tuple[T any, R any, S any] struct {
	First  T
	Second R
	Third  S
}

// NewTuple returns a struct pointer.
func NewTuple[T any, R any, S any](first T, second R, third S) *Tuple[T, R, S] {
	return &Tuple[T, R, S]{First: first, Second: second, Third: third}
}

// Clone returns a deepcopy object.
func (t *Tuple[T, R, S]) Clone() *Tuple[T, R, S] {
	return NewTuple(t.First, t.Second, t.Third)
}
