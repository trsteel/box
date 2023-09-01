package funcs

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMust(t *testing.T) {
	assert.Equal(t, 1, Must(strconv.Atoi("1")))
	assert.Panics(t, func() { Must(strconv.Atoi("a")) })
}

func TestMustTrue(t *testing.T) {
	assert.Equal(t, 1, MustTrue(1, true))
	assert.Panics(t, func() { MustTrue(1, false) })
}
