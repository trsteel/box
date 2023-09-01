package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEqual(t *testing.T) {
	assert.True(t, Equal(1, 1))
	assert.False(t, Equal("1", "2"))
}

func TestLessThan(t *testing.T) {
	assert.False(t, LessThan(1, 1))
	assert.False(t, LessThan(2, 1))
	assert.True(t, LessThan(1, 2))
}

func TestLargerThan(t *testing.T) {
	assert.False(t, LargerThan(1, 1))
	assert.True(t, LargerThan(2, 1))
	assert.False(t, LargerThan(1, 2))
}

func TestBetween(t *testing.T) {
	assert.Equal(t, true, Between(0, 0, 2))
	assert.Equal(t, false, Between(-1, 0, 2))
}

func TestMax(t *testing.T) {
	assert.Equal(t, 3, Max(1, 2, 3))
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, Min(1, 2, 3))
}
