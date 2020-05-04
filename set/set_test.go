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

	s.Delete("value2")

	if s.Contain("value2") {
		t.Error("delete: value is contain")
	}

	if s.Len() != 1 {
		t.Error("delete: Set length is not 1")
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
