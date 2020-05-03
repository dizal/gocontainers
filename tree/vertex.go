package tree

import "github.com/dizal/gocontainers/set"

// Vertex ...
type Vertex struct {
	Parents  *set.Set
	Siblings *set.Set
	Children *set.Set
	Marked   bool
}

// AddParent ...
func (v *Vertex) AddParent(parent uint32) bool {
	if !v.Parents.Contain(parent) {
		v.Parents.Add(parent)
		return true
	}
	return false
}

// AddChild ...
func (v *Vertex) AddChild(child uint32) bool {
	if !v.Children.Contain(child) {
		v.Children.Add(child)
		return true
	}
	return false
}

// AddSibling ...
func (v *Vertex) AddSibling(sib uint32) bool {
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
