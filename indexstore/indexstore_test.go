package indexstore

import (
	"strconv"
	"testing"
)

func TestStoreNew(t *testing.T) {
	s := New()
	if s.Size() != 0 {
		t.Error("new: store is not empry")
	}
}

func TestStoreAdd(t *testing.T) {
	s := New()
	s.Store("value")
	s.Store("value")
	s.Store("value")

	if s.Size() != 1 {
		t.Error("add: Store length is not 1")
	}
}

func TestStoreGetIndex(t *testing.T) {
	s := New()
	s.Store("value1")
	s.Store("value2")
	s.Store("value3")

	if v, ok := s.GetIndex("value2"); !ok || v != 1 {
		t.Error("getIndex: index is not 1")
	}

	if s.Size() != 3 {
		t.Error("getIndex: Store length is not 3")
	}
}

func TestStoreGetValue(t *testing.T) {
	s := New()
	s.Store("value1")
	s.Store("value2")
	s.Store("value3")

	if v, ok := s.GetValue(1); !ok || v.(string) != "value2" {
		t.Error("getValue: value is not 'value2'")
	}
	if s.Size() != 3 {
		t.Error("getValue: Store length is not 3")
	}
}

func TestStoreDelete(t *testing.T) {
	s := New()
	s.Store("value1")
	s.Store("value2")
	s.Store("value3")

	if s.Size() != 3 {
		t.Error("delete: Store length is not 3")
	}

	s.Del("value1")
	s.Del("value2")
	s.Del("value3")

	if s.Size() != 0 {
		t.Error("delete: Store length is not 0")
	}
	if _, ok := s.GetIndex("value2"); ok {
		t.Error("delete: index exists")
	}
}

func BenchmarkStoreAdd(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	store := New()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, elem := range testSet {
			store.Store(elem)
		}
	}
}

func BenchmarkStoreDelete(b *testing.B) {
	b.ReportAllocs()
	var testSet []string
	for i := 0; i < 1024; i++ {
		testSet = append(testSet, strconv.Itoa(i))
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		store := New()
		for _, elem := range testSet {
			store.Store(elem)
		}
		b.StartTimer()
		for _, elem := range testSet {
			store.Del(elem)
		}
	}
}
