package collections

import (
	"fmt"

	"github.com/avivatedgi/go-rust-std/iterator"
	"github.com/avivatedgi/go-rust-std/option"
)

type Vec[T any] []T

// Moves all the elements of other into vec, leaving other empty.
func (vec *Vec[T]) Append(other *Vec[T]) {
	*vec = append(*vec, *other...)
	other.Clear()
}

// Returns the number of elements the vector can hold without reallocating.
func (vec Vec[T]) Capacity() int {
	return cap(vec)
}

// Returns the number of elements in the vector, also referred to as its ‘length’.
func (vec Vec[T]) Len() int {
	return len(vec)
}

// Clears the vector, removing all values.
// Note that this method has no effect on the allocated capacity of the vector.
func (vec *Vec[T]) Clear() {
	*vec = (*vec)[:0]
}

// Removes consecutive repeated elements in the vector.
// If the vector is sorted, this removes all duplicates.
//
// NOTE: This function isn't a method of the vector because it can only work on comparable types,
// and I didn't wanted to limit the Vector structure to hold only comparable types.
func Dedup[T comparable](vec *Vec[T]) {
	vec.DedupBy(func(a, b T) bool { return a == b })
}

// Removes all but the first of consecutive elements in the vector that resolve to the same key.
// If the vector is sorted, this removes all duplicates.
//
// NOTE: This function isn't a method of the vector because it can only work on comparable types,
// and I didn't wanted to limit the Vector structure to hold only comparable types.
func DedupByKey[T comparable](vec *Vec[T], key func(T) T) {
	vec.DedupBy(func(a, b T) bool { return key(a) == key(b) })
}

// Removes all but the first of consecutive elements in the vector satisfying a given equality relation.
// The f function is passed the two elements from the vector and must determine if the elements compare equal.
// The elements are passed in opposite order from their order in the slice, so if f(a, b) returns true, a is removed.
func (vec *Vec[T]) DedupBy(f func(T, T) bool) {
	temp := make(Vec[T], 0, len(*vec))

	for idx, item := range *vec {
		// If this is the first index, it can not be a duplicated item
		if idx == 0 {
			temp = append(temp, item)
			continue
		}

		previousItem := (*vec)[idx-1]
		if f(item, previousItem) {
			continue
		}

		temp = append(temp, item)
	}

	*vec = temp
}

// Removes the specified range from the vector in bulk, returning all removed elements as an iterator.
func (vec *Vec[T]) Drain(start, end int) iterator.Iterator[T] {
	return &sliceIterator[T]{data: vec.Splice(start, end, Vec[T]{})}
}

// Appends all elements in a slice to the Vec.
func (vec *Vec[T]) Extend(other *Vec[T]) {
	*vec = append(*vec, *other...)
}

// Inserts an element at position index within the vector, shifting all elements after it to the right.
// Panics if index > len.
func (vec *Vec[T]) Insert(index int, item T) {
	if index > vec.Len() || index < 0 {
		// Check if index is out of bound
		panic(fmt.Sprintf("insertion index (is %d) should be >= 0 <= len (is %d)", index, vec.Len()))
	} else if index == vec.Len() {
		// If index is equal to the length of the vector, append the item, it's better performance
		vec.Push(item)
		return
	}

	temp := make(Vec[T], 0, vec.Len()+1)
	temp = append(temp, (*vec)[:index]...)
	temp = append(temp, item)
	temp = append(temp, (*vec)[index:]...)
	*vec = temp
}

// Returns true if the vector contains no elements.
func (vec Vec[T]) IsEmpty() bool {
	return vec.Len() == 0
}

// Removes the last element from a vector and returns it, or None if it is empty.
func (vec *Vec[T]) Pop() option.Option[T] {
	if vec.Len() == 0 {
		return option.None[T]()
	}

	last := vec.Len() - 1
	value := (*vec)[last]
	*vec = (*vec)[:last]
	return option.Some(value)
}

// Appends an element to the back of a collection.
func (vec *Vec[T]) Push(item T) {
	*vec = append(*vec, item)
}

// Removes and returns the element at position index within the vector, shifting all elements after it to the left.
// Note: Because this shifts over the remaining elements, it has a worst-case performance of O(n).
// If you don’t need the order of elements to be preserved, use SwapRemove instead.
// Panics if index is out of bounds.
func (vec *Vec[T]) Remove(index int) T {
	if index >= vec.Len() || index < 0 {
		panic(fmt.Sprintf("index (is %d) should be >= 0 and < len (is %d)", index, vec.Len()))
	}

	value := (*vec)[index]
	*vec = append((*vec)[:index], (*vec)[index+1:]...)
	return value
}

// Resizes the Vec in-place so that len is equal to newLength.
// If newLength is greater than len, the Vec is extended by the difference, with each additional slot filled with value.
// If newLength is less than len, the Vec is simply truncated.
// Panics if index is less than zero
func (vec *Vec[T]) Resize(newLength int, value T) {
	vec.ResizeWith(newLength, func() T { return value })
}

// Resizes the Vec in-place so that len is equal to newLength.
// If newLength is greater than len, the Vec is extended by the difference, with each additional slot filled with the result of calling the function f.
// The return values from f will end up in the Vec in the order they have been generated.
//
// If new_len is less than len, the Vec is simply truncated.
// Panics if index is less than zero
func (vec *Vec[T]) ResizeWith(newLength int, f func() T) {
	if newLength < 0 {
		panic("length must be >= 0")
	}

	if newLength < vec.Len() {
		*vec = (*vec)[:newLength]
		return
	}

	value := f()
	for i := vec.Len(); i < newLength; i++ {
		*vec = append(*vec, value)
	}
}

// Retains only the elements specified by the predicate.
// In other words, remove all elements e such that f(&e) returns false.
// This method operates in place, visiting each element exactly once in the original order, and preserves the order of the retained elements.
func (vec *Vec[T]) Retain(f func(T) bool) {
	temp := make(Vec[T], 0, vec.Len())
	for _, item := range *vec {
		if f(item) {
			temp = append(temp, item)
		}
	}

	*vec = temp
}

// Creates a splicing vector that replaces the specified range in the vector with the given replaceWith vector and returns the removed items.
// replaceWith does not need to be the same length as range.
// range is removed even if the vector is not consumed until the end.
func (vec *Vec[T]) Splice(start, end int, replaceWith Vec[T]) Vec[T] {
	// Create a copy of the vecotr
	copied := make(Vec[T], len(*vec))
	copy(copied, *vec)

	// Change the vector to match the splice algorithm
	*vec = append((*vec)[:start], replaceWith...)
	*vec = append(*vec, copied[end:]...)

	// Return the splice
	return copied[start:end]
}

// Splits the collection into two at the given index.
// Returns a newly allocated vector containing the elements in the range [at, len].
// After the call, the original vector will be left containing the elements [0, at] with its previous capacity unchanged.
func (vec Vec[T]) SplitOff(at int) Vec[T] {
	return vec[:at]
}

// Removes an element from the vector and returns it.
// The removed element is replaced by the last element of the vector.
// This does not preserve ordering, but is O(1). If you need to preserve the element order, use remove instead.
// Panics if index is out of bounds.
func (vec *Vec[T]) SwapRemove(index int) T {
	if index >= vec.Len() || index < 0 {
		panic(fmt.Sprintf("index (is %d) should be >= 0 and < len (is %d)", index, vec.Len()))
	}

	value := (*vec)[index]
	last := vec.Len() - 1
	(*vec)[index] = (*vec)[last]
	*vec = (*vec)[:last]
	return value
}

// Shortens the vector, keeping the first len elements and dropping the rest.
// If len is greater than the vector’s current length, this has no effect.
// The drain method can emulate truncate, but causes the excess elements to be returned instead of dropped.
// Note that this method has no effect on the allocated capacity of the vector.
// Panics if index is negative.
func (vec *Vec[T]) Truncate(len int) {
	if len > vec.Len() {
		len = vec.Len()
	}

	*vec = (*vec)[:len]
}
