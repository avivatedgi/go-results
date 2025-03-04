<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# result

```go
import "github.com/avivatedgi/go-rust-std/result"
```

## Index

- [func MapOr[T any, E error, U any](result Result[T, E], other U, f func(*T) U) U](<#func-mapor>)
- [func MapOrElse[T any, E error, U any](result Result[T, E], def func(*E) U, f func(*T) U) U](<#func-maporelse>)
- [type Result](<#type-result>)
  - [func And[T any, E error, U any](result Result[T, E], other Result[U, E]) Result[U, E]](<#func-and>)
  - [func AndThen[T any, E error, U any](result Result[T, E], f func(*T) Result[U, E]) Result[U, E]](<#func-andthen>)
  - [func Err[T any, E error](err E) Result[T, E]](<#func-err>)
  - [func Map[T any, E error, U any](result Result[T, E], f func(*T) U) Result[U, E]](<#func-map>)
  - [func MapErr[T any, E error, F error](result Result[T, E], f func(*E) F) Result[T, F]](<#func-maperr>)
  - [func Ok[T any, E error](value T) Result[T, E]](<#func-ok>)
  - [func Or[T any, E error, F error](result Result[T, E], other Result[T, F]) Result[T, F]](<#func-or>)
  - [func OrElse[T any, E error, F error](result Result[T, E], f func(*E) Result[T, F]) Result[T, F]](<#func-orelse>)
  - [func (result Result[_, E]) Err() option.Option[E]](<#func-result_-e-err>)
  - [func (result Result[T, E]) Expect(message string) T](<#func-resultt-e-expect>)
  - [func (result Result[T, E]) ExpectErr(message string) E](<#func-resultt-e-expecterr>)
  - [func (result Result[T, E]) IsErr() bool](<#func-resultt-e-iserr>)
  - [func (result Result[T, E]) IsErrWith(f func(*E) bool) bool](<#func-resultt-e-iserrwith>)
  - [func (result Result[T, E]) IsOk() bool](<#func-resultt-e-isok>)
  - [func (result Result[T, E]) IsOkWith(f func(*T) bool) bool](<#func-resultt-e-isokwith>)
  - [func (result Result[T, _]) Ok() option.Option[T]](<#func-resultt-_-ok>)
  - [func (result Result[T, E]) Unwrap() T](<#func-resultt-e-unwrap>)
  - [func (result Result[T, E]) UnwrapErr() E](<#func-resultt-e-unwraperr>)
  - [func (result Result[T, E]) UnwrapOr(other T) T](<#func-resultt-e-unwrapor>)
  - [func (result Result[T, E]) UnwrapOrDefault() T](<#func-resultt-e-unwrapordefault>)
  - [func (result Result[T, E]) UnwrapOrElse(f func(*E) T) T](<#func-resultt-e-unwraporelse>)


## func MapOr

```go
func MapOr[T any, E error, U any](result Result[T, E], other U, f func(*T) U) U
```

Returns the provided default \(if Err\)\, or applies a function to the contained value \(if Ok\)\,

## func MapOrElse

```go
func MapOrElse[T any, E error, U any](result Result[T, E], def func(*E) U, f func(*T) U) U
```

Maps a Result\[T\, E\] to U by applying fallback function default to a contained Err value\, or function f to a contained Ok value\.

## type Result

This Result implementation is based on the one in the Rust's standart library \(https://doc.rust-lang.org/stable/std/result/enum.Result.html\) The Result represents the result of an operation that may either succeed \(Ok\) or fail \(Err\)\.

```go
type Result[T any, E error] struct {
    // contains filtered or unexported fields
}
```

### func And

```go
func And[T any, E error, U any](result Result[T, E], other Result[U, E]) Result[U, E]
```

Returns other if the result is Ok\, otherwise returns the Err value of result\.

### func AndThen

```go
func AndThen[T any, E error, U any](result Result[T, E], f func(*T) Result[U, E]) Result[U, E]
```

Calls f if the result is Ok\, otherwise returns the Err value of result\. This function can be used for control flow based on Result values\.

### func Err

```go
func Err[T any, E error](err E) Result[T, E]
```

Return a new Result containing an error\.

### func Map

```go
func Map[T any, E error, U any](result Result[T, E], f func(*T) U) Result[U, E]
```

Maps a Result\[T\, E\] to Result\[U\, E\] by applying a function to a contained Ok value\, leaving an Err value untouched\. This function can be used to compose the results of two functions\.

### func MapErr

```go
func MapErr[T any, E error, F error](result Result[T, E], f func(*E) F) Result[T, F]
```

Maps a Result\[T\, E\] to Result\[T\, F\] by applying a function to a contained Err value\, leaving an Ok value untouched\. This function can be used to pass through a successful result while handling an error\.

### func Ok

```go
func Ok[T any, E error](value T) Result[T, E]
```

Return a new Result containing a value\.

### func Or

```go
func Or[T any, E error, F error](result Result[T, E], other Result[T, F]) Result[T, F]
```

Returns other if the result is Err\, otherwise returns the Ok value of result\. Arguments passed to or are eagerly evaluated; if you are passing the result of a function call\, it is recommended to use or\_else\, which is lazily evaluated\.

### func OrElse

```go
func OrElse[T any, E error, F error](result Result[T, E], f func(*E) Result[T, F]) Result[T, F]
```

Calls f if the result is Err\, otherwise returns the Ok value of result\. This function can be used for control flow based on result values\.

### func \(Result\[\_\, E\]\) Err

```go
func (result Result[_, E]) Err() option.Option[E]
```

Converts from Result\[T\, E\] to Option\[E\]\. Converts result into an Option\[E\]\, consuming the error\, and discarding the value\, if any\.

### func \(Result\[T\, E\]\) Expect

```go
func (result Result[T, E]) Expect(message string) T
```

Returns the contained Ok value\, consuming the self value\. Panics if the value is an Err\, with a panic message including the passed message\, and the content of the Err\.

### func \(Result\[T\, E\]\) ExpectErr

```go
func (result Result[T, E]) ExpectErr(message string) E
```

Returns the contained Err value\, consuming the self value\. Panics if the value is an Ok\, with a panic message including the passed message\, and the content of the Ok\.

### func \(Result\[T\, E\]\) IsErr

```go
func (result Result[T, E]) IsErr() bool
```

Returns true if the result is Err\.

### func \(Result\[T\, E\]\) IsErrWith

```go
func (result Result[T, E]) IsErrWith(f func(*E) bool) bool
```

Returns true if the result is Err wrapping an error matching the predicate\.

### func \(Result\[T\, E\]\) IsOk

```go
func (result Result[T, E]) IsOk() bool
```

Returns true if the result is Ok\.

### func \(Result\[T\, E\]\) IsOkWith

```go
func (result Result[T, E]) IsOkWith(f func(*T) bool) bool
```

Returns true if the result is Ok wrapping a value matching the predicate\.

### func \(Result\[T\, \_\]\) Ok

```go
func (result Result[T, _]) Ok() option.Option[T]
```

Converts from Result\[T\, E\] to Option\[T\]\. Converts result into an Option\[T\]\, consuming the result value\, and discarding the error\, if any\.

### func \(Result\[T\, E\]\) Unwrap

```go
func (result Result[T, E]) Unwrap() T
```

Returns the contained Ok value\, consuming the self value\. Because this function may panic\, its use is generally discouraged\. Instead\, prefer to use pattern matching and handle the Err case explicitly\, or call unwrap\_or\, unwrap\_or\_else\, or unwrap\_or\_default\. Panics if the value is an Err\, with a panic message provided by the Err’s value\.

### func \(Result\[T\, E\]\) UnwrapErr

```go
func (result Result[T, E]) UnwrapErr() E
```

Returns the contained Err value\, consuming the self value\. Panics if the value is an Ok\, with a custom panic message provided by the Ok’s value\.

### func \(Result\[T\, E\]\) UnwrapOr

```go
func (result Result[T, E]) UnwrapOr(other T) T
```

Returns the contained Ok value or a provided default\. Arguments passed to unwrap\_or are eagerly evaluated; if you are passing the result of a function call\, it is recommended to use unwrap\_or\_else\, which is lazily evaluated\.

### func \(Result\[T\, E\]\) UnwrapOrDefault

```go
func (result Result[T, E]) UnwrapOrDefault() T
```

Returns the contained Ok value or a default Consumes the self argument then\, if Ok\, returns the contained value\, otherwise if Err\, returns the default value for that type\.

### func \(Result\[T\, E\]\) UnwrapOrElse

```go
func (result Result[T, E]) UnwrapOrElse(f func(*E) T) T
```

Returns the contained Ok value or computes it from a closure\.



Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)
