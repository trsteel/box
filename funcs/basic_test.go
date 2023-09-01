package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	assert.True(t, Ternary(true, true, false))
	assert.False(t, Ternary(false, true, false))
}

func TestReflectIF(t *testing.T) {
	v, ok := ReflectIF[string]("aa", true)
	assert.True(t, ok)
	assert.Equal(t, "aa", v)
	v, ok = ReflectIF[string]("aa", false)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestReflectOR(t *testing.T) {
	assert.Equal(t, ReflectOR[string]("aa", "empty"), "aa")
	assert.Equal(t, ReflectOR[string](123, "empty"), "empty")
}

func TestSwap(t *testing.T) {
	first, second := Swap(123, "123")
	assert.Equal(t, first, "123")
	assert.Equal(t, second, 123)
}
