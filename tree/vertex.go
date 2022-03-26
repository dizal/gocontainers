package tree

import (
	"fmt"
	"strings"

	"github.com/dizal/gocontainers/set"
)

//  Vertex implements a tree node.
//	L-1      parent1      parent2
//	               \      |
//	L0  sibling1 -- current -- sibling2
//	               /      \
//	L1         child1    child2
type Vertex[T comparable] interface {
	AddParent(parent T) bool
	AddChild(child T) bool
	AddSibling(sibling T) bool
	Parents() set.Set[T]
	Siblings() set.Set[T]
	Children() set.Set[T]
	Degree() int
	Mark()
	IsMarked() bool
	String() string
}

type vertex[T comparable] struct {
	parents  set.Set[T]
	siblings set.Set[T]
	children set.Set[T]
	marked   bool
}

// NewVertex ...
func NewVertex[T comparable]() Vertex[T] {
	return &vertex[T]{
		parents:  set.New[T](),
		siblings: set.New[T](),
		children: set.New[T](),
	}
}

func (v *vertex[T]) Parents() set.Set[T] {
	return v.parents
}

func (v *vertex[T]) Siblings() set.Set[T] {
	return v.siblings
}

func (v *vertex[T]) Children() set.Set[T] {
	return v.children
}

// AddParent ...
func (v *vertex[T]) AddParent(parent T) bool {
	return v.parents.Add(parent)
}

// AddChild ...
func (v *vertex[T]) AddChild(child T) bool {
	return v.children.Add(child)
}

// AddSibling ...
func (v *vertex[T]) AddSibling(sib T) bool {
	return v.siblings.Add(sib)
}

// Degree ...
func (v *vertex[T]) Degree() int {
	d := 0
	if v.parents != nil {
		d += v.parents.Len()
	}
	if v.children != nil {
		d += v.children.Len()
	}
	if v.siblings != nil {
		d += v.siblings.Len()
	}
	return d
}

// Mark ...
func (v *vertex[T]) Mark() {
	v.marked = true
}

func (v *vertex[T]) IsMarked() bool {
	return v.marked
}

func (v *vertex[T]) String() string {
	var buffer strings.Builder
	buffer.WriteString("V[")
	if v.parents != nil {
		buffer.WriteString(fmt.Sprintf("P:%v, ", v.parents))
	}
	if v.children != nil {
		buffer.WriteString(fmt.Sprintf("C:%v, ", v.children))
	}
	if v.siblings != nil {
		buffer.WriteString(fmt.Sprintf("S:%v, ", v.siblings))
	}
	buffer.WriteString(fmt.Sprintf("M:%v", v.marked))
	buffer.WriteString("]")
	return buffer.String()
}
