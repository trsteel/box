package funcs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	basicfuncs "github.com/trsteel/box/funcs/basic"
)

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	Reverse(slice)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, slice)
}

func TestMap(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, Map([]int{1, 2, 3}, strconv.Itoa))
}

func TestMapPtr(t *testing.T) {
	assert.Equal(t, []string{"1", "2", "3"}, MapPtr([]int{1, 2, 3}, basicfuncs.Combine(FromPtr[int], strconv.Itoa)))
}

func TestRepeat(t *testing.T) {
	assert.Equal(t, []int{1, 1, 1}, Repeat(3, 1))
}
