package tree

import (
	"fmt"
)

type IndexLevel[SV comparable] interface {
	Level[uint32]
	ToStoreSlice() ([]SV, error)
	AddVertexWithStore(v SV) error
	AddEdgeWithStore(source, target SV) error
}

type indexLevel[SV comparable, T IndexTree[SV]] struct {
	level[uint32, T]
}

func NewIndexLevel[SV comparable](index int16, tree IndexTree[SV]) (IndexLevel[SV], error) {
	if tree == nil {
		return nil, fmt.Errorf("tree.Level.NewIndexLevel: tree cannot be null")
	}
	return &indexLevel[SV, IndexTree[SV]]{
		level: level[uint32, IndexTree[SV]]{
			index:    index,
			tree:     tree,
			vertexes: make(map[uint32]Vertex[uint32]),
		},
	}, nil
}

// ToStoreSlice ...
func (l *indexLevel[SV, T]) ToStoreSlice() ([]SV, error) {
	s := make([]SV, 0, l.Len())

	for k := range l.vertexes {
		if v, ok2 := l.tree.store().GetValue(k); ok2 {
			s = append(s, v)
		}
	}

	return s, nil
}

// AddVertexWithStore ...
func (l *indexLevel[SV, T]) AddVertexWithStore(value SV) error {
	l.AddVertex(l.tree.store().Add(value))
	return nil
}

// AddEdgeWithStore ...
func (l *indexLevel[SV, T]) AddEdgeWithStore(source, target SV) error {
	e0 := l.tree.store().Add(source)
	e1 := l.tree.store().Add(target)
	l.AddEdge(e0, e1)
	return nil
}
