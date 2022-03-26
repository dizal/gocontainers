package tree_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/dizal/gocontainers/indexstore"
	. "github.com/dizal/gocontainers/tree"
)

func TestLevelNew(t *testing.T) {
	_, err := NewLevel[int](0, nil)
	assert.Error(t, err)

	tree := New[int]()

	level, err := NewLevel(0, tree)
	assert.NoError(t, err)
	assert.Equal(t, 0, level.Len())
}

func TestLevelAddVertex(t *testing.T) {
	level, _ := NewLevel(0, New[int]())

	level.AddVertex(1)
	level.AddVertex(1)

	assert.Equal(t, 1, level.Len())
}

func TestLevelAddVertexWithStore(t *testing.T) {
	tree := NewIndex(indexstore.New[string]())
	level, _ := NewIndexLevel(0, tree)

	assert.NoError(t, level.AddVertexWithStore("v1"))
	assert.NoError(t, level.AddVertexWithStore("v1"))

	assert.Equal(t, 1, level.Len())
}

func TestLevelContain(t *testing.T) {
	level, _ := NewLevel(0, New[int]())

	V1, V2 := 1, 2

	level.AddVertex(V1)

	assert.True(t, level.Contain(V1))
	assert.False(t, level.Contain(V2))
}

func TestLevelGet(t *testing.T) {
	level, _ := NewLevel(0, New[int]())

	V1 := 1
	level.AddVertex(V1)

	_, ok := level.Get(V1)
	assert.True(t, ok)
}

func TestLevelSlice(t *testing.T) {
	level, _ := NewLevel(0, New[int]())
	V1, V2 := 1, 2
	level.AddVertex(V1)
	level.AddVertex(V1)
	level.AddVertex(V2)

	s := level.ToSlice()

	assert.Len(t, s, 2)
}

func TestLevelStoreSlice(t *testing.T) {
	level, _ := NewIndexLevel(0, NewIndex(indexstore.New[string]()))
	V1, V2 := "v1", "v2"
	level.AddVertexWithStore(V1)
	level.AddVertexWithStore(V1)
	level.AddVertexWithStore(V2)

	s, err := level.ToStoreSlice()
	assert.NoError(t, err)
	assert.Len(t, s, 2)
}
