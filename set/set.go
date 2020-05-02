package set

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
func (s *Set) Add(v ...interface{}) {
	for _, vv := range v {
		s.arr[vv] = struct{}{}
	}
}

// Delete ...
func (s *Set) Delete(v ...interface{}) {
	for _, vv := range v {
		delete(s.arr, vv)
	}
}

// Contain ...
func (s *Set) Contain(v interface{}) bool {
	_, ok := s.arr[v]
	return ok
}

// Len ...
func (s *Set) Len() int {
	return len(s.arr)
}

// Each ...
func (s *Set) Each(cb func(v interface{}) bool) {
	for v := range s.arr {
		if !cb(v) {
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
