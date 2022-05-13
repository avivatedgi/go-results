package iterator

import "github.com/avivatedgi/results/option"

type Iterator[T any] interface {
	// Advances the iterator and return the next value.
	// Returns None when iteration is finished.
	// Individual iterator implementations may choose to resume iteration, and so calling Next() again may or may not eventually start returning Some(T) again at some point.
	Next() option.Option[T]
}
