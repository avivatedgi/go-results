package tests

import (
	"testing"

	"github.com/avivatedgi/go-rust-std/collections"
)

func TestVectorSplice(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3, 4}
	new := collections.Vec[int]{7, 8, 9}
	splice := vec.Splice(1, 3, new)

	expectedVecValues := []int{1, 7, 8, 9, 4}
	expectedSpliceValues := []int{2, 3}

	// Check the lengths of both the vector and the splice
	if len(vec) != len(expectedVecValues) {
		t.Errorf("expected `len(vec)` to be %d", len(expectedVecValues))
	} else if len(splice) != len(expectedSpliceValues) {
		t.Errorf("expected `len(splice)` to be %d", len(expectedSpliceValues))
	}

	// Validate the vector values
	for i, v := range vec {
		if v != expectedVecValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedVecValues[i], vec[i])
		}
	}

	// Validate the splice values
	for i, v := range splice {
		if v != expectedSpliceValues[i] {
			t.Errorf("expected `splice[%d]` to be %d but got %d", i, expectedSpliceValues[i], splice[i])
		}
	}
}

// No need to check DedupBy because both Dedup and DedupByKey uses DedupBy
func TestVectorDedup(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 2, 3, 2}
	collections.Dedup(&vec)
	expectedValues := []int{1, 2, 3, 2}

	// Check that the length of the vector and the length of the expected values are the same
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `len(vec)` to be %d but got %d", len(expectedValues), len(vec))
	}

	// Validate the vector values
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorDedupByKey(t *testing.T) {
	vec := collections.Vec[int]{10, 20, 21, 22, 30, 20}
	collections.DedupByKey(&vec, func(v int) int {
		return v / 10
	})

	expectedValues := []int{10, 20, 30, 20}

	// Check that the length of the vector and the length of the expected values are the same
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `len(vec)` to be %d but got %d", len(expectedValues), len(vec))
	}

	// Validate the vector values
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorDrain(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3}
	iter := vec.Drain(1, vec.Len())

	if vec.Len() != 1 {
		t.Errorf("expected `vec.Len()` to be 1 but got %d", vec.Len())
	} else if vec[0] != 1 {
		t.Errorf("expected `vec[0]` to be 1 but got %d", vec[0])
	}

	expectedValues := []int{2, 3}
	counter := 0
	for {
		value := iter.Next()
		if value.IsNone() {
			break
		}

		if value.Unwrap() != expectedValues[counter] {
			t.Errorf("expected `iter[%d]` to be %d but got %d", counter, expectedValues[counter], value.Unwrap())
		}

		counter++
	}
}

func TestVectorExtend(t *testing.T) {
	vec := collections.Vec[int]{1}
	other := collections.Vec[int]{2, 3, 4}
	vec.Extend(&other)

	expectedValues := []int{1, 2, 3, 4}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], v)
		}
	}

	expectedValues = []int{2, 3, 4}
	for i, v := range other {
		if v != expectedValues[i] {
			t.Errorf("expected `other[%d]` to be %d but got %d", i, expectedValues[i], v)
		}
	}
}

func TestVectorInsert(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3}
	vec.Insert(1, 4)

	expectedValues := []int{1, 4, 2, 3}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}

	vec.Insert(4, 5)
	expectedValues = []int{1, 4, 2, 3, 5}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorInsertPanic(t *testing.T) {
	defer ShouldPanic(t)

	vec := collections.Vec[int]{}
	vec.Insert(1, 5)
}

func TestVectorPop(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3}
	value := vec.Pop().Unwrap()
	if value != 3 {
		t.Errorf("expected `vec.Pop.Unwrap` to be 3 but got %d", value)
	}

	expectedValues := []int{1, 2}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorPush(t *testing.T) {
	vec := collections.Vec[int]{1, 2}
	vec.Push(3)

	expectedValues := []int{1, 2, 3}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorRemove(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3}
	value := vec.Remove(1)
	if value != 2 {
		t.Errorf("expected `vec.Remove(1).Unwrap` to be 2 but got %d", value)
	}

	expectedValues := []int{1, 3}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorRemovePanic(t *testing.T) {
	defer ShouldPanic(t)

	vec := collections.Vec[int]{1}
	vec.Remove(1)
}

func TestVectorRemovePanicNegativeIndex(t *testing.T) {
	defer ShouldPanic(t)

	vec := collections.Vec[int]{1}
	vec.Remove(-1)
}

func TestVectorResize(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3}
	vec.Resize(2, 0)

	expectedValues := []int{1, 2}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}

	vec.Resize(5, 3)
	expectedValues = []int{1, 2, 3, 3, 3}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}

	vec.Resize(0, 0)
	if vec.Len() != 0 {
		t.Errorf("expected `vec.Len()` to be 0 but got %d", vec.Len())
	}
}

func TestVectorResizePanic(t *testing.T) {
	defer ShouldPanic(t)

	vec := collections.Vec[int]{1}
	vec.Resize(-1, 0)
}

func TestVectorResizeWithPanic(t *testing.T) {
	defer ShouldPanic(t)

	vec := collections.Vec[int]{1}
	vec.ResizeWith(-1, func() int { return 0 })
}

func TestVectorRetain(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3, 4}
	vec.Retain(func(value int) bool { return value%2 == 0 })
	expectedValues := []int{2, 4}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorSwapRemove(t *testing.T) {
	vec := collections.Vec[string]{"foo", "bar", "baz", "qux"}

	// Test swap remove
	if vec.SwapRemove(1) != "bar" {
		t.Error("expected `vec.SwapRemove(1)` to be `bar`")
	}

	expectedValues := []string{"foo", "qux", "baz"}
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `vec.Len()` to be %d but got %d", len(expectedValues), vec.Len())
	}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %s but got %s", i, expectedValues[i], vec[i])
		}
	}

	// Test again
	if vec.SwapRemove(0) != "foo" {
		t.Error("expected `vec.SwapRemove(0)` to be `foo`")
	}

	expectedValues = []string{"baz", "qux"}
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `vec.Len()` to be %d but got %d", len(expectedValues), vec.Len())
	}

	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %s but got %s", i, expectedValues[i], vec[i])
		}
	}
}

func TestVectorTruncate(t *testing.T) {
	vec := collections.Vec[int]{1, 2, 3, 4, 5}

	// Check truncate with normal index
	vec.Truncate(2)
	expectedValues := []int{1, 2}
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `vec.Len()` to be %d but got %d", len(expectedValues), vec.Len())
	}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}

	// Check truncate with out of bounds index
	vec.Truncate(10)
	expectedValues = []int{1, 2}
	if len(vec) != len(expectedValues) {
		t.Errorf("expected `vec.Len()` to be %d but got %d", len(expectedValues), vec.Len())
	}
	for i, v := range vec {
		if v != expectedValues[i] {
			t.Errorf("expected `vec[%d]` to be %d but got %d", i, expectedValues[i], vec[i])
		}
	}

	// Check truncate with zero index
	vec.Truncate(0)
	if len(vec) != 0 {
		t.Errorf("expected `vec.Len()` to be 0 but got %d", vec.Len())
	}
}

func TestVectorTruncatePanic(t *testing.T) {
	defer ShouldPanic(t)
	vec := collections.Vec[int]{1}

	// Check truncate with negative index
	vec.Truncate(-2)
}
