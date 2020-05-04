package indexstore

import "sync"

// IndexStore ...
type IndexStore struct {
	mx    sync.RWMutex
	store map[interface{}]uint32
	revst map[uint32]interface{}
	size  uint32
	index uint32
}

// New ...
func New() *IndexStore {
	return &IndexStore{
		revst: make(map[uint32]interface{}),
		store: make(map[interface{}]uint32),
	}
}

// Store - store value and return index
func (s *IndexStore) Store(value interface{}) uint32 {
	i, ok := s.GetIndex(value)
	if !ok {
		s.mx.Lock()
		s.store[value] = s.index
		s.revst[s.index] = value

		i = s.index
		s.index++
		s.size++
		s.mx.Unlock()
	}
	return i
}

// GetIndex - get index by value
func (s *IndexStore) GetIndex(value interface{}) (uint32, bool) {
	s.mx.RLock()
	i, ok := s.store[value]
	s.mx.RUnlock()
	return i, ok
}

// GetValue - get value by index
func (s *IndexStore) GetValue(index uint32) (interface{}, bool) {
	s.mx.RLock()
	v, ok := s.revst[index]
	s.mx.RUnlock()
	return v, ok
}

// GetStringValue - get value by index
func (s *IndexStore) GetStringValue(index uint32) (string, bool) {
	if v, ok := s.GetValue(index); ok {
		if s, ok2 := v.(string); ok2 {
			return s, ok
		}
	}
	return "", false
}

// Del -- delete value and index from Store
func (s *IndexStore) Del(value interface{}) {
	if i, ok := s.GetIndex(value); ok {
		s.mx.Lock()
		delete(s.store, value)
		delete(s.revst, i)
		s.size--
		s.mx.Unlock()
	}
}

// Size of store
func (s *IndexStore) Size() uint32 {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.size
}

// Erase store
func (s *IndexStore) Erase() {
	s.mx.Lock()
	s.store = make(map[interface{}]uint32)
	s.revst = make(map[uint32]interface{})
	s.index = 0
	s.size = 0
	s.mx.Unlock()
}
