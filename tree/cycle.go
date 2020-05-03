package tree

import "github.com/dizal/gocontainers/set"

// CyclicData ...
type CyclicData struct {
	Vertexes map[uint32]int16
	Count    uint32
}

// SearchCyclicVertexes ...
func (t *Tree) SearchCyclicVertexes(
	level int16,
	onlyMainCycle, calcDegree bool,
	additionalCycleCheck func(vertexes map[uint32]int16) bool,
) *CyclicData {
	c := CyclicData{make(map[uint32]int16), 0}

	for ; level > 1; level-- {
		saw := set.New()
		l := t.Level(level)

		l.Each(func(vID uint32, vData *Vertex) bool {
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
			} else if vData.Parents.Len() > 0 && !saw.Contain(vID) {
				initVertexes = []interface{}{vID}
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
		for v, l := range c.Vertexes {
			if vv, ok := t.Level(l).Get(v); ok {
				vv.Mark()
			}
		}

		for node, level := range c.Vertexes {
			c.Vertexes[node] = int16(t.Level(level).GetVertexDegreeMarked(node))
		}
	}

	return &c
}

func (t *Tree) makeCycle(
	c *CyclicData,
	initLevel int16,
	initVertexes []interface{},
	onlyMain bool,
	additionalCheck func(vertexes map[uint32]int16) bool,
) {
	var temp set.Set

	cycleVertexes := make(map[uint32]int16)

	// узлы, которые находятся на текущем уровне
	vertexesOnLevel := *set.New()

	for _, v := range initVertexes {
		vv := v.(uint32)
		cycleVertexes[vv] = initLevel
		vertexesOnLevel.Add(vv)
	}

	for level := initLevel; level > 0; level-- {
		temp = *set.New()

		l := t.Level(level)

		parentContain := func(parent interface{}) bool {
			if t.Level(level - 1).Contain(parent.(uint32)) {
				temp.Add(parent)
			}
			return true
		}

		vertexesOnLevel.Each(func(v interface{}) bool {
			if vertex, ok := l.Get(v.(uint32)); ok {

				vertex.Parents.Each(parentContain)

				vertex.Siblings.Each(func(s interface{}) bool {
					if sib, ok := l.Get(s.(uint32)); ok {
						sib.Parents.Each(parentContain)
					}

					return true
				})
			}

			return true
		})

		vertexesOnLevel = temp

		temp.Each(func(v interface{}) bool {
			cycleVertexes[v.(uint32)] = level - 1
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
		c.Vertexes[v] = level
	}
	c.Count++
}