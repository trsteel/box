package set

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Add(t *testing.T) {
	st := NewSet(1)
	assert.Equal(t, 1, st.Len())
	st.Add(1)
	assert.Equal(t, 1, st.Len())
	st.Add(2, 3)
	assert.Equal(t, 3, st.Len())
}

func TestSet_Contains(t *testing.T) {
	st := NewSet(1, 2, 3)
	assert.True(t, st.Contains(1))
	assert.False(t, st.Contains(0))
}

func TestSet_ContainsAny(t *testing.T) {
	st := NewSet(1, 2, 3)
	assert.True(t, st.ContainsAny([]int{0, 1, 2}))
	assert.False(t, st.ContainsAny([]int{0, -1, -2}))
}

func TestSet_ContainsAll(t *testing.T) {
	st := NewSet(1, 2, 3)
	assert.True(t, st.ContainsAll([]int{1, 2, 3}))
	assert.False(t, st.ContainsAll([]int{0, 1, 2}))
}
