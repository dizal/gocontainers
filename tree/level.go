package tree

import (
	"fmt"
	"strings"

	"github.com/dizal/gocontainers/set"
)

type Level[V comparable] interface {
	Each(f func(vertexID V, vertex Vertex[V]) bool)
	Get(vertexID V) (Vertex[V], bool)
	Contain(vertexID V) bool
	Len() int
	ToSlice() []V
	AddVertex(vertexID V) Vertex[V]
	AddEdge(source, target V)
	GetRecursionSibling(checked set.Set[V], vertexID V)
	GetVertexDegreeMarked(vertex V) int
	String() string
}

// Level ...
type level[V comparable, T Tree[V]] struct {
	index    int16
	vertexes map[V]Vertex[V]
	tree     T
}

// NewLevel ...
func NewLevel[V comparable](index int16, tree Tree[V]) (Level[V], error) {
	if tree == nil {
		return nil, fmt.Errorf("tree.Level.NewLevel: tree cannot be null")
	}

	return &level[V, Tree[V]]{
		index:    index,
		tree:     tree,
		vertexes: make(map[V]Vertex[V]),
	}, nil
}

// Each ...
func (l *level[V, T]) Each(f func(vertexID V, vertex Vertex[V]) bool) {
	for v := range l.vertexes {
		f(v, l.vertexes[v])
	}
}

// Get ...
func (l *level[V, T]) Get(vertexID V) (Vertex[V], bool) {
	vv, ok := l.vertexes[vertexID]
	return vv, ok
}

// Contain ...
func (l *level[V, T]) Contain(vertexID V) bool {
	_, ok := l.vertexes[vertexID]
	return ok
}

// Len ...
func (l *level[V, T]) Len() int {
	return len(l.vertexes)
}

// ToSlice ...
func (l *level[V, T]) ToSlice() []V {
	s := make([]V, 0, l.Len())

	for v := range l.vertexes {
		s = append(s, v)
	}

	return s
}

// AddVertex ...
func (l *level[V, SV]) AddVertex(vertexID V) Vertex[V] {
	if v, ok := l.Get(vertexID); ok {
		return v
	}

	l.tree.addVetrex()
	v := NewVertex[V]()
	l.vertexes[vertexID] = v

	return v
}

// AddEdge ...
func (l *level[V, T]) AddEdge(source, target V) {
	// | old | cur | new |
	// |level|level|level|
	// |_____|_____|_____|
	// | l2  | l1  |  l  |
	// |-----|-----|-----|
	// |  A←---→B  |     | this Edge is exists
	// |  C←---→D  |     | this Edge is exists
	// |     |     |     |
	// |     |  B  |     |
	// |     |  ↕  |     | sibling
	// |     |  D  |     |
	// |     |     |     |
	// |     |  B←---→X  | new edge
	// |     |  D←---→X  |
	l1, l2 := l.tree.L(l.index-1), l.tree.L(l.index-2)

	_, A := l2.Get(target)
	vB, B := l1.Get(source)
	_, C := l2.Get(source)
	vD, D := l1.Get(target)

	// exists
	if (B && A) || (C && D) {
		return
	}

	// sibling
	if B && D {
		vB.AddSibling(target)
		vD.AddSibling(source)
		l.tree.addEdge()
		return
	}

	if B {
		l.AddVertex(target).AddParent(source)
		vB.AddChild(target)
		l.tree.addEdge()
	} else if D {
		l.AddVertex(source).AddParent(target)
		vD.AddChild(source)
		l.tree.addEdge()
	}
}

// GetRecursionSibling ...
func (l *level[V, T]) GetRecursionSibling(checked set.Set[V], vertexID V) {
	checked.Add(vertexID)
	if v, ok := l.Get(vertexID); ok {
		v.Siblings().Each(func(vv V) bool {
			if !checked.Contain(vv) {
				l.GetRecursionSibling(checked, vv)
			}
			return true
		})
	}
}

// GetVertexDegreeMarked ...
func (l *level[V, T]) GetVertexDegreeMarked(vertexID V) int {
	if v, ok := l.Get(vertexID); ok {
		d := 0
		pLevel, cLevel := l.tree.L(l.index-1), l.tree.L(l.index+1)

		v.Parents().Each(func(parent V) bool {
			if vv, ok := pLevel.Get(parent); ok && vv.IsMarked() {
				d++
			}
			return true
		})

		v.Children().Each(func(child V) bool {
			if vv, ok := cLevel.Get(child); ok && vv.IsMarked() {
				d++
			}
			return true
		})

		v.Siblings().Each(func(sib V) bool {
			if vv, ok := l.Get(sib); ok && vv.IsMarked() {
				d++
			}
			return true
		})

		return d
	}

	return 0
}

func (l *level[V, T]) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("L%v{\n", l.index))

	l.Each(func(i V, v Vertex[V]) bool {
		buffer.WriteString(fmt.Sprintf("\t%v = %v\n", i, v))
		return true
	})

	buffer.WriteString("}\n")

	return buffer.String()
}
