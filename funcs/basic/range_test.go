package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func TestForEachKV(t *testing.T) {
	assert.True(t, ForEachKV(map[string]bool{}, func(s string, b bool) bool { return b }))
	assert.True(t, ForEachKV(map[string]bool{"a": true, "b": true}, func(s string, b bool) bool { return b }))
	assert.False(t, ForEachKV(map[string]bool{"a": true, "b": false}, func(s string, b bool) bool { return b }))
}
