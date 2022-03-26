package set

import (
	"fmt"
	"unsafe"
)

type Set[T comparable] interface {
	Add(value T) bool
	Delete(value T) bool
	Contain(value T) bool
	Len() int
	Each(f func(value T) bool)
	ToSlice() []T
	IsEmpty() bool
	Erase()
	String() string
}

type set[T comparable] struct {
	arr map[T]struct{}
}

// New creates a new Set.
func New[T comparable]() Set[T] {
	return &set[T]{
		arr: make(map[T]struct{}),
	}
}

// Add adds a new value to the set.
func (s *set[T]) Add(value T) bool {
	if !s.Contain(value) {
		s.arr[value] = struct{}{}
		return true
	}
	return false
}

// Delete removes a value from the set.
func (s *set[T]) Delete(value T) bool {
	if s.Contain(value) {
		delete(s.arr, value)
		return true
	}
	return false
}

// Contain checks that the value is present in the set.
func (s *set[T]) Contain(value T) bool {
	_, ok := s.arr[value]
	return ok
}

// Len returns the length of the set.
func (s *set[T]) Len() int {
	return len(s.arr)
}

// Each goes through all the values of the set in the callback function.
func (s *set[T]) Each(f func(value T) bool) {
	for v := range s.arr {
		if !f(v) {
			break
		}
	}
}

// ToSlice converts the set to a slice.
func (s *set[T]) ToSlice() []T {
	arr := make([]T, 0, s.Len())
	for k := range s.arr {
		arr = append(arr, k)
	}
	return arr
}

func (s *set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *set[T]) Erase() {
	s.arr = make(map[T]struct{})
}

func (s *set[T]) String() string {
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
