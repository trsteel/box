package funcs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToSnakeCase(t *testing.T) {
	assert.Equal(t, "abc123+-*/", ToSnakeCase("ABC123+-*/"))
	assert.Equal(t, "123_abc+-*/", ToSnakeCase("123ABC+-*/"))
}
