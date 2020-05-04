package tree

import "github.com/dizal/gocontainers/set"

// SearchShortestPath ...
func SearchShortestPath(tA, tB *Tree, l int16) (*set.Set, int16) {
	lA := tA.Level(l).Len()
	lB := tB.Level(l).Len()
	lA2 := tA.Level(l - 1).Len()
	lB2 := tB.Level(l - 1).Len()

	var inter []interface{}
	bIsShort := lA+lB2 > lB+lA2

	// even length (A - x - (x - x) - x - B)
	if bIsShort {
		inter = intersect(tA.Level(l), tB.Level(l-1))
	} else {
		inter = intersect(tA.Level(l-1), tB.Level(l))
	}

	if len(inter) > 0 {
		if bIsShort {
			return makePath(tA, tB, inter, l, l-1)
		}
		return makePath(tA, tB, inter, l-1, l)
	}
	// odd length (A - x - (x) - x - B)
	inter = intersect(tA.Level(l), tB.Level(l))
	if len(inter) > 0 {
		return makePath(tA, tB, inter, l, l)
	}

	return nil, 0
}

func makePath(tA, tB *Tree, inter []interface{}, l1, l2 int16) (*set.Set, int16) {
	path := set.New()
	for _, vertex := range inter {
		path.Add(vertex)
		getNextVertex(path, tA, vertex, l1)
		getNextVertex(path, tB, vertex, l2)
	}
	return path, l1 + l2
}

func getNextVertex(path *set.Set, t *Tree, vertex interface{}, d int16) {
	if p, ok := t.Level(d).Get(vertex); p != nil && ok {
		p.Parents.Each(func(k interface{}) bool {
			if t.Level(d - 1).Contain(k) {
				path.Add(k)
				getNextVertex(path, t, k, d-1)
			}
			return true
		})
	}
}
