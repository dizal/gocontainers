package indexstore

import "sync"

// IndexStore ...
type IndexStore struct {
	m     sync.RWMutex
	store sync.Map
	revst sync.Map
	size  uint32
	index uint32
}

// New ...
func New() *IndexStore {
	return &IndexStore{}
}

// Store - store value and return index
func (s *IndexStore) Store(value interface{}) uint32 {
	i, ok := s.store.Load(value)
	if !ok {
		s.m.Lock()
		s.store.Store(value, s.index)
		s.revst.Store(s.index, value)

		i = s.index
		s.index++
		s.size++
		s.m.Unlock()
	}

	return i.(uint32)
}

// GetIndex - get index by value
func (s *IndexStore) GetIndex(key interface{}) (uint32, bool) {
	if i, ok := s.store.Load(key); ok {
		return i.(uint32), ok
	}
	return 0, false
}

// GetValue - get value by index
func (s *IndexStore) GetValue(index uint32) (interface{}, bool) {
	if v, ok := s.revst.Load(index); ok {
		return v, ok
	}
	return nil, false
}

// GetStringValue - get value by index
func (s *IndexStore) GetStringValue(index uint32) (string, bool) {
	if v, ok := s.revst.Load(index); ok {
		if s, ok2 := v.(string); ok2 {
			return s, ok2
		}
	}
	return "", false
}

// Del -- delete value and index from Store
func (s *IndexStore) Del(value interface{}) {
	if i, ok := s.store.Load(value); ok {
		s.m.Lock()
		s.store.Delete(value)
		s.revst.Delete(i)
		s.size--
		s.m.Unlock()
	}
}

// Size of store
func (s *IndexStore) Size() uint32 {
	s.m.RLock()
	defer s.m.RUnlock()
	return s.size
}

// Erase store
func (s *IndexStore) Erase() {
	s.m.Lock()
	s.store = sync.Map{}
	s.revst = sync.Map{}
	s.index = 0
	s.size = 0
	s.m.Unlock()
}
