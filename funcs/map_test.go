package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 1, "b": 3, "d": 4}, Merge(
		map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "d": 4},
	))
}
