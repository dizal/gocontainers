package tree

import (
	"fmt"
	"strings"

	"github.com/dizal/gocontainers/set"
)

// Level ...
type Level struct {
	id int16
	l  map[interface{}]*Vertex
	t  *Tree
}

// Each ...
func (l *Level) Each(f func(i interface{}, v *Vertex) bool) {
	for v := range l.l {
		f(v, l.l[v])
	}
}

// Get ...
func (l *Level) Get(v interface{}) (*Vertex, bool) {
	vv, ok := l.l[v]
	return vv, ok
}

// Contain ...
func (l *Level) Contain(v interface{}) bool {
	_, ok := l.l[v]
	return ok
}

// Len ...
func (l *Level) Len() int {
	return len(l.l)
}

// ToSlice ...
func (l *Level) ToSlice() []interface{} {
	s := make([]interface{}, 0, l.Len())

	for v := range l.l {
		s = append(s, v)
	}

	return s
}

// ToStoreSlice ...
func (l *Level) ToStoreSlice() []interface{} {
	s := make([]interface{}, 0, l.Len())

	for k := range l.l {
		if kk, ok := k.(uint32); ok {
			if v, ok2 := l.t.Store.GetValue(kk); ok2 {
				s = append(s, v)
			}
		}

	}

	return s
}

// AddVertexWithStore ...
func (l *Level) AddVertexWithStore(v string) {
	l.AddVertex(l.t.Store.Store(v))
}

// AddVertex ...
func (l *Level) AddVertex(vertex interface{}) *Vertex {
	if v, ok := l.Get(vertex); ok {
		return v
	}

	l.t.CountVertex++
	v := NewVertex()
	l.l[vertex] = v

	return v
}

// AddEdgeWithStore ...
func (l *Level) AddEdgeWithStore(source, target interface{}) {
	e0 := l.t.Store.Store(source)
	e1 := l.t.Store.Store(target)
	l.AddEdge(e0, e1)
}

// AddEdge ...
func (l *Level) AddEdge(source, target interface{}) {
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
	l1, l2 := l.t.Level(l.id-1), l.t.Level(l.id-2)

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
		l.t.CountEdges++
		return
	}

	if B {
		l.AddVertex(target).AddParent(source)
		vB.AddChild(target)
		l.t.CountEdges++
	} else if D {
		l.AddVertex(source).AddParent(target)
		vD.AddChild(source)
		l.t.CountEdges++
	}
}

// GetRecursionSibling ...
func (l *Level) GetRecursionSibling(checked *set.Set, vertex interface{}) {
	checked.Add(vertex)
	if v, ok := l.Get(vertex); ok {
		v.Siblings.Each(func(vv interface{}) bool {
			if !checked.Contain(vv) {
				l.GetRecursionSibling(checked, vv)
			}
			return true
		})
	}
}

// GetVertexDegreeMarked ...
func (l *Level) GetVertexDegreeMarked(vertex interface{}) int {
	if v, ok := l.Get(vertex); ok {
		d := 0
		pLevel, cLevel := l.t.Level(l.id-1), l.t.Level(l.id+1)

		v.Parents.Each(func(parent interface{}) bool {
			if vv, ok := pLevel.Get(parent); ok && vv.Marked {
				d++
			}
			return true
		})
		v.Children.Each(func(child interface{}) bool {
			if vv, ok := cLevel.Get(child); ok && vv.Marked {
				d++
			}
			return true
		})
		v.Siblings.Each(func(sib interface{}) bool {
			if vv, ok := l.Get(sib); ok && vv.Marked {
				d++
			}
			return true
		})

		return d
	}

	return 0
}

// intersect ...
func intersect(l1, l2 *Level) []interface{} {
	var c []interface{}

	l2.Each(func(i interface{}, _ *Vertex) bool {
		if l1.Contain(i) {
			c = append(c, i)
		}
		return true
	})

	return c
}

func (l *Level) String() string {
	var buffer strings.Builder
	buffer.WriteString(fmt.Sprintf("L%v{\n", l.id))

	l.Each(func(i interface{}, v *Vertex) bool {
		buffer.WriteString(fmt.Sprintf("\t%v = %v\n", i, v))
		return true
	})

	buffer.WriteString("}\n")

	return buffer.String()
}
