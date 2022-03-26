package tree

import (
	"fmt"
	"strings"
)

type Tree[V comparable] interface {
	L(levelIndex int16) Level[V]
	Erase()
	String() string

	addVetrex()
	addEdge()
}

// Tree ...
//    root          level 0
//   /    \
//  A      B        level 1
// / \    / \
// C  D---E  F      level 2
type tree[V comparable, L Level[V]] struct {
	levels      map[int16]L
	countVertex uint32
	countEdges  uint32
}

func New[V comparable]() Tree[V] {
	return &tree[V, Level[V]]{
		levels: make(map[int16]Level[V]),
	}
}

// Level return level from Tree
func (t *tree[V, L]) L(levelIndex int16) Level[V] {
	if l, ok := t.levels[levelIndex]; ok {
		return l
	}

	l, _ := NewLevel[V](levelIndex, t)

	t.levels[levelIndex] = l.(L)

	return l
}

// Erase ..
func (t *tree[V, L]) Erase() {
	t.levels = make(map[int16]L)
	t.countEdges = 0
	t.countVertex = 0
}

func (t *tree[V, L]) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("Tree: V:%v, E:%v (\n", t.countVertex, t.countEdges))

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

func (t *tree[V, L]) addVetrex() {
	t.countVertex++
}

func (t *tree[V, L]) addEdge() {
	t.countEdges++
}
