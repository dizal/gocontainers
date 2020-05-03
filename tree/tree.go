package tree

import (
	"github.com/dizal/gocontainers/indexstore"
)

// Tree ...
//    root          level 0
//   /    \
//  A      B        level 1
// / \    / \
// C  D---E  F      level 2
type Tree struct {
	levels map[int16]*Level
	Store  *indexstore.IndexStore

	// info
	CountVertex, CountEdges uint32
}

// New - создание нового дерева
func New(s *indexstore.IndexStore) *Tree {
	t := &Tree{
		levels: make(map[int16]*Level),
		Store:  s,
	}
	return t
}

// Level return level from Tree
func (t *Tree) Level(level int16) *Level {
	if l, ok := t.levels[level]; ok {
		return l
	}

	l := &Level{
		id: level,
		l:  make(map[uint32]*Vertex),
		t:  t,
	}
	t.levels[level] = l

	return l
}

// Erase ..
func (t *Tree) Erase() {
	t.levels = make(map[int16]*Level)
	t.CountEdges = 0
	t.CountVertex = 0
}
