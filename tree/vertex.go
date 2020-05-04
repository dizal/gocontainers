package tree

import (
	"fmt"
	"strings"

	"github.com/dizal/gocontainers/set"
)

// Vertex ...
type Vertex struct {
	Parents  *set.Set
	Siblings *set.Set
	Children *set.Set
	Marked   bool
}

// NewVertex ...
func NewVertex() *Vertex {
	return &Vertex{
		Parents:  set.New(),
		Siblings: set.New(),
		Children: set.New(),
	}
}

// AddParent ...
func (v *Vertex) AddParent(parent interface{}) bool {
	if !v.Parents.Contain(parent) {
		v.Parents.Add(parent)
		return true
	}
	return false
}

// AddChild ...
func (v *Vertex) AddChild(child interface{}) bool {
	if !v.Children.Contain(child) {
		v.Children.Add(child)
		return true
	}
	return false
}

// AddSibling ...
func (v *Vertex) AddSibling(sib interface{}) bool {
	if !v.Siblings.Contain(sib) {
		v.Siblings.Add(sib)
		return true
	}
	return false
}

// Degree ...
func (v *Vertex) Degree() int {
	d := 0
	if v.Parents != nil {
		d += v.Parents.Len()
	}
	if v.Children != nil {
		d += v.Children.Len()
	}
	if v.Siblings != nil {
		d += v.Siblings.Len()
	}
	return d
}

// Mark ...
func (v *Vertex) Mark() {
	v.Marked = true
}

func (v *Vertex) String() string {
	var buffer strings.Builder
	buffer.WriteString("V[")
	buffer.WriteString(fmt.Sprintf("P:%v, ", v.Parents))
	buffer.WriteString(fmt.Sprintf("C:%v, ", v.Children))
	buffer.WriteString(fmt.Sprintf("S:%v, ", v.Siblings))
	buffer.WriteString(fmt.Sprintf("M:%v", v.Marked))
	buffer.WriteString("]")
	return buffer.String()
}
