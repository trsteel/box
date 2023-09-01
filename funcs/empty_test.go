package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsZero(t *testing.T) {
	assert.True(t, IsZero(Zero[int]()))
	assert.True(t, IsZero(Zero[string]()))
	assert.True(t, IsZero(Zero[uintptr]()))
}
