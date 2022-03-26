package graph

import (
	"github.com/dizal/gocontainers/set"
	"github.com/dizal/gocontainers/tree"
)

// SearchShortestPath looking for the shortest path between node A and node B.
// maxLength limits the search to a given length.
//
// A tree is built from each of the two nodes. Expanding continues until there
// are identical nodes in two different trees on the same or neighboring level.
func (g *Graph[T]) SearchShortestPath(nodeA, nodeB T, maxLength int16) (set.Set[T], int16) {
	t1 := tree.New[T]()
	t1.L(0).AddVertex(nodeA)
	t2 := tree.New[T]()
	t2.L(0).AddVertex(nodeB)

	// the maximum depth for each tree is half of the maximum length
	maxDepth := maxLength/2 + (maxLength % 2)

	for level := int16(0); level <= maxDepth; level++ {
		if t1.L(level).Len() == 0 || t2.L(level).Len() == 0 {
			// no more nodes to expand
			break
		}

		t1.L(level).Each(func(sourceNode T, v tree.Vertex[T]) bool {
			g.Range(sourceNode, func(targetNode T) bool {
				t1.L(level+1).AddEdge(sourceNode, targetNode)
				return true
			})
			return true
		})

		t2.L(level).Each(func(sourceNode T, v tree.Vertex[T]) bool {
			g.Range(sourceNode, func(targetNode T) bool {
				t2.L(level+1).AddEdge(sourceNode, targetNode)
				return true
			})
			return true
		})

		path, length := tree.SearchShortestPath(t1, t2, level)
		if path != nil {
			return path, length
		}
	}
	return nil, 0
}
