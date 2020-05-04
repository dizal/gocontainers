package tree

import (
	"fmt"
	"strings"

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

// New ...
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
		l:  make(map[interface{}]*Vertex),
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

func (t *Tree) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("Tree: V:%v, E:%v (\n", t.CountVertex, t.CountEdges))

	for level := int16(0); true; level++ {
		if v, ok := t.levels[level]; ok {
			buffer.WriteString(v.String())
		} else {
			break
		}
	}
	buffer.WriteString(")")

	return buffer.String()
}
