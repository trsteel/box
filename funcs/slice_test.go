package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5}
	Reverse(slice)
	assert.Equal(t, []int{5, 4, 3, 2, 1}, slice)
}
