package syncmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestSyncMap_Delete(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// before add
	syncMap.Delete("a")

	syncMap.Store("a", "aa")
	// after add
	value, ok := syncMap.Load("a")
	assert.Equal(t, "aa", value)
	assert.True(t, ok)

	syncMap.Delete("a")
	// after delete
	value, ok = syncMap.Load("a")
	assert.Equal(t, "", value)
	assert.False(t, ok)
}

func TestSyncMap_Load(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// before add
	value, ok := syncMap.Load("a")
	assert.Equal(t, "", value)
	assert.False(t, ok)

	syncMap.Store("a", "aa")
	// after add
	value, ok = syncMap.Load("a")
	assert.Equal(t, "aa", value)
	assert.True(t, ok)
}

func TestSyncMap_LoadAndDelete(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// before add
	value, ok := syncMap.LoadAndDelete("a")
	assert.Equal(t, "", value)
	assert.False(t, ok)

	syncMap.Store("a", "aa")
	// after add
	value, ok = syncMap.LoadAndDelete("a")
	assert.Equal(t, "aa", value)
	assert.True(t, ok)
}

func TestSyncMap_LoadOrStore(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// first
	value, ok := syncMap.LoadOrStore("a", "aa")
	assert.Equal(t, "", value)
	assert.False(t, ok)
	// second
	value, ok = syncMap.LoadOrStore("a", "bb")
	assert.Equal(t, "aa", value)
	assert.True(t, ok)
}

func TestSyncMap_Range(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	syncMap.Store("a", "aa")
	syncMap.Store("b", "bb")

	var keys, values []string
	syncMap.Range(func(key string, value string) bool {
		keys = append(keys, key)
		values = append(values, value)
		return true
	})

	slices.Sort(keys)
	assert.Equal(t, []string{"a", "b"}, keys)
	slices.Sort(values)
	assert.Equal(t, []string{"aa", "bb"}, values)
}

func TestSyncMap_Store(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// before add
	value, ok := syncMap.Load("a")
	assert.Equal(t, "", value)
	assert.False(t, ok)

	syncMap.Store("a", "aa")
	// after add
	value, ok = syncMap.Load("a")
	assert.Equal(t, "aa", value)
	assert.True(t, ok)
}
