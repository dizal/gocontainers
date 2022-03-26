package graph

import (
	"github.com/dizal/gocontainers/tree"
)

// SearchCycle ...
func (g *Graph[T]) SearchCycle(
	vertex T,
	maxDepth int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(vertexes map[T]int16) bool,
) *tree.SearchCyclicResponse[T] {

	resp, _ := g.searchCycle(vertex, maxDepth, onlyMainCycle, calcDegree, additionalCycleCheck)
	return resp
}

func (g *Graph[T]) searchCycle(
	vertex T,
	maxDepth int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(v map[T]int16) bool,
) (*tree.SearchCyclicResponse[T], tree.Tree[T]) {

	t := tree.New[T]()
	t.L(0).AddVertex(vertex)

	level := int16(0)

	for ; level < maxDepth; level++ {
		if t.L(level).Len() > 0 {

			t.L(level).Each(func(source T, v tree.Vertex[T]) bool {
				g.Range(source, func(target T) bool {
					t.L(level+1).AddEdge(source, target)
					return true
				})
				return true
			})
		} else {
			break
		}
	}

	return tree.SearchCyclicVertexes(t, level, onlyMainCycle, calcDegree, additionalCycleCheck), t
}
