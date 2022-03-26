package tree

import (
	"fmt"

	"github.com/dizal/gocontainers/set"
)

type SearchCyclicResponse[V comparable] struct {
	Vertexes map[V]CyclicVertexData
	Count    uint32
}

type CyclicVertexData struct {
	Level  int16
	Degree int
}

func (c CyclicVertexData) String() string {
	return fmt.Sprintf("Vdata<L:%v, D:%v>", c.Level, c.Degree)
}

// SearchCyclicVertexes ...
func SearchCyclicVertexes[V comparable](
	t Tree[V],
	level int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(vertexes map[V]int16) bool,
) *SearchCyclicResponse[V] {

	res := &SearchCyclicResponse[V]{make(map[V]CyclicVertexData), 0}

	for ; level > 0; level-- {
		saw := set.New[V]()
		l := t.L(level)

		l.Each(func(vID V, vData Vertex[V]) bool {
			var initVertexes []V

			if vData.Siblings().Len() > 0 {
				sibs := set.New[V]()
				l.GetRecursionSibling(sibs, vID)

				sibs.Each(func(v V) bool {
					if saw.Contain(v) {
						sibs.Delete(v)
					}
					return true
				})

				initVertexes = sibs.ToSlice()
			} else if vData.Parents().Len() > 1 {
				if !saw.Contain(vID) {
					initVertexes = []V{vID}
				}
			}
			for _, v := range initVertexes {
				saw.Add(v)
			}

			if len(initVertexes) > 0 {
				makeCycle(t, res, level, initVertexes, onlyMainCycle, additionalCycleCheck)
			}

			return true
		})
	}

	if calcDegree {
		for v, d := range res.Vertexes {
			if vv, ok := t.L(d.Level).Get(v); ok {
				vv.Mark()
			}
		}

		for v, d := range res.Vertexes {
			d.Degree = t.L(d.Level).GetVertexDegreeMarked(v)
			res.Vertexes[v] = d
		}
	}

	return res
}

func makeCycle[V comparable](
	t Tree[V],
	res *SearchCyclicResponse[V],
	initLevel int16,
	initVertexes []V,
	onlyMain bool,
	additionalCheck func(vertexes map[V]int16) bool,
) {
	var temp set.Set[V]

	cycleVertexes := make(map[V]int16)

	vertexesOnLevel := set.New[V]()

	for _, v := range initVertexes {
		cycleVertexes[v] = initLevel
		vertexesOnLevel.Add(v)
	}

	for level := initLevel; level > 0; level-- {
		temp = set.New[V]()

		l := t.L(level)

		parentContain := func(parent V) bool {
			if t.L(level - 1).Contain(parent) {
				temp.Add(parent)
			}
			return true
		}

		vertexesOnLevel.Each(func(v V) bool {
			if vertex, ok := l.Get(v); ok {

				vertex.Parents().Each(parentContain)

				vertex.Siblings().Each(func(s V) bool {
					if sib, ok := l.Get(s); ok {
						sib.Parents().Each(parentContain)
					}

					return true
				})
			}

			return true
		})

		vertexesOnLevel = temp

		temp.Each(func(v V) bool {
			cycleVertexes[v] = level - 1
			return true
		})

		if temp.Len() < 2 {
			if onlyMain && level > 1 {
				return
			}
			break
		}
	}

	if additionalCheck != nil {
		if !additionalCheck(cycleVertexes) {
			return
		}
	}

	for v, level := range cycleVertexes {
		res.Vertexes[v] = CyclicVertexData{
			Level: level,
		}
	}
	res.Count++
}
