package automic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValue(t *testing.T) {
	value := &Value[string]{}

	_, loaded := value.Load()
	assert.False(t, loaded)

	value.Store("a")

	v, loaded := value.Load()
	assert.Equal(t, "a", v)
	assert.True(t, loaded)

	old, loaded := value.Swap("b")
	assert.Equal(t, "a", old)
	assert.True(t, loaded)

	assert.False(t, value.CompareAndSwap("a", "c"))
	assert.True(t, value.CompareAndSwap("b", "c"))
}
