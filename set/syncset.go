package set

import "sync"

type syncSet[V comparable] struct {
	set[V]
	mx sync.RWMutex
}

// New ...
func NewSync[V comparable]() Set[V] {
	return &syncSet[V]{
		set: set[V]{
			arr: make(map[V]struct{}),
		},
	}
}

// Add ...
func (s *syncSet[V]) Add(value V) bool {
	if !s.Contain(value) {
		s.mx.Lock()
		s.arr[value] = struct{}{}
		s.mx.Unlock()
		return true
	}
	return false
}

// Delete ...
func (s *syncSet[V]) Delete(value V) bool {
	if s.Contain(value) {
		s.mx.Lock()
		delete(s.arr, value)
		s.mx.Unlock()
		return true
	}
	return false
}

// Contain ...
func (s *syncSet[V]) Contain(value V) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.set.Contain(value)
}

// Len ...
func (s *syncSet[V]) Len() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.set.Len()
}

// Each ...
func (s *syncSet[V]) Each(f func(value V) bool) {
	s.mx.RLock()
	s.set.Each(f)
	s.mx.RUnlock()
}

// ToSlice ...
func (s *syncSet[V]) ToSlice() []V {
	arr := make([]V, 0, s.Len())
	s.Each(func(v V) bool {
		arr = append(arr, v)
		return true
	})
	return arr
}

// Erase ...
func (s *syncSet[V]) Erase() {
	s.mx.Lock()
	s.set.Erase()
	s.mx.Unlock()
}
