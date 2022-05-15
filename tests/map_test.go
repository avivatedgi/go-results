package tests

import (
	"testing"

	"github.com/avivatedgi/go-rust-std/collections"
)

func TestMapInsertAndGet(t *testing.T) {
	m := make(collections.Map[int, int])

	if m.Insert(1, 5).IsSome() {
		t.Error("expected 1st `m.Insert` to be `None`")
	} else if m.Insert(1, 6).Unwrap() != 5 {
		t.Error("expected 2nd `m.Insert` to be `Some(5)`")
	} else if m.Get(1).Unwrap() != 6 {
		t.Error("expected `m.Get` to be `Some(6)`")
	}
}

func TestMapClear(t *testing.T) {
	m := make(collections.Map[int, int])

	m.Insert(1, 5)
	m.Insert(2, 6)
	m.Insert(3, 7)

	if len(m) != 3 {
		t.Error("expected `len(m)` to be 3")
	}

	m.Clear()
	if len(m) > 0 {
		t.Error("expected `len(m)` to be 0")
	}
}

func TestMapContainsKey(t *testing.T) {
	m := make(collections.Map[int, int])
	m.Insert(1, 5)

	if !m.ContainsKey(1) {
		t.Error("expected `m.ContainsKey(1)` to be true")
	} else if m.ContainsKey(2) {
		t.Error("expected `m.ContainsKey(2)` to be false")
	}
}

func TestMapDrain(t *testing.T) {
	// This test also checks Map.Clear and Map.Iter because drain is basiclly implemented
	// by both of this functions
	m := make(collections.Map[int, int])

	m.Insert(1, 5)
	m.Insert(2, 6)
	m.Insert(3, 7)

	if len(m) != 3 {
		t.Error("expected `len(m)` to be 3")
	}

	values := [][]int{
		{1, 5},
		{2, 6},
		{3, 7},
	}

	index := 0
	iter := m.Drain()

	for {
		item := iter.Next()
		if item.IsNone() {
			break
		}

		pair := item.Unwrap()
		if pair.First != values[index][0] || pair.Second != values[index][1] {
			t.Errorf("Expected pair to be (%d, %d) but got (%d, %d)", values[index][0], values[index][1], pair.First, pair.Second)
		}

		index += 1
	}

	if len(m) != 0 {
		t.Errorf("Expected map to be cleared")
	}
}

func TestMapEntry(t *testing.T) {
	m := make(collections.Map[int, int])

	// Test Map.Entry.OrInsert
	value := m.Entry(1).OrInsert(1)
	if m[1] != 1 || value != 1 {
		t.Errorf("Failed Map.Entry.OrInsert (m[1] = %d, value = %d)", m[1], value)
	}

	// Test Map.Entry.OrDefault
	value = m.Entry(2).OrDefault()
	if m[2] != 0 || value != 0 {
		t.Errorf("Failed Map.Entry.OrDefault (m[2] = %d, value = %d)", m[2], value)
	}

	// Test Map.Entry.AndModify on non existing key
	m.Entry(3).AndModify(func(v *int) {
		*v = 3
	})
	if !m.Get(3).IsNone() {
		t.Errorf("Expected Map.Get(3) to be None")
	}

	// Test Map.Entry.AndModify on existing key
	m.Entry(2).AndModify(func(v *int) {
		(*v)++
	})
	if m[2] != 1 {
		t.Errorf("Expected Map.Entry.AndModify to modify value to 1")
	}

	// Test Map.Entry.OrInsertWith
	m.Entry(4).OrInsertWith(func() int {
		return 3
	})
	if m[4] != 3 {
		t.Errorf("Failed Map.Entry.OrInsertWith")
	}

	// Test Map.Entry.OrInsertWithKey
	value = m.Entry(5).OrInsertWithKey(func(k int) int {
		return k
	})
	if m[5] != 5 || value != 5 {
		t.Errorf("Failed Map.Entry.OrInsertWithKey (m[5] = %d, value = %d)", m[5], value)
	}

	// Test Map.Entry.AndModify.OrInsert
	value = m.Entry(6).AndModify(func(v *int) { (*v)++ }).OrInsert(3)
	if m[6] != 3 || value != 3 {
		t.Errorf("Failed Map.Entry.AndModify.OrInsert (m[6] = %d, value = %d)", m[6], value)
	}

	value = m.Entry(6).AndModify(func(v *int) { (*v)++ }).OrInsert(3)
	if m[6] != 4 || value != 4 {
		t.Errorf("Failed Map.Entry.AndModify.OrInsert (m[6] = %d, value = %d)", m[6], value)
	}
}

func TestMapGetKeyValue(t *testing.T) {
	m := make(collections.Map[int, int])
	m.Insert(1, 5)

	pair := m.GetKeyValue(1)
	if pair.IsNone() {
		t.Errorf("Expected pair to be Some")
	}

	if pair.Unwrap().First != 1 || pair.Unwrap().Second != 5 {
		t.Errorf("Expected pair to be (1, 5) but got (%d, %d)", pair.Unwrap().First, pair.Unwrap().Second)
	}
}

func TestMapKeysValuesAndPairs(t *testing.T) {
	m := make(collections.Map[int, int])
	m.Insert(1, 5)
	m.Insert(2, 6)
	m.Insert(3, 7)

	// Test Map.Keys
	// Cannot predict order so just check if it exists
	expectedKeys := make(map[int]bool)
	expectedKeys[1] = false
	expectedKeys[2] = false
	expectedKeys[3] = false
	for _, key := range m.Keys() {
		if _, ok := expectedKeys[key]; !ok {
			t.Errorf("Expected key %d to be in expectedKeys", key)
		}

		expectedKeys[key] = true
	}

	for key, found := range expectedKeys {
		if !found {
			t.Errorf("Expected key %d to be in expectedKeys", key)
		}
	}

	// Test Map.Values
	// Cannot predict order so just check if it exists
	expectedValues := make(map[int]bool)
	expectedValues[5] = false
	expectedValues[6] = false
	expectedValues[7] = false
	for _, value := range m.Values() {
		if _, ok := expectedValues[value]; !ok {
			t.Errorf("Expected value %d to be in expectedValues", value)
		}

		expectedValues[value] = true
	}

	for value, found := range expectedValues {
		if !found {
			t.Errorf("Expected value %d to be in expectedValues", value)
		}
	}

	// Test Map.KeyValuePairs
	// Cannot predict order so just check if it exists
	expectedPairs := make(map[int]int)
	found := make(map[int]bool)
	expectedPairs[1] = 5
	expectedPairs[2] = 6
	expectedPairs[3] = 7
	for _, pair := range m.KeyValuePairs() {
		if _, ok := expectedPairs[pair.First]; !ok {
			t.Errorf("Expected key %d to be in expectedPairs", pair.First)
		} else if expectedPairs[pair.First] != pair.Second {
			t.Errorf("Expected value to be %d but got %d", expectedPairs[pair.First], pair.Second)
		}

		found[pair.First] = true
	}

	for key, found := range found {
		if !found {
			t.Errorf("Expected key %d to be in found", key)
		}
	}
}

func TestMapForEach(t *testing.T) {
	m := make(collections.Map[int, int])
	m.Insert(1, 5)
	m.Insert(2, 6)
	m.Insert(3, 7)

	// Test Map.ForEach
	x := make(map[int]int)
	m.ForEach(func(key *int, value *int) bool {
		x[*key] = *value
		return true
	})

	found := make(map[int]bool)
	for key, value := range x {
		if m[key] != value {
			t.Errorf("Expected value to be %d but got %d", m[key], value)
		}

		found[key] = true
	}

	for key, found := range found {
		if !found {
			t.Errorf("Expected key %d to be in found", key)
		}
	}

	// Test Map.ForEach with pausing the loop
	counter := 0
	m.ForEach(func(*int, *int) bool {
		counter++
		return false
	})

	if counter != 1 {
		t.Errorf("Expected counter to be 1 but got %d", counter)
	}
}
