package automic

import (
	"sync/atomic"

	basicfuncs "github.com/trsteel/box/funcs/basic"
)

// Value wraps automic.Value with certain value type.
type Value[T any] struct {
	atomic.Value
}

// Load returns the value set by the most recent Store.
func (v *Value[T]) Load() (val T, loaded bool) {
	value := v.Value.Load()
	return basicfuncs.ReflectIF[T](value, value != nil)
}

// Store sets the value of the Value v to val.
func (v *Value[T]) Store(val T) {
	v.Value.Store(val)
}

// Swap stores new into Value and returns the previous value.
func (v *Value[T]) Swap(new T) (old T, loaded bool) {
	value := v.Value.Swap(new)
	return basicfuncs.ReflectIF[T](value, value != nil)
}

// CompareAndSwap executes the compare-and-swap operation for the Value.
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return v.Value.CompareAndSwap(old, new)
}
