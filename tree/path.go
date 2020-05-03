package tree

import "github.com/dizal/gocontainers/set"

// SearchShortestPath ...
func SearchShortestPath(tA, tB *Tree, l int16) (*set.Set, int16) {
	lA := tA.Level(l).Len()
	lB := tB.Level(l).Len()
	lA2 := tA.Level(l - 1).Len()
	lB2 := tB.Level(l - 1).Len()

	var inter []uint32
	bIsShort := false

	// even length (A - x - (x - x) - x - B)
	if lA+lB2 > lB+lA2 {
		inter = intersect(tA.Level(l-1), tB.Level(l))
	} else {
		bIsShort = true
		inter = intersect(tA.Level(l), tB.Level(l-1))
	}

	if len(inter) > 0 {
		if bIsShort {
			return makePath(tA, tB, inter, l, l-1)
		}
		return makePath(tA, tB, inter, l-1, l)
	}
	// odd length (А - x - (x) - x - B)
	inter = intersect(tA.Level(l), tB.Level(l))
	if len(inter) > 0 {
		return makePath(tA, tB, inter, l, l)
	}

	return nil, 0
}

func makePath(tA, tB *Tree, inter []uint32, l1, l2 int16) (*set.Set, int16) {
	path := set.New()
	for _, vertex := range inter {
		path.Add(vertex)
		getNextVertex(path, tA, vertex, l1)
		getNextVertex(path, tB, vertex, l2)
	}
	return path, l1 + l2
}

func getNextVertex(path *set.Set, t *Tree, vertex uint32, d int16) {
	if p, ok := t.Level(d).Get(vertex); p != nil && ok {
		p.Parents.Each(func(k interface{}) bool {
			if t.Level(d - 1).Contain(k.(uint32)) {
				path.Add(k)
				getNextVertex(path, t, k.(uint32), d-1)
			}
			return true
		})
	}
}
