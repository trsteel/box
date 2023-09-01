package funcs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsAny(t *testing.T) {
	assert.True(t, ContainsAny([]int{1, 2, 3, 4, 5}, []int{5, 6, 7}))
	assert.False(t, ContainsAny([]int{1, 2, 3, 4, 5}, []int{6, 7}))
}

func TestContainsAll(t *testing.T) {
	assert.True(t, ContainsAll([]int{1, 2, 3, 4, 5}, []int{3, 4, 5}))
	assert.False(t, ContainsAll([]int{1, 2, 3, 4, 5}, []int{4, 5, 6}))
}

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	Reverse(slice)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, slice)
}

func TestMap(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, Map([]int{1, 2, 3}, strconv.Itoa))
}

func TestMapPtr(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, MapPtr([]int{1, 2, 3}, Combine(FromPtr[int], strconv.Itoa)))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, []int{1, 1, 1}, Repeat(3, 1))
}
