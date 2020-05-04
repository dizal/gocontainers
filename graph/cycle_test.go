package graph

import (
	"strings"
	"testing"

	"github.com/dizal/gocontainers/tree"
)

func TestGraphSearchCycle1(t *testing.T) {
	g := g1()

	cd := g.SearchCycle("v1", 3, true, true, nil)

	target := map[string]tree.CyclicVertexData{
		"v1": {Level: 0, Degree: 2},
		"v2": {Level: 1, Degree: 2},
		"v3": {Level: 2, Degree: 2},
		"v4": {Level: 1, Degree: 2},
	}
	testCycleResult(t, cd, target, 1)
}

func TestGraphSearchCycle2(t *testing.T) {
	g := g2()

	cd := g.SearchCycle("v1", 3, true, false, func(vertexes map[interface{}]int16) bool {
		// skip all cycles smaller than 3
		return len(vertexes) > 3
	})

	testCycleResult(t, cd, nil, 0)

	cd = g.SearchCycle("v1", 3, true, true, nil)

	target := map[string]tree.CyclicVertexData{
		"v1": {Level: 0, Degree: 2},
		"v2": {Level: 1, Degree: 2},
		"v3": {Level: 1, Degree: 2},
	}
	testCycleResult(t, cd, target, 1)
}

func TestGraphSearchCycle3(t *testing.T) {
	g := g3()

	cd := g.SearchCycle("v0", 3, true, false, nil)

	testCycleResult(t, cd, nil, 0)

	cd = g.SearchCycle("v0", 4, false, true, nil)

	target := map[string]tree.CyclicVertexData{
		"v1": {Level: 1, Degree: 3},
		"v2": {Level: 2, Degree: 2},
		"v3": {Level: 3, Degree: 3},
		"v4": {Level: 2, Degree: 2},
		"v5": {Level: 2, Degree: 2},
	}

	testCycleResult(t, cd, target, 1)
}

func TestGraphSearchCycle4(t *testing.T) {
	g := g4()

	cd, _ := g.searchCycle("v1", 3, true, true, nil)

	target := map[string]tree.CyclicVertexData{
		"v1": {Level: 0, Degree: 3},
		"v2": {Level: 1, Degree: 2},
		"v3": {Level: 1, Degree: 3},
		"v4": {Level: 1, Degree: 2},
	}
	testCycleResult(t, cd, target, 1)
}

func TestGraphSearchCycle5(t *testing.T) {
	g := g5()

	cd, _ := g.searchCycle("c_v1", 7, true, true, func(vertexes map[interface{}]int16) bool {
		// skip all cycles less then 3
		if len(vertexes) <= 3 {
			return false
		}

		cCount := 0
		for k := range vertexes {
			if strings.HasPrefix(k.(string), "c_") {
				cCount++
			}
		}
		// skip all cycles where number of c_ type vertexes less then 3
		return cCount >= 3
	})

	target := map[string]tree.CyclicVertexData{
		"c_v1":  {Level: 0, Degree: 2},
		"v_v2":  {Level: 1, Degree: 3},
		"c_v3":  {Level: 2, Degree: 3},
		"k_v4":  {Level: 3, Degree: 2},
		"c_v5":  {Level: 4, Degree: 2},
		"k_v6":  {Level: 5, Degree: 2},
		"c_v7":  {Level: 6, Degree: 2},
		"k_v8":  {Level: 7, Degree: 2},
		"k_v9":  {Level: 1, Degree: 4},
		"v_v10": {Level: 2, Degree: 2},
		"k_v11": {Level: 3, Degree: 2},
		"c_v12": {Level: 4, Degree: 2},
		"k_v13": {Level: 5, Degree: 2},
		"c_v14": {Level: 6, Degree: 2},
		// "c_v15": {Level: 3, Degree: 2},
	}
	testCycleResult(t, cd, target, 1)
}

func BenchmarkSearchCycle(b *testing.B) {
	b.ReportAllocs()
	g := g5()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.SearchCycle("c_v1", 7, true, true, nil)
	}
}

func testCycleResult(t *testing.T, cd *tree.CyclicData, target map[string]tree.CyclicVertexData, countCycles uint32) {
	if cd.Count != countCycles {
		t.Errorf("testResult: number of cycles is not equal to %v. Response: %v", countCycles, cd.Count)
	}

	if len(cd.Vertexes) != len(target) {
		t.Errorf("testResult: number of vertices in a cycle is not equal to %v. Response %v", len(target), len(cd.Vertexes))
	}

	for k, tData := range target {
		respData, ok := cd.Vertexes[k]

		if !ok {
			t.Errorf("testResult: vertex %v is not found in the cycle", k)
			continue
		}
		if tData.Degree != respData.Degree {
			t.Errorf("testResult: vertex %v. Target degree %v. Response %v", k, tData.Degree, respData.Degree)
		}
		if tData.Level != respData.Level {
			t.Errorf("testResult: vertex %v. Target level %v. Response %v", k, tData.Level, respData.Level)
		}
	}
}
