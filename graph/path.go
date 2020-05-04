package graph

import (
	"github.com/dizal/gocontainers/set"
	"github.com/dizal/gocontainers/tree"
)

// SearchShortestPath ...
func (g *Graph) SearchShortestPath(
	v1, v2 interface{},
	maxLength int16,
) (*set.Set, int16) {
	t1 := tree.New(nil)
	t1.Level(0).AddVertex(v1)
	t2 := tree.New(nil)
	t2.Level(0).AddVertex(v2)

	level := int16(0)
	maxDepth := maxLength/2 + (maxLength % 2)

	for ; level <= maxDepth; level++ {
		if t1.Level(level).Len() > 0 {
			t1.Level(level).Each(func(source interface{}, v *tree.Vertex) bool {
				g.Range(source, func(target interface{}) bool {
					t1.Level(level+1).AddEdge(source, target)
					return true
				})
				return true
			})
		} else {
			break
		}

		if t2.Level(level).Len() > 0 {

			t2.Level(level).Each(func(source interface{}, v *tree.Vertex) bool {
				g.Range(source, func(target interface{}) bool {
					t2.Level(level+1).AddEdge(source, target)
					return true
				})
				return true
			})
		} else {
			break
		}

		path, length := tree.SearchShortestPath(t1, t2, level)
		if path != nil {
			return path, length
		}
	}
	return nil, 0
}
