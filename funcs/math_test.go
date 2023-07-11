package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
