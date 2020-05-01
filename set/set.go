package set

type Set struct {
	arr map[interface{}]struct{}
}

func New() *Set {
	return &Set{
		arr: make(map[interface{}]struct{}),
	}
}

func (set *Set) Add(v interface{}) {
	set.arr[v] = struct{}{}
}

func (set *Set) Delete(v interface{}) {
	delete(set.arr, v)
}

func (set *Set) Contain(v interface{}) bool {
	_, contain := set.arr[v]
	return contain
}

func (set *Set) Len() int {
	return len(set.arr)
}

func (set *Set) Each(cb func(interface{}) bool) {
	for v := range set.arr {
		if !cb(v) {
			break
		}
	}
}

func (set *Set) ToSlice() []interface{} {
	arr := make([]interface{}, 0, len(set.arr))
	for k := range set.arr {
		arr = append(arr, k)
	}
	return arr
}
