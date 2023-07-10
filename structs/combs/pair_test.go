package combs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPair_Clone(t *testing.T) {
	p := NewPair(1, 'a')
	cp := p.Clone()
	assert.Equal(t, p, cp)
	cp.First, cp.Second = 2, 'b'
	assert.NotEqual(t, p, cp)
}
