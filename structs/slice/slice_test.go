package slice

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlice_Clone(t *testing.T) {
	slice := New(1, 2, 3, 4, 5)
	slice2 := slice.Clone()
	assert.Equal(t, slice, slice2)
	slice2[0] = -1
	assert.NotEqual(t, slice, slice2)
}

func TestSlice_Reverse(t *testing.T) {
	slice := New(1, 2, 3, 4, 5)
	slice.Reverse()
	assert.Equal(t, slice, New(5, 4, 3, 2, 1))
}

func TestSlice_SortBy(t *testing.T) {
	slice := New(4, 5, 1, 2, 3)
	slice.SortBy(func(a, b int) bool { return a > b })
	assert.Equal(t, slice, New(5, 4, 3, 2, 1))
}
