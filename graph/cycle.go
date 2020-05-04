package graph

import (
	"github.com/dizal/gocontainers/tree"
)

// SearchCycle ...
func (g *Graph) SearchCycle(
	vertex interface{},
	maxDepth int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(vertexes map[interface{}]int16) bool,
) *tree.CyclicData {
	resp, _ := g.searchCycle(vertex, maxDepth, onlyMainCycle, calcDegree, additionalCycleCheck)
	return resp
}

func (g *Graph) searchCycle(
	vertex interface{},
	maxDepth int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(v map[interface{}]int16) bool,
) (*tree.CyclicData, *tree.Tree) {
	t := tree.New(nil)
	t.Level(0).AddVertex(vertex)

	level := int16(0)

	for ; level < maxDepth; level++ {
		if t.Level(level).Len() > 0 {

			t.Level(level).Each(func(source interface{}, v *tree.Vertex) bool {
				g.Range(source, func(target interface{}) bool {
					t.Level(level+1).AddEdge(source, target)
					return true
				})
				return true
			})
		} else {
			break
		}
	}

	return t.SearchCyclicVertexes(level, onlyMainCycle, calcDegree, additionalCycleCheck), t
}
