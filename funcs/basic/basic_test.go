package basic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	assert.True(t, Ternary(true, true, false))
	assert.False(t, Ternary(false, true, false))
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
