package tree

import "github.com/dizal/gocontainers/set"

// Level ...
type Level struct {
	id int16
	l  map[uint32]*Vertex
	t  *Tree
}

// Each ...
func (l *Level) Each(f func(i uint32, v *Vertex) bool) {
	for v := range l.l {
		f(v, l.l[v])
	}
}

// Get ...
func (l *Level) Get(v uint32) (*Vertex, bool) {
	vv, ok := l.l[v]
	return vv, ok
}

// Contain ...
func (l *Level) Contain(v uint32) bool {
	_, ok := l.l[v]
	return ok
}

// Len ...
func (l *Level) Len() int {
	return len(l.l)
}

// ToSlice ...
func (l *Level) ToSlice() []uint32 {
	s := make([]uint32, 0, l.Len())

	for v := range l.l {
		s = append(s, v)
	}

	return s
}

// ToStoreSlice ...
func (l *Level) ToStoreSlice() []interface{} {
	s := make([]interface{}, 0, l.Len())

	for k := range l.l {
		if v, ok := l.t.Store.GetValue(k); ok {
			s = append(s, v)
		}
	}

	return s
}

// AddVertexWithStore ...
func (l *Level) AddVertexWithStore(v string) {
	l.AddVertex(l.t.Store.Store(v))
}

// AddVertex ...
func (l *Level) AddVertex(vertex uint32) *Vertex {
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
func (l *Level) AddEdge(source, target uint32) {
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
func (l *Level) GetRecursionSibling(checked *set.Set, vertex uint32) {
	checked.Add(vertex)
	if v, ok := l.Get(vertex); ok {
		v.Siblings.Each(func(vv interface{}) bool {
			l.GetRecursionSibling(checked, vv.(uint32))
			return true
		})
	}
}

// GetVertexDegreeMarked ...
func (l *Level) GetVertexDegreeMarked(vertex uint32) int {
	if v, ok := l.Get(vertex); ok {
		d := 0
		pLevel, cLevel := l.t.Level(l.id-1), l.t.Level(l.id+1)

		v.Parents.Each(func(parent interface{}) bool {
			if vv, ok := pLevel.Get(parent.(uint32)); ok && vv.Marked {
				d++
			}
			return true
		})
		v.Children.Each(func(child interface{}) bool {
			if vv, ok := cLevel.Get(child.(uint32)); ok && vv.Marked {
				d++
			}
			return true
		})
		v.Siblings.Each(func(sib interface{}) bool {
			if vv, ok := l.Get(sib.(uint32)); ok && vv.Marked {
				d++
			}
			return true
		})

		return d
	}

	return 0
}

// intersect ...
func intersect(l1, l2 *Level) []uint32 {
	var c []uint32

	l2.Each(func(i uint32, _ *Vertex) bool {
		if l1.Contain(i) {
			c = append(c, i)
		}
		return true
	})

	return c
}
