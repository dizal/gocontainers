package indexstore

import "sync"

// IndexStore ...
type IndexStore struct {
	mx     sync.RWMutex
	store  map[interface{}]uint32
	revst  []interface{}
	size   uint32
	cap    uint32
	index  uint32
	bucket uint32
}

// New ...
func New() *IndexStore {
	return &IndexStore{
		store:  make(map[interface{}]uint32),
		revst:  make([]interface{}, 0),
		bucket: 10000,
	}
}

// NewWithBucket ...
func NewWithBucket(bucket uint32) *IndexStore {
	s := New()
	if bucket < 10 {
		bucket = 10
	}
	s.bucket = bucket
	return s
}

// Store - store value and return index
func (s *IndexStore) Store(value interface{}) uint32 {
	i, ok := s.GetIndex(value)
	if !ok {
		s.mx.Lock()
		s.store[value] = s.index
		s.increaseRevSize()
		s.revst[s.index] = value

		i = s.index
		s.index++
		s.size++
		s.mx.Unlock()
	}
	return i
}

func (s *IndexStore) increaseRevSize() {
	if s.index+1 <= s.cap {
		s.revst = s.revst[:s.index+1]
	} else {
		s.cap += s.bucket
		z := make([]interface{}, s.index+1, s.cap)
		copy(z, s.revst)
		s.revst = z
	}
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
	if index > s.index {
		return nil, false
	}
	v := s.revst[index]
	s.mx.RUnlock()
	return v, v != nil
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
		s.revst[i] = nil
		if s.size > 0 {
			s.size--
		}
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
	s.revst = make([]interface{}, 0)
	s.index = 0
	s.size = 0
	s.mx.Unlock()
}
