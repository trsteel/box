package automic

import (
	"sync/atomic"

	"github.com/trsteel/box/funcs"
)

// Value wraps automic.Value with certain value type.
type Value[T any] struct {
	atomic.Value
}

// Load returns the value set by the most recent Store.
func (v *Value[T]) Load() (val T, loaded bool) {
	if v := v.Value.Load(); v != nil {
		return v.(T), true
	}
	return funcs.Zero[T](), false
}

// Store sets the value of the Value v to val.
func (v *Value[T]) Store(val T) {
	v.Value.Store(val)
}

// Swap stores new into Value and returns the previous value.
func (v *Value[T]) Swap(new T) (old T, loaded bool) {
	if v := v.Value.Swap(new); v != nil {
		return v.(T), true
	}
	return funcs.Zero[T](), false
}

// CompareAndSwap executes the compare-and-swap operation for the Value.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.Value.CompareAndSwap(old, new)
}
