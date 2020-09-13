package set

import (
	"strconv"
	"testing"
)

func TestSetNew(t *testing.T) {
	s := New()

	if !s.IsEmpty() {
		t.Error("new: set is not empry")
	}
}

func TestSetAdd(t *testing.T) {
	s := New()
	s.Add("value")
	s.Add("value")

	if s.Len() != 1 {
		t.Error("add: Set length is not 1")
	}
}

func TestSetContain(t *testing.T) {
	s := New()
	s.Add("value")

	if !s.Contain("value") {
		t.Error("contain: value is not contain")
	}
}

func TestSetDelete(t *testing.T) {
	s := New()
	s.Add("value")
	s.Add("value2")

	ok := s.Delete("value2")

	if !ok {
		t.Error("delete: delete status is false")
	}

	if s.Contain("value2") {
		t.Error("delete: value is contain")
	}

	if s.Len() != 1 {
		t.Error("delete: Set length is not 1")
	}

	ok = s.Delete("value2")

	if ok {
		t.Error("delete: delete status is true")
	}
}

func TestSetErase(t *testing.T) {
	s := New()
	s.Add("value")
	s.Add("value1")
	s.Add("value2")
	s.Add("value3")

	s.Erase()

	if s.Contain("value") {
		t.Error("erase: value is contain")
	}

	if s.Len() != 0 {
		t.Error("erase: set length is not null")
	}
}

func TestSetSlice(t *testing.T) {
	s := New()
	s.Add("value1")
	s.Add("value2")

	slice := s.ToSlice()

	if len(slice) != 2 {
		t.Error("slice: slice length must be 2")
	}
}

func TestSetString(t *testing.T) {
	s := New()

	if s.String() != "SET<>" {
		t.Errorf("string: Expected SET<>. Received %s", s.String())
	}

	s.Add("v1")
	s.Add("v2")

	if s.String() != "SET<v1, v2>" {
		t.Errorf("string: Expected SET<v1, v2>. Received %s", s.String())
	}
}

func TestSetEach(t *testing.T) {
	s := New()

	s.Add("v1")
	s.Add("v2")
	s.Add("v3")

	s.Each(func(v interface{}) bool {
		if v != "v1" && v != "v2" && v != "v3" {
			t.Errorf("Each: Expected 'v1' or 'v2' or 'v3'. Received '%s'", v)
		}
		return true
	})
}

func BenchmarkSetAdd(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		set := New()
		b.StartTimer()
		for _, elem := range testSet {
			set.Add(elem)
		}
	}
}
