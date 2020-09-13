package set

import (
	"fmt"
	"unsafe"
)

// Set ...
type Set struct {
	arr map[interface{}]struct{}
}

// New ...
func New() *Set {
	return &Set{
		arr: make(map[interface{}]struct{}),
	}
}

// Add ...
func (s *Set) Add(value interface{}) bool {
	if !s.Contain(value) {
		s.arr[value] = struct{}{}
		return true
	}
	return false
}

// Delete ...
func (s *Set) Delete(value interface{}) bool {
	if s.Contain(value) {
		delete(s.arr, value)
		return true
	}
	return false
}

// Contain ...
func (s *Set) Contain(value interface{}) bool {
	_, ok := s.arr[value]
	return ok
}

// Len ...
func (s *Set) Len() int {
	return len(s.arr)
}

// Each ...
func (s *Set) Each(f func(value interface{}) bool) {
	for v := range s.arr {
		if !f(v) {
			break
		}
	}
}

// ToSlice ...
func (s *Set) ToSlice() []interface{} {
	arr := make([]interface{}, 0, s.Len())
	for k := range s.arr {
		arr = append(arr, k)
	}
	return arr
}

// IsEmpty ...
func (s *Set) IsEmpty() bool {
	return s.Len() == 0
}

// Erase ...
func (s *Set) Erase() {
	s.arr = make(map[interface{}]struct{})
}

func (s *Set) String() string {
	if s.Len() == 0 {
		return "SET<>"
	}
	buf := []byte("SET<")
	for v := range s.arr {
		buf = append(buf, fmt.Sprintf("%v, ", v)...)
	}
	buf = append(buf[:len(buf)-2], ">"...)
	return *(*string)(unsafe.Pointer(&buf))
}
