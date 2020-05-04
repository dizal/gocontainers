package syncset

import "sync"

// SyncSet ...
type SyncSet struct {
	mx  sync.RWMutex
	arr map[interface{}]struct{}
}

// New ...
func New() *SyncSet {
	return &SyncSet{
		arr: make(map[interface{}]struct{}),
	}
}

// Add ...
func (s *SyncSet) Add(value interface{}) bool {
	if !s.Contain(value) {
		s.mx.Lock()
		s.arr[value] = struct{}{}
		s.mx.Unlock()
		return true
	}
	return false
}

// Delete ...
func (s *SyncSet) Delete(value interface{}) bool {
	if s.Contain(value) {
		s.mx.Lock()
		delete(s.arr, value)
		s.mx.Unlock()
		return true
	}
	return false
}

// Contain ...
func (s *SyncSet) Contain(value interface{}) bool {
	s.mx.RLock()
	defer s.mx.RUnlock()
	_, ok := s.arr[value]
	return ok
}

// Len ...
func (s *SyncSet) Len() int {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return len(s.arr)
}

// Each ...
func (s *SyncSet) Each(f func(value interface{}) bool) {
	s.mx.RLock()
	for k := range s.arr {
		if !f(k) {
			break
		}
	}
	s.mx.RUnlock()
}

// ToSlice ...
func (s *SyncSet) ToSlice() []interface{} {
	arr := make([]interface{}, 0, s.Len())
	s.Each(func(v interface{}) bool {
		arr = append(arr, v)
		return true
	})
	return arr
}

// IsEmpty ...
func (s *SyncSet) IsEmpty() bool {
	return s.Len() == 0
}

// Erase ...
func (s *SyncSet) Erase() {
	s.mx.Lock()
	s.arr = map[interface{}]struct{}{}
	s.mx.Unlock()
}
