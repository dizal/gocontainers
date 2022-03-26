package tree_test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	. "github.com/dizal/gocontainers/tree"
)

func TestVertex(t *testing.T) {
	v := NewVertex[string]()
	assert.Equal(t, 0, v.Degree())
	println(v.String())
}

func TestVertexAdd(t *testing.T) {
	v := NewVertex[string]()

	assert.True(t, v.AddChild("c1"))
	assert.False(t, v.AddChild("c1"))
	assert.True(t, v.AddChild("p1"))
	assert.False(t, v.AddChild("p1"))
	assert.True(t, v.AddChild("s1"))
	assert.False(t, v.AddChild("s1"))

	assert.Equal(t, 3, v.Degree())

	println(v.String())
}

func TestVertexMark(t *testing.T) {
	v := NewVertex[string]()
	v.Mark()
	assert.True(t, v.IsMarked())
}

func TestVertexAddParent(t *testing.T) {
	v := NewVertex[string]()

	assert.True(t, v.AddParent("p1"))
	assert.False(t, v.AddParent("p1"))
	assert.True(t, v.AddParent("p2"))

	assert.Equal(t, 2, v.Parents().Len())

	v.Parents().Each(func(value string) bool {
		assert.Contains(t, []string{"p1", "p2"}, value)
		return true
	})
}

func TestVertexAddChild(t *testing.T) {
	v := NewVertex[string]()

	assert.True(t, v.AddChild("p1"))
	assert.False(t, v.AddChild("p1"))
	assert.True(t, v.AddChild("p2"))

	assert.Equal(t, 2, v.Children().Len())

	v.Children().Each(func(value string) bool {
		assert.Contains(t, []string{"p1", "p2"}, value)
		return true
	})
}

func TestVertexAddSibling(t *testing.T) {
	v := NewVertex[string]()

	assert.True(t, v.AddSibling("p1"))
	assert.False(t, v.AddSibling("p1"))
	assert.True(t, v.AddSibling("p2"))

	assert.Equal(t, 2, v.Siblings().Len())

	v.Siblings().Each(func(value string) bool {
		assert.Contains(t, []string{"p1", "p2"}, value)
		return true
	})
}

func TestVertexDegree(t *testing.T) {
	v := NewVertex[string]()

	assert.True(t, v.AddParent("p1"))
	assert.False(t, v.AddParent("p1"))
	assert.True(t, v.AddParent("p2"))

	assert.True(t, v.AddChild("p1"))
	assert.False(t, v.AddChild("p1"))
	assert.True(t, v.AddChild("p2"))

	assert.True(t, v.AddSibling("p1"))
	assert.False(t, v.AddSibling("p1"))
	assert.True(t, v.AddSibling("p2"))

	assert.Equal(t, 6, v.Degree())
}

func BenchmarkVertexAdd(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	v := NewVertex[string]()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			v.AddSibling(elem)
		}
	}
}
