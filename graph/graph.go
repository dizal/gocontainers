package graph

import "sync"

// Type - Directed or undirected.
type Type byte

const (
	// Undirected type of Graph
	Undirected Type = iota
	// Directed type of Graph
	Directed
)

// Target ...
type Target map[interface{}]struct{}

// Graph ...
type Graph struct {
	m    sync.RWMutex
	s    map[interface{}]Target
	kind Type
}

// New ...
func New(kind Type) *Graph {
	return &Graph{
		s:    make(map[interface{}]Target),
		kind: kind,
	}
}

// AddEdge ...
func (g *Graph) AddEdge(source, target interface{}) {
	g.m.Lock()
	if _, ok := g.s[source]; !ok {
		g.s[source] = make(Target)
	}
	g.s[source][target] = struct{}{}

	if g.kind == Undirected {
		if _, ok := g.s[target]; !ok {
			g.s[target] = make(Target)
		}
		g.s[target][source] = struct{}{}
	}
	g.m.Unlock()
}

// GetTarget ...
func (g *Graph) GetTarget(source interface{}) (Target, bool) {
	g.m.RLock()
	defer g.m.RUnlock()
	t, ok := g.s[source]
	return t, ok
}

// Range ...
func (g *Graph) Range(source interface{}, f func(target interface{}) bool) {
	if v, ok := g.GetTarget(source); ok {
		for target := range v {
			if !f(target) {
				break
			}
		}
	}
}

// Len ..
func (g *Graph) Len() int {
	return len(g.s)
}

// Erase ...
func (g *Graph) Erase() {
	g.m.Lock()
	g.s = make(map[interface{}]Target)
	g.m.Unlock()
}
