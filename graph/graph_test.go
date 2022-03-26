package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGraphNew(t *testing.T) {
	g := New[string](Undirected)
	assert.Equal(t, 0, g.VertexesCount())
}

func TestGraphAddEdge(t *testing.T) {
	g := New[string](Undirected).
		AddEdge("v1", "v2").
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v1", "v3")

	assert.Equal(t, 3, g.VertexesCount())
	assert.Equal(t, 4, g.EdgesCount())

	g = New[string](Directed).
		AddEdge("v1", "v2").
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v1", "v3")

	assert.Equal(t, 3, g.VertexesCount())
	assert.Equal(t, 2, g.EdgesCount())
}

func TestGraphGetTarget(t *testing.T) {
	g := New[string](Undirected).
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v2", "v3")

	assert.Equal(t, 3, g.VertexesCount())
	assert.Equal(t, 6, g.EdgesCount())

	s, ok := g.GetTargetNodes("v3")
	assert.True(t, ok)
	assert.Equal(t, 2, s.Len())
	s.Each(func(value string) bool {
		assert.Contains(t, []string{"v1", "v2"}, value)
		return true
	})

	g = New[string](Directed).
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v2", "v3")

	assert.Equal(t, 3, g.VertexesCount())
	assert.Equal(t, 3, g.EdgesCount())

	s, ok = g.GetTargetNodes("v3")
	assert.True(t, ok)
	assert.Equal(t, 0, s.Len())
}

func TestGraphErase(t *testing.T) {
	g := New[string](Undirected).
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v2", "v3")

	g.Erase()

	assert.Equal(t, 0, g.VertexesCount())
	assert.Equal(t, 0, g.EdgesCount())
}

func TestGraphRange(t *testing.T) {
	g := New[string](Undirected).
		AddEdge("v1", "v2").
		AddEdge("v1", "v3").
		AddEdge("v2", "v3")

	g.Range("v3", func(target string) bool {
		assert.Contains(t, []string{"v1", "v2"}, target)
		return true
	})

}
