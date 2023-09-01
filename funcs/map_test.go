package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedKeys(t *testing.T) {
	assert.Equal(t, []string{"a", "b", "c"}, SortedKeys(map[string]struct{}{"c": {}, "b": {}, "a": {}}))
}

func TestContainsKey(t *testing.T) {
	assert.True(t, ContainsKey(map[string]struct{}{"a": {}}, "a"))
	assert.False(t, ContainsKey(map[string]struct{}{}, "a"))
}

func TestContainsAnyKey(t *testing.T) {
	assert.True(t, ContainsAnyKey(map[string]struct{}{"a": {}}, []string{"a", "b"}))
	assert.False(t, ContainsAnyKey(map[string]struct{}{"a": {}}, []string{"b", "c"}))
}

func TestContainsAllKeys(t *testing.T) {
	assert.True(t, ContainsAllKeys(map[string]struct{}{"a": {}}, []string{"a"}))
	assert.False(t, ContainsAllKeys(map[string]struct{}{"a": {}}, []string{"a", "b"}))
}

func TestForEachKV(t *testing.T) {
	assert.True(t, ForEachKV(map[string]bool{}, func(s string, b bool) bool { return b }))
	assert.True(t, ForEachKV(map[string]bool{"a": true, "b": true}, func(s string, b bool) bool { return b }))
	assert.False(t, ForEachKV(map[string]bool{"a": true, "b": false}, func(s string, b bool) bool { return b }))
}

func TestMerge(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 1, "b": 3, "d": 4}, Merge(
		map[string]int{"a": 1, "b": 2}, map[string]int{"b": 3, "d": 4},
	))
}
