package collections

import (
	"github.com/avivatedgi/go-rust-std/option"
)

type sliceIterator[T any] struct {
	data  []T
	index int
}

func (iterator *sliceIterator[T]) Next() option.Option[T] {
	if iterator.index >= len(iterator.data) {
		return option.None[T]()
	}

	pair := iterator.data[iterator.index]
	iterator.index++
	return option.Some(pair)
}
