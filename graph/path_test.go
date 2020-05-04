package graph

import (
	"testing"

	"github.com/dizal/gocontainers/set"
)

func TestGraphSearchPath1(t *testing.T) {
	path, length := g1().SearchShortestPath("v1", "v3", 2)
	testPath(t, path, length, 2)
}

func TestGraphSearchPath2(t *testing.T) {
	path, length := g2().SearchShortestPath("v1", "v2", 2)
	testPath(t, path, length, 1)
}

func TestGraphSearchPath3(t *testing.T) {
	path, length := g3().SearchShortestPath("v0", "v7", 4)
	testPath(t, path, length, 4)
}

func TestGraphSearchPath4(t *testing.T) {
	path, length := g4().SearchShortestPath("v2", "v4", 2)
	testPath(t, path, length, 2)
}

func TestGraphSearchPath5(t *testing.T) {
	g := g5()
	path, length := g.SearchShortestPath("c_v1", "k_v8", 7)
	testPath(t, path, length, 7)

	path, length = g.SearchShortestPath("c_v1", "k_v18", 4)
	testPath(t, path, length, 4)

	path, length = g.SearchShortestPath("c_v5", "k_v18", 6)
	testPath(t, path, length, 6)
}

func BenchmarkSearchPath(b *testing.B) {
	b.ReportAllocs()
	g := g5()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = g.SearchShortestPath("c_v1", "k_v8", 7)
	}
}

func testPath(t *testing.T, path *set.Set, length, targetLen int16) {
	if path == nil {
		t.Error("testPath: path is nil")
	}

	if length != targetLen {
		t.Errorf("testPath: uncorrenct path length. Target %v. Response: %v", targetLen, length)
	}
}
