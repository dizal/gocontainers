package tree

import "github.com/dizal/gocontainers/indexstore"

type IndexTree[SV comparable] interface {
	Tree[uint32]
	store() indexstore.IndexStore[SV]
}

type indexTree[SV comparable, L IndexLevel[SV]] struct {
	tree[uint32, L]
	s indexstore.IndexStore[SV]
}

func NewIndex[SV comparable](s indexstore.IndexStore[SV]) IndexTree[SV] {
	it := &indexTree[SV, IndexLevel[SV]]{
		tree: tree[uint32, IndexLevel[SV]]{
			levels: make(map[int16]IndexLevel[SV]),
		},
		s: s,
	}
	return it
}

func (t *indexTree[SV, L]) Level(levelID int16) Level[uint32] {
	if l, ok := t.levels[levelID]; ok {
		return l
	}

	l, _ := NewIndexLevel[SV](levelID, t)

	t.levels[levelID] = l.(L)

	return l
}

func (t *indexTree[SV, L]) store() indexstore.IndexStore[SV] {
	return t.s
}
