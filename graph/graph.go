package graph

import (
	"fmt"
	"strings"
	"sync"

	"github.com/dizal/gocontainers/set"
)

// Kind - Directed or undirected.
type Kind byte

const (
	// Undirected type of Graph
	Undirected Kind = iota
	// Directed type of Graph
	Directed
)

// Graph ...
type Graph struct {
	mx    sync.RWMutex
	s     map[interface{}]*set.Set
	kind  Kind
	edges uint32
}

// New ...
func New(kind Kind) *Graph {
	return &Graph{
		s:    make(map[interface{}]*set.Set),
		kind: kind,
	}
}

// AddEdge ...
func (g *Graph) AddEdge(source, target interface{}) *Graph {
	g.mx.Lock()
	if _, ok := g.s[source]; !ok {
		g.s[source] = set.New()
	}
	if _, ok := g.s[target]; !ok {
		g.s[target] = set.New()
	}
	if g.s[source].Add(target) {
		g.edges++
	}

	if g.kind == Undirected {
		if g.s[target].Add(source) {
			g.edges++
		}
	}
	g.mx.Unlock()
	return g
}

// GetTarget ...
func (g *Graph) GetTarget(source interface{}) (*set.Set, bool) {
	g.mx.RLock()
	t, ok := g.s[source]
	g.mx.RUnlock()
	return t, ok
}

// Range ...
func (g *Graph) Range(source interface{}, f func(target interface{}) bool) {
	if v, ok := g.GetTarget(source); ok {
		v.Each(func(target interface{}) bool {
			return f(target)
		})
	}
}

// VertexesCount ..
func (g *Graph) VertexesCount() uint32 {
	g.mx.RLock()
	defer g.mx.RUnlock()
	return uint32(len(g.s))
}

// EdgesCount ...
func (g *Graph) EdgesCount() uint32 {
	g.mx.RLock()
	defer g.mx.RUnlock()
	return g.edges
}

// Erase ...
func (g *Graph) Erase() {
	g.mx.Lock()
	g.s = make(map[interface{}]*set.Set)
	g.edges = 0
	g.mx.Unlock()
}

func (g *Graph) String() string {
	var buffer strings.Builder
	buffer.WriteString("\nGraph[\n")
	for k, v := range g.s {
		buffer.WriteString(fmt.Sprintf("\t{%v => %s}\n", k, v))
	}
	buffer.WriteString("]")
	return buffer.String()
}
