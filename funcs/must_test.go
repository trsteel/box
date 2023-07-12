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
