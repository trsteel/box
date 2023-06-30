package sortedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortedMap_Range(t *testing.T) {
	sortedMap := SortedMap[string, string](map[string]string{"a": "aa", "b": "bb", "c": "cc"})
	var keys, values []string
	sortedMap.Range(func(key string, val string) bool {
		keys = append(keys, key)
		values = append(values, val)
		return true
	})
	assert.Equal(t, []string{"a", "b", "c"}, keys)
	assert.Equal(t, []string{"aa", "bb", "cc"}, values)
}
