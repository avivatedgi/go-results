package results

// This Option implementation is based on the one in the Rust's standart library (https://doc.rust-lang.org/std/option/enum.Option.html)
// The Option represents an optional value: every Option is either Some and contains a value, or None, and does not.
type Option[T any] struct {
	value *T
}

// Return an Option containing no value.
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// Return an Option containing the value `value`.
func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

// Returns true if the option is a Some value.
func (option Option[T]) IsSome() bool {
	return option.value != nil
}

// Returns true if the option is a Some wrapping a value matching the predicate.
func (option Option[T]) IsSomeWith(f func(*T) bool) bool {
	return option.IsSome() && f(option.value)
}

// Returns true if the option is a None value.
func (option Option[T]) IsNone() bool {
	return !option.IsSome()
}

// Returns the contained Some value, consuming the option value.
// Panics if the value is a None with a custom panic message provided by message.
func (option Option[T]) Expect(message string) T {
	if option.IsNone() {
		panic(message)
	}

	return *option.value
}

// Returns the contained Some value, consuming the option value.
// Because this function may panic, its use is generally discouraged.
// Instead, prefer to use pattern matching and handle the None case explicitly, or call UnwrapOr, UnwrapOrElse, or UnwrapOrDefault.
// Panics if the self value equals None.
func (option Option[T]) Unwrap() T {
	return option.Expect("called `Option::Unwrap()` on a `None` value")
}

// Returns the contained Some value or a provided default.
func (option Option[T]) UnwrapOr(other T) T {
	if option.IsNone() {
		return other
	}

	return *option.value
}

// Returns the contained Some value or computes it from a closure.
func (option Option[T]) UnwrapOrElse(f func() T) T {
	if option.IsNone() {
		return f()
	}

	return *option.value
}

// Returns the contained Some value or a default.
// Consumes the self argument then, if Some, returns the contained value, otherwise if None, returns the default value for that type.
func (option Option[T]) UnwrapOrDefault() T {
	var zeroValue T
	return option.UnwrapOr(zeroValue)
}

// Maps an Option[T] to Option[U] by applying a function to a contained value.
// This function is not a method of option because method must have no type parameter.
// https://github.com/golang/go/issues/48793
func Map[T any, U any](option Option[T], f func(*T) U) Option[U] {
	if option.IsNone() {
		return None[U]()
	}

	return Some(f(option.value))
}

// Returns the provided default result (if none), or applies a function to the contained value (if any).
// This function is not a method of option because method must have no type parameter.
// https://github.com/golang/go/issues/48793
func MapOr[T any, U any](option Option[T], other U, f func(*T) U) U {
	if option.IsNone() {
		return other
	}

	return f(option.value)
}

// Computes a default function result (if none), or applies a different function to the contained value (if any).
// This function is not a method of option because method must have no type parameter.
// https://github.com/golang/go/issues/48793
func MapOrElse[T any, U any](option Option[T], def func() U, f func(*T) U) U {
	if option.IsNone() {
		return def()
	}

	return f(option.value)
}
