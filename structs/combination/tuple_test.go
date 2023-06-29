package combination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTuple_Clone(t *testing.T) {
	tu := NewTuple(1, 'a', "aa")
	ctu := tu.Clone()
	assert.Equal(t, tu, ctu)
	ctu.First, ctu.Second, ctu.Third = 2, 'b', "bb"
	assert.NotEqual(t, tu, ctu)
}
