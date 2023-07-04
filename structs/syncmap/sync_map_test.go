package syncmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/slices"
)

func TestSyncMap_CompareAndDelete(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	assert.False(t, syncMap.CompareAndDelete("a", "aa")) // before add key
	syncMap.Store("a", "aa")
	assert.False(t, syncMap.CompareAndDelete("a", "bb")) // old value mismatch
	assert.True(t, syncMap.CompareAndDelete("a", "aa"))  // old value match
	assert.False(t, syncMap.CompareAndDelete("a", "aa")) // after delete
}

func TestSyncMap_CompareAndSwap(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	assert.False(t, syncMap.CompareAndSwap("a", "aa", "bb")) // before add key
	syncMap.Store("a", "aa")
	assert.False(t, syncMap.CompareAndSwap("a", "bb", "cc")) // old value mismatch
	assert.True(t, syncMap.CompareAndSwap("a", "aa", "bb"))  // old value match
	assert.True(t, syncMap.CompareAndSwap("a", "bb", "cc"))  // old value match
}

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
	assert.Equal(t, "aa", value)
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

func TestSyncMap_Swap(t *testing.T) {
	syncMap := SyncMap[string, string]{}
	// first
	previous, loaded := syncMap.Swap("a", "aa")
	assert.False(t, loaded)
	assert.Equal(t, "", previous)
	// second
	previous, loaded = syncMap.Swap("a", "bb")
	assert.True(t, loaded)
	assert.Equal(t, "aa", previous)
}
