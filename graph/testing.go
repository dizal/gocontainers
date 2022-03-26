package graph

func g1() *Graph[string] {
	g := New[string](Undirected)

	//   v1
	//  /  \
	// v2  v4
	//  \  /
	//   v3
	g.AddEdge("v1", "v2")
	g.AddEdge("v2", "v3")
	g.AddEdge("v3", "v4")
	g.AddEdge("v4", "v1")
	return g
}

func g2() *Graph[string] {
	g := New[string](Undirected)
	//   v1
	//  /  \
	// v2--v3
	g.AddEdge("v1", "v2")
	g.AddEdge("v2", "v3")
	g.AddEdge("v3", "v1")
	return g
}

func g3() *Graph[string] {
	g := New[string](Undirected)

	//    v0
	//    |
	//    v1--v6
	//  / | \
	// v2 v4 v5--v8
	//  \ | /
	//   v3--v9
	//    |
	//    v7
	g.AddEdge("v0", "v1")
	g.AddEdge("v1", "v2")
	g.AddEdge("v2", "v3")
	g.AddEdge("v3", "v4")
	g.AddEdge("v4", "v1")
	g.AddEdge("v1", "v5")
	g.AddEdge("v5", "v3")
	g.AddEdge("v1", "v6")
	g.AddEdge("v3", "v7")
	g.AddEdge("v5", "v8")
	g.AddEdge("v3", "v9")
	return g
}

func g4() *Graph[string] {
	g := New[string](Undirected)

	//    v1
	//   / | \
	//  /  |  \
	// v2--v3--v4
	g.AddEdge("v1", "v2")
	g.AddEdge("v1", "v3")
	g.AddEdge("v1", "v4")
	g.AddEdge("v2", "v3")
	g.AddEdge("v3", "v4")
	return g
}

func g5() *Graph[string] {
	g := New[string](Undirected)

	//   c_v1                           L0
	//  /    \
	// v_v2--k_v9                       L1
	// |    / |
	// c_v3  v_v10----------            L2
	// |      |     \        \
	// k_v4  k_v11--c_v15   c_v16       L3
	// |      |             /   \
	// c_v5  c_v12      c_v17---k_v18   L4
	// |      |
	// k_v6  k_v13                      L5
	// |      |
	// c_v7  c_v14                      L6
	// |     /
	// k_v8                             L7

	g.AddEdge("c_v1", "v_v2").AddEdge("v_v2", "c_v3")
	g.AddEdge("c_v3", "k_v4").AddEdge("k_v4", "c_v5")
	g.AddEdge("c_v5", "k_v6").AddEdge("k_v6", "c_v7")
	g.AddEdge("c_v7", "k_v8").AddEdge("c_v1", "k_v9")
	g.AddEdge("k_v9", "v_v10").AddEdge("v_v10", "k_v11")
	g.AddEdge("k_v11", "c_v12").AddEdge("c_v12", "k_v13")
	g.AddEdge("k_v13", "c_v14").AddEdge("c_v14", "k_v8")
	g.AddEdge("v_v2", "k_v9").AddEdge("c_v3", "k_v9")
	g.AddEdge("c_v15", "v_v10").AddEdge("c_v15", "k_v11")
	g.AddEdge("c_v16", "v_v10").AddEdge("c_v16", "c_v17")
	g.AddEdge("c_v16", "k_v18").AddEdge("c_v17", "k_v18")

	return g
}
