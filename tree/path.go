package tree

import "github.com/dizal/gocontainers/set"

// SearchShortestPath ...
func SearchShortestPath[V comparable](tA, tB Tree[V], l int16) (set.Set[V], int16) {
	lA := tA.L(l).Len()
	lB := tB.L(l).Len()
	lA2 := tA.L(l - 1).Len()
	lB2 := tB.L(l - 1).Len()

	var inter []V
	bIsShort := lA+lB2 > lB+lA2

	// even length (A - x - (x - x) - x - B)
	if bIsShort {
		inter = intersect(tA.L(l), tB.L(l-1))
	} else {
		inter = intersect(tA.L(l-1), tB.L(l))
	}

	if len(inter) > 0 {
		if bIsShort {
			return makePath(tA, tB, inter, l, l-1)
		}
		return makePath(tA, tB, inter, l-1, l)
	}
	// odd length (A - x - (x) - x - B)
	inter = intersect(tA.L(l), tB.L(l))
	if len(inter) > 0 {
		return makePath(tA, tB, inter, l, l)
	}

	return nil, 0
}

func intersect[V comparable](l1, l2 Level[V]) []V {
	var c []V

	l2.Each(func(i V, _ Vertex[V]) bool {
		if l1.Contain(i) {
			c = append(c, i)
		}
		return true
	})

	return c
}

func makePath[V comparable](tA, tB Tree[V], inter []V, l1, l2 int16) (set.Set[V], int16) {
	path := set.New[V]()
	for _, vertex := range inter {
		path.Add(vertex)
		getNextVertex(path, tA, vertex, l1)
		getNextVertex(path, tB, vertex, l2)
	}
	return path, l1 + l2
}

func getNextVertex[V comparable](path set.Set[V], t Tree[V], vertexID V, depth int16) {
	if p, ok := t.L(depth).Get(vertexID); p != nil && ok {
		p.Parents().Each(func(k V) bool {
			if t.L(depth - 1).Contain(k) {
				path.Add(k)
				getNextVertex(path, t, k, depth-1)
			}
			return true
		})
	}
}
