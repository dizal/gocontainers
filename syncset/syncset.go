package syncset

import "sync"

// SyncSet ...
type SyncSet struct {
	m    sync.RWMutex
	arr  sync.Map
	size int
}

// New ...
func New() *SyncSet {
	return &SyncSet{}
}

// Add ...
func (s *SyncSet) Add(v ...interface{}) {
	for _, vv := range v {
		if _, ok := s.arr.Load(vv); !ok {
			s.arr.Store(vv, struct{}{})
			s.m.Lock()
			s.size++
			s.m.Unlock()
		}
	}
}

// Delete ...
func (s *SyncSet) Delete(v ...interface{}) {
	for _, vv := range v {
		if _, ok := s.arr.Load(vv); ok {
			s.arr.Delete(vv)
			s.m.Lock()
			if s.size > 0 {
				s.size--
			}
			s.m.Unlock()
		}
	}
}

// Contain ...
func (s *SyncSet) Contain(v interface{}) bool {
	_, ok := s.arr.Load(v)
	return ok
}

// Len ...
func (s *SyncSet) Len() int {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.size
}

// Each ...
func (s *SyncSet) Each(cb func(v interface{})) {
	s.arr.Range(func(k, _ interface{}) bool {
		cb(k)
		return true
	})
}

// ToSlice ...
func (s *SyncSet) ToSlice() []interface{} {
	arr := make([]interface{}, 0, s.Len())
	s.Each(func(v interface{}) {
		arr = append(arr, v)

	})
	return arr
}

// IsEmpty ...
func (s *SyncSet) IsEmpty() bool {
	return s.Len() == 0
}

// Erase ...
func (s *SyncSet) Erase() {
	s.m.Lock()
	defer s.m.Unlock()
	s.arr = sync.Map{}
	s.size = 0
}
