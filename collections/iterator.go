package collections

// A channel based iterator, used to abstarct a channel as an iterator.
// Usage example (taken from collections.Vector[T]):
// 	for value := range vec.Iter() {
//		fmt.Println(value)
// 	}
type Iterator[T any] chan T

// Push a value into the channel, passing nil as the value will be ignored.
func (ch *Iterator[T]) Push(value *T) {
	if value == nil {
		return
	}

	*ch <- *value
}

// Close the channel, this function MUST be called after we are done use Push for all of the values we want to iterate.
func (ch *Iterator[T]) Close() {
	close(*ch)
}

// Convert an iterator into a vector of the same type.
func (ch *Iterator[T]) IntoVector() *Vec[T] {
	vec := Vec[T]{}

	for value := range *ch {
		vec.Push(value)
	}

	return &vec
}
