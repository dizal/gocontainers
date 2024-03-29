package set_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/dizal/gocontainers/set"
)

func TestSetNew(t *testing.T) {
	assert.True(t, New[string]().IsEmpty())
}

func TestSetAdd(t *testing.T) {
	s := New[string]()
	assert.True(t, s.Add("value"))
	assert.False(t, s.Add("value"))
	assert.Equal(t, 1, s.Len())
}

func TestSetContain(t *testing.T) {
	s := New[string]()
	assert.True(t, s.Add("value"))
	assert.True(t, s.Contain("value"))
}

func TestSetDelete(t *testing.T) {
	s := New[string]()
	assert.True(t, s.Add("value"))
	assert.True(t, s.Add("value2"))
	assert.True(t, s.Delete("value2"))
	assert.False(t, s.Contain("value2"))
	assert.Equal(t, 1, s.Len())
	assert.False(t, s.Delete("value2"))
}

func TestSetErase(t *testing.T) {
	s := New[string]()
	assert.True(t, s.Add("value1"))
	assert.True(t, s.Add("value2"))
	assert.True(t, s.Add("value3"))
	assert.True(t, s.Add("value4"))

	assert.Equal(t, 4, s.Len())

	s.Erase()

	assert.False(t, s.Contain("value1"))
	assert.True(t, s.IsEmpty())
}

func TestSetSlice(t *testing.T) {
	s := New[string]()
	assert.True(t, s.Add("value1"))
	assert.True(t, s.Add("value2"))

	slice := s.ToSlice()
	assert.Len(t, slice, 2)
}

func TestSetString(t *testing.T) {
	s := New[string]()

	assert.Equal(t, "SET<>", s.String())

	assert.True(t, s.Add("v1"))
	assert.True(t, s.Add("v2"))

	assert.Contains(t, []string{"SET<v1, v2>", "SET<v2, v1>"}, s.String())
}

func TestSetEach(t *testing.T) {
	s := New[string]()

	assert.True(t, s.Add("v1"))
	assert.True(t, s.Add("v2"))
	assert.True(t, s.Add("v3"))

	s.Each(func(v string) bool {
		return assert.Contains(t, []string{"v1", "v2", "v3"}, v)
	})

	i := 0
	s.Each(func(v string) bool {
		i++
		return false
	})
	assert.Equal(t, 1, i)
}

func BenchmarkSetAdd(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		set := New[string]()
		b.StartTimer()
		for _, elem := range testSet {
			set.Add(elem)
		}
	}
}
