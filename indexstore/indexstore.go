package indexstore

import "sync"

type IndexStore[V comparable] interface {
	Add(value V) uint32
	GetIndex(value V) (uint32, bool)
	GetValue(index uint32) (V, bool)
	Delete(value V)
	Size() uint32
	Erase()
}

type store[V comparable] struct {
	mx     sync.RWMutex
	store  map[V]uint32
	revst  []V
	size   uint32
	cap    uint32
	index  uint32
	bucket uint32
}

// New ...
func New[V comparable]() IndexStore[V] {
	return NewWithBucket[V](10000)
}

// NewWithBucket ...
func NewWithBucket[V comparable](bucket uint32) IndexStore[V] {
	return &store[V]{
		store:  make(map[V]uint32),
		revst:  make([]V, 0),
		bucket: bucket,
	}
}

// Store new value and return index.
func (s *store[V]) Add(value V) uint32 {
	i, ok := s.GetIndex(value)
	if !ok {
		s.mx.Lock()
		defer s.mx.Unlock()

		s.store[value] = s.index
		s.increaseRevSize()
		s.revst[s.index] = value

		i = s.index
		s.index++
		s.size++

	}
	return i
}

func (s *store[V]) increaseRevSize() {
	if uint32(s.index+1) <= s.cap {
		s.revst = s.revst[:s.index+1]
	} else {
		s.cap += s.bucket
		z := make([]V, s.index+1, s.cap)
		copy(z, s.revst)
		s.revst = z
	}
}

// GetIndex get the index from the value in the store.
func (s *store[V]) GetIndex(value V) (uint32, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	i, ok := s.store[value]
	return i, ok
}

// GetValue get the value from the store by the index.
func (s *store[V]) GetValue(index uint32) (V, bool) {
	s.mx.RLock()
	defer s.mx.RUnlock()
	if index > s.index {
		return *new(V), false
	}

	v := s.revst[index]
	if v != *new(V) {
		return v, true
	}

	return *new(V), false
}

// Delete value and index from Store.
func (s *store[V]) Delete(value V) {
	if i, ok := s.GetIndex(value); ok {
		s.mx.Lock()
		delete(s.store, value)
		s.revst[i] = *new(V)
		if s.size > 0 {
			s.size--
		}
		s.mx.Unlock()
	}
}

// Size of store
func (s *store[V]) Size() uint32 {
	s.mx.RLock()
	defer s.mx.RUnlock()
	return s.size
}

// Erase store
func (s *store[V]) Erase() {
	s.mx.Lock()
	defer s.mx.Unlock()
	s.store = make(map[V]uint32)
	s.revst = make([]V, 0)
	s.index = 0
	s.size = 0
	s.cap = 0
}
