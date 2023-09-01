package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestForRange(t *testing.T) {
	arr := make([]int, 5)
	ForRange(0, len(arr), 1, func(i int) bool { arr[i] = i; return true })
	assert.Equal(t, []int{0, 1, 2, 3, 4}, arr)
	arr = make([]int, 5)
	ForRange(0, len(arr), 2, func(i int) bool { arr[i] = i; return true })
	assert.Equal(t, []int{0, 0, 2, 0, 4}, arr)
}

func TestForEach(t *testing.T) {
	assert.True(t, ForEach([]int{}, func(i int) bool { return true }))
	assert.True(t, ForEach([]int{1, 2, 3}, func(i int) bool { return true }))
	assert.False(t, ForEach([]int{1, 2, 3}, func(i int) bool { return false }))
}

func TestForEachPtr(t *testing.T) {
	assert.True(t, ForEachPtr([]int{}, func(i *int) bool { return true }))
	assert.True(t, ForEachPtr([]int{1, 2, 3}, func(i *int) bool { return true }))
	assert.False(t, ForEachPtr([]int{1, 2, 3}, func(i *int) bool { return false }))
}

func TestReduce(t *testing.T) {
	assert.Equal(t, 5, Reduce(Repeat(5, 1), 0, Sum2[int]))
	assert.Equal(t, "aaaaa", Reduce(Repeat(5, "a"), "", Sum2[string]))
}
