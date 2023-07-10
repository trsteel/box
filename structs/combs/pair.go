package combs

// Pair is a combination of two elements.
type Pair[T any, R any] struct {
	First  T
	Second R
}

// NewPair returns a struct pointer.
func NewPair[T any, R any](first T, second R) *Pair[T, R] {
	return &Pair[T, R]{First: first, Second: second}
}

// Clone returns a deepcopy object.
func (p *Pair[T, R]) Clone() *Pair[T, R] {
	return NewPair(p.First, p.Second)
}
