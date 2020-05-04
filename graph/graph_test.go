package graph

import "testing"

func TestGraphNew(t *testing.T) {
	g := New(Undirected)

	if g.VertexesCount() != 0 {
		t.Error("new: graph is not empry")
	}
}

func TestGraphAddEdge(t *testing.T) {
	g := New(Undirected)
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v1", "v3")

	if g.VertexesCount() != 3 {
		t.Error("add: Undirected Graph Vertexes is not 3")
	}

	if g.EdgesCount() != 4 {
		t.Error("add: Undirected Graph Edges is not 4")
	}

	g = New(Directed)
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v1", "v3")

	if g.VertexesCount() != 3 {
		t.Error("add: Directed Graph Vertexes is not 3")
	}

	if g.EdgesCount() != 2 {
		t.Error("add: Directed Graph Edges is not 2")
	}
}

func TestGraphGetTarget(t *testing.T) {
	g := New(Undirected)
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v2", "v3")

	if g.VertexesCount() != 3 {
		t.Error("get: Undirected Graph Vertexes is not 3")
	}

	if g.EdgesCount() != 6 {
		t.Error("get: Undirected Graph Edges is not 4")
	}

	s, ok := g.GetTarget("v3")
	if !ok {
		t.Error("get: Undirected Graph Target is empty")
	}
	if s.Len() != 2 {
		t.Error("get: Undirected Graph Target length is not 2")
	}

	g = New(Directed)
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v2", "v3")

	if g.VertexesCount() != 3 {
		t.Error("get: Directed Graph Vertexes is not 3")
	}

	if g.EdgesCount() != 3 {
		t.Error("get: Directed Graph Edges is not 2")
	}

	s, ok = g.GetTarget("v3")
	if !ok {
		t.Error("get: Undirected Graph Target is empty")
	}
	if s.Len() != 0 {
		t.Error("get: Undirected Graph Target length is not 2")
	}
}

func TestGraphErase(t *testing.T) {
	g := New(Undirected)

	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v2", "v3")

	g.Erase()

	if g.VertexesCount() != 0 {
		t.Error("erase: graph is not empry")
	}
	if g.EdgesCount() != 0 {
		t.Error("erase: graph is not empry")
	}
}

func TestGraphRange(t *testing.T) {
	g := New(Undirected)

	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v2", "v3")

	g.Range("v3", func(target interface{}) bool {
		s, ok := target.(string)
		if !ok {
			t.Error("range: cannot convert target value to string")
		}
		if !(s == "v1" || s == "v2") {
			t.Error("range: undefined value")
		}

		return true
	})

}
