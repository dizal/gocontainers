package indexstore

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStoreNew(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Size())
}

func TestStoreAdd(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Add("value"))
	assert.EqualValues(t, 0, s.Add("value"))
	assert.EqualValues(t, 0, s.Add("value"))

	assert.EqualValues(t, 1, s.Size())
}

func TestStoreGetIndex(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Add("value1"))
	assert.EqualValues(t, 1, s.Add("value2"))
	assert.EqualValues(t, 2, s.Add("value3"))

	v, ok := s.GetIndex("value2")
	assert.True(t, ok)
	assert.EqualValues(t, 1, v)

	assert.EqualValues(t, 3, s.Size())
}

func TestStoreGetValue(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Add("value1"))
	assert.EqualValues(t, 1, s.Add("value2"))
	assert.EqualValues(t, 2, s.Add("value3"))

	v, ok := s.GetValue(1)
	assert.True(t, ok)
	assert.Equal(t, "value2", v)

	assert.EqualValues(t, 3, s.Size())

	v, ok = s.GetValue(999)
	assert.False(t, ok)
	assert.Equal(t, "", v)
}

func TestStoreDelete(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Add("value1"))
	assert.EqualValues(t, 1, s.Add("value2"))
	assert.EqualValues(t, 2, s.Add("value3"))

	assert.EqualValues(t, 3, s.Size())

	s.Delete("value1")
	s.Delete("value2")

	assert.EqualValues(t, 1, s.Size())

	_, ok := s.GetIndex("value2")
	assert.False(t, ok)

	_, ok = s.GetIndex("value3")
	assert.True(t, ok)

	_, ok = s.GetValue(1)
	assert.False(t, ok)
}

func TestStoreErase(t *testing.T) {
	s := New[string]()
	assert.EqualValues(t, 0, s.Add("value1"))
	assert.EqualValues(t, 1, s.Add("value2"))
	assert.EqualValues(t, 2, s.Add("value3"))

	s.Erase()

	assert.EqualValues(t, 0, s.Add("value3"))
	assert.EqualValues(t, 1, s.Size())
}

func BenchmarkStoreAdd(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	store := New[string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			store.Add(elem)
		}
	}
}

func BenchmarkStoreDelete(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		store := New[string]()
		for _, elem := range testSet {
			store.Add(elem)
		}
		b.StartTimer()
		for _, elem := range testSet {
			store.Delete(elem)
		}
	}
}
