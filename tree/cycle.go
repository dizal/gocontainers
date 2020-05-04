package tree

import (
	"fmt"

	"github.com/dizal/gocontainers/set"
)

// CyclicData ...
type CyclicData struct {
	Vertexes map[interface{}]CyclicVertexData
	Count    uint32
}

// CyclicVertexData ...
type CyclicVertexData struct {
	Level  int16
	Degree int
}

func (c CyclicVertexData) String() string {
	return fmt.Sprintf("Vdata<L:%v, D:%v>", c.Level, c.Degree)
}

// SearchCyclicVertexes ...
func (t *Tree) SearchCyclicVertexes(
	level int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(vertexes map[interface{}]int16) bool,
) *CyclicData {
	c := CyclicData{make(map[interface{}]CyclicVertexData), 0}

	for ; level > 0; level-- {
		saw := set.New()
		l := t.Level(level)

		l.Each(func(vID interface{}, vData *Vertex) bool {
			var initVertexes []interface{}

			if vData.Siblings.Len() > 0 {
				sibs := set.New()
				l.GetRecursionSibling(sibs, vID)

				sibs.Each(func(v interface{}) bool {
					if saw.Contain(v) {
						sibs.Delete(v)
					}
					return true
				})

				initVertexes = sibs.ToSlice()
			} else if vData.Parents.Len() > 1 {
				if !saw.Contain(vID) {
					initVertexes = []interface{}{vID}
				}
			}
			for _, v := range initVertexes {
				saw.Add(v)
			}

			if len(initVertexes) > 0 {
				t.makeCycle(&c, level, initVertexes, onlyMainCycle, additionalCycleCheck)
			}

			return true
		})
	}

	if calcDegree {
		for v, d := range c.Vertexes {
			if vv, ok := t.Level(d.Level).Get(v); ok {
				vv.Mark()
			}
		}

		for v, d := range c.Vertexes {
			d.Degree = t.Level(d.Level).GetVertexDegreeMarked(v)
			c.Vertexes[v] = d
		}
	}

	return &c
}

func (t *Tree) makeCycle(
	c *CyclicData,
	initLevel int16,
	initVertexes []interface{},
	onlyMain bool,
	additionalCheck func(vertexes map[interface{}]int16) bool,
) {
	var temp set.Set

	cycleVertexes := make(map[interface{}]int16)

	vertexesOnLevel := *set.New()

	for _, v := range initVertexes {
		cycleVertexes[v] = initLevel
		vertexesOnLevel.Add(v)
	}

	for level := initLevel; level > 0; level-- {
		temp = *set.New()

		l := t.Level(level)

		parentContain := func(parent interface{}) bool {
			if t.Level(level - 1).Contain(parent) {
				temp.Add(parent)
			}
			return true
		}

		vertexesOnLevel.Each(func(v interface{}) bool {
			if vertex, ok := l.Get(v); ok {

				vertex.Parents.Each(parentContain)

				vertex.Siblings.Each(func(s interface{}) bool {
					if sib, ok := l.Get(s); ok {
						sib.Parents.Each(parentContain)
					}

					return true
				})
			}

			return true
		})

		vertexesOnLevel = temp

		temp.Each(func(v interface{}) bool {
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
		c.Vertexes[v] = CyclicVertexData{
			Level: level,
		}
	}
	c.Count++
}
