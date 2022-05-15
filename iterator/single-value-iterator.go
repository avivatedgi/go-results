package iterator

import "github.com/avivatedgi/go-rust-std/option"

type SingleValueIterator[T any] struct {
	Value *T
}

func (iterator *SingleValueIterator[T]) Next() option.Option[T] {
	if iterator.Value == nil {
		return option.None[T]()
	}

	value := *iterator.Value
	iterator.Value = nil
	return option.Some(value)
}
