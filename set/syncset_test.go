package set_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/dizal/gocontainers/set"
)

func TestSyncSetNew(t *testing.T) {
	assert.True(t, NewSync[string]().IsEmpty())
}

func TestSyncSetAdd(t *testing.T) {
	s := NewSync[string]()
	assert.True(t, s.Add("value"))
	assert.False(t, s.Add("value"))
	assert.Equal(t, 1, s.Len())
}

func TestSyncSetContain(t *testing.T) {
	s := NewSync[string]()
	assert.True(t, s.Add("value"))
	assert.True(t, s.Contain("value"))
}

func TestSyncSetDelete(t *testing.T) {
	s := NewSync[string]()
	assert.True(t, s.Add("value"))
	assert.True(t, s.Add("value2"))
	assert.True(t, s.Delete("value2"))
	assert.False(t, s.Contain("value2"))
	assert.Equal(t, 1, s.Len())
	assert.False(t, s.Delete("value2"))
}

func TestSyncSetErase(t *testing.T) {
	s := NewSync[string]()
	assert.True(t, s.Add("value1"))
	assert.True(t, s.Add("value2"))
	assert.True(t, s.Add("value3"))
	assert.True(t, s.Add("value4"))

	assert.Equal(t, 4, s.Len())

	s.Erase()

	assert.False(t, s.Contain("value1"))
	assert.True(t, s.IsEmpty())
}

func TestSyncSetSlice(t *testing.T) {
	s := NewSync[string]()
	assert.True(t, s.Add("value1"))
	assert.True(t, s.Add("value2"))

	slice := s.ToSlice()
	assert.Len(t, slice, 2)
}

func TestSyncSetEach(t *testing.T) {
	s := NewSync[string]()

	assert.True(t, s.Add("v1"))
	assert.True(t, s.Add("v2"))
	assert.True(t, s.Add("v3"))

	s.Each(func(v string) bool {
		return assert.Contains(t, []string{"v1", "v2", "v3"}, v)
	})
}
