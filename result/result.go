package result

import (
	"github.com/avivatedgi/go-rust-std/iterator"
	"github.com/avivatedgi/go-rust-std/option"
)

// This Result implementation is based on the one in the Rust's standart library (https://doc.rust-lang.org/stable/std/result/enum.Result.html)
// The Result represents the result of an operation that may either succeed (Ok) or fail (Err).
type Result[T any, E error] struct {
	value *T
	err   *E
}

// Return a new Result containing a value.
func Ok[T any, E error](value T) Result[T, E] {
	return Result[T, E]{value: &value, err: nil}
}

// Return a new Result containing an error.
func Err[T any, E error](err E) Result[T, E] {
	return Result[T, E]{value: nil, err: &err}	
}

// Returns true if the result is Ok.
func (result Result[T, E]) IsOk() bool {
	return result.err == nil
}

// Returns true if the result is Ok wrapping a value matching the predicate.
func (result Result[T, E]) IsOkWith(f func(*T) bool) bool {
	return result.IsOk() && f(result.value)
}

// Returns true if the result is Err.
func (result Result[T, E]) IsErr() bool {
	return !result.IsOk()
}

// Returns true if the result is Err wrapping an error matching the predicate.
func (result Result[T, E]) IsErrWith(f func(*E) bool) bool {
	return result.IsErr() && f(result.err)
}

// Converts from Result[T, E] to Option[T].
// Converts result into an Option[T], consuming the result value, and discarding the error, if any.
func (result Result[T, _]) Ok() option.Option[T] {
	if result.IsOk() {
		return option.Some(*result.value)
	}

	return option.None[T]()
}

// Converts from Result[T, E] to Option[E].
// Converts result into an Option[E], consuming the error, and discarding the value, if any.
func (result Result[_, E]) Err() option.Option[E] {
	if result.IsErr() {
		return option.Some(*result.err)
	}

	return option.None[E]()
}

func (result Result[T, _]) Iter() iterator.Iterator[T] {
	return &iterator.SingleValueIterator[T]{Value: result.value}
}

// Returns the contained Ok value, consuming the self value.
// Panics if the value is an Err, with a panic message including the passed message, and the content of the Err.
func (result Result[T, E]) Expect(message string) T {
	if result.IsOk() {
		return *result.value
	}

	panic(message)
}

// Returns the contained Err value, consuming the self value.
// Panics if the value is an Ok, with a panic message including the passed message, and the content of the Ok.
func (result Result[T, E]) ExpectErr(message string) E {
	if result.IsErr() {
		return *result.err
	}

	panic(message)
}

// Returns the contained Ok value, consuming the self value.
// Because this function may panic, its use is generally discouraged.
// Instead, prefer to use pattern matching and handle the Err case explicitly, or call unwrap_or, unwrap_or_else, or unwrap_or_default.
// Panics if the value is an Err, with a panic message provided by the Err’s value.
func (result Result[T, E]) Unwrap() T {
	var err string

	if result.IsErr() {
		err = (*result.err).Error()
	}

	return result.Expect("called `Result::Unwrap()` on a `Err` value: \"" + err + "\"")
}

// Returns the contained Err value, consuming the self value.
// Panics if the value is an Ok, with a custom panic message provided by the Ok’s value.
func (result Result[T, E]) UnwrapErr() E {
	return result.ExpectErr("called `Result::UnwrapErr()` on a `Ok` value")
}

// Returns the contained Ok value or a provided default.
// Arguments passed to unwrap_or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use unwrap_or_else, which is lazily evaluated.
func (result Result[T, E]) UnwrapOr(other T) T {
	if result.IsOk() {
		return *result.value
	}

	return other
}

// Returns the contained Ok value or a default
// Consumes the self argument then, if Ok, returns the contained value, otherwise if Err, returns the default value for that type.
func (result Result[T, E]) UnwrapOrDefault() T {
	var zeroValue T
	return result.UnwrapOr(zeroValue)
}

// Returns the contained Ok value or computes it from a closure.
func (result Result[T, E]) UnwrapOrElse(f func(*E) T) T {
	if result.IsOk() {
		return *result.value
	}

	return f(result.err)
}

// Maps a Result[T, E] to Result[U, E] by applying a function to a contained Ok value, leaving an Err value untouched.
// This function can be used to compose the results of two functions.
func Map[T any, E error, U any](result Result[T, E], f func(*T) U) Result[U, E] {
	if result.IsOk() {
		return Ok[U, E](f(result.value))
	}

	return Err[U](*result.err)
}

// Returns the provided default (if Err), or applies a function to the contained value (if Ok),
func MapOr[T any, E error, U any](result Result[T, E], other U, f func(*T) U) U {
	if result.IsOk() {
		return f(result.value)
	}

	return other
}

// Maps a Result[T, E] to U by applying fallback function default to a contained Err value, or function f to a contained Ok value.
func MapOrElse[T any, E error, U any](result Result[T, E], def func(*E) U, f func(*T) U) U {
	if result.IsOk() {
		return f(result.value)
	}

	return def(result.err)
}

// Maps a Result[T, E] to Result[T, F] by applying a function to a contained Err value, leaving an Ok value untouched.
// This function can be used to pass through a successful result while handling an error.
func MapErr[T any, E error, F error](result Result[T, E], f func(*E) F) Result[T, F] {
	if result.IsOk() {
		return Ok[T, F](*result.value)
	}

	return Err[T](f(result.err))
}

// Returns other if the result is Ok, otherwise returns the Err value of result.
func And[T any, E error, U any](result Result[T, E], other Result[U, E]) Result[U, E] {
	if result.IsOk() {
		return other
	}

	return Err[U](*result.err)
}

// Calls f if the result is Ok, otherwise returns the Err value of result.
// This function can be used for control flow based on Result values.
func AndThen[T any, E error, U any](result Result[T, E], f func(*T) Result[U, E]) Result[U, E] {
	if result.IsOk() {
		return f(result.value)
	}

	return Err[U](*result.err)
}

// Returns other if the result is Err, otherwise returns the Ok value of result.
// Arguments passed to or are eagerly evaluated; if you are passing the result of a function call, it is recommended to use or_else, which is lazily evaluated.
func Or[T any, E error, F error](result Result[T, E], other Result[T, F]) Result[T, F] {
	if result.IsOk() {
		return Ok[T, F](*result.value)
	}

	return other
}

// Calls f if the result is Err, otherwise returns the Ok value of result.
// This function can be used for control flow based on result values.
func OrElse[T any, E error, F error](result Result[T, E], f func(*E) Result[T, F]) Result[T, F] {
	if result.IsOk() {
		return Ok[T, F](*result.value)
	}

	return f(result.err)
}
