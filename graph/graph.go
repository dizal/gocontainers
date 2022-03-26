package graph

import (
	"fmt"
	"strings"
	"sync"

	"github.com/dizal/gocontainers/set"
)

// Kind of Graph. Directed or undirected.
type Kind byte

const (
	// Undirected type of Graph
	Undirected Kind = iota
	// Directed type of Graph
	Directed
)

type Graph[T comparable] struct {
	// several target nodes can be attached to each source node
	s map[T]set.Set[T]
	// kind of Graph
	kind Kind
	// the number of edges in the graph
	edges int

	mx sync.RWMutex
}

// New creates new Graph.
func New[T comparable](kind Kind) *Graph[T] {
	return &Graph[T]{
		s:    make(map[T]set.Set[T]),
		kind: kind,
	}
}

// AddEdge adds a new edge to the graph.
func (g *Graph[T]) AddEdge(sourceNode, targetNode T) *Graph[T] {
	g.mx.Lock()
	defer g.mx.Unlock()

	if _, ok := g.s[sourceNode]; !ok {
		g.s[sourceNode] = set.New[T]()
	}
	if _, ok := g.s[targetNode]; !ok {
		g.s[targetNode] = set.New[T]()
	}
	if g.s[sourceNode].Add(targetNode) {
		g.edges++
	}

	if g.kind == Undirected {
		if g.s[targetNode].Add(sourceNode) {
			g.edges++
		}
	}

	return g
}

// GetTargetNodes returns the set of nodes to which the source node is connected.
func (g *Graph[T]) GetTargetNodes(sourceNode T) (set.Set[T], bool) {
	g.mx.RLock()
	defer g.mx.RUnlock()
	t, ok := g.s[sourceNode]
	return t, ok
}

// Range goes through all target nodes of the source node
func (g *Graph[T]) Range(sourceNode T, f func(targetNode T) bool) {
	if v, ok := g.GetTargetNodes(sourceNode); ok {
		v.Each(func(target T) bool {
			return f(target)
		})
	}
}

// VertexesCount returns the number of nodes in the graph.
func (g *Graph[T]) VertexesCount() int {
	g.mx.RLock()
	defer g.mx.RUnlock()
	return len(g.s)
}

// EdgesCount returns the number of edges in the graph.
func (g *Graph[T]) EdgesCount() int {
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.edges
}

// Erase clears the graph of nodes and edges.
func (g *Graph[T]) Erase() {
	g.mx.Lock()
	defer g.mx.Unlock()
	g.s = make(map[T]set.Set[T])
	g.edges = 0
}

func (g *Graph[T]) String() string {
	var buffer strings.Builder
	buffer.WriteString("\nGraph[\n")
	for k, v := range g.s {
		buffer.WriteString(fmt.Sprintf("\t{%v => %s}\n", k, v))
	}
	buffer.WriteString("]")
	return buffer.String()
}
