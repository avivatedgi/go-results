package tests

import (
	"fmt"
	"testing"

	"github.com/avivatedgi/results/result"
)

type TestError struct {
	Value int
}

func (err TestError) Error() string {
	return fmt.Sprintf("test error %d", err.Value)
}

func TestResultIsOk(t *testing.T) {
	if !result.Ok[int, error](1).IsOk() {
		t.Error("expected `result.Ok(...).IsOk()` to be true")
	} else if result.Err[int](TestError{}).IsOk() {
		t.Error("expected `result.Err(...).IsOk()` to be false")
	}
}

func TestResultIsErr(t *testing.T) {
	if !result.Err[int](TestError{}).IsErr() {
		t.Error("expected `result.Err(...).IsErr()` to be true")
	} else if result.Ok[int, TestError](1).IsErr() {
		t.Error("expected `result.Ok(...).IsErr()` to be false")
	}
}

func TestResultIsOkWith(t *testing.T) {
	f := func(i *int) bool {
		return *i == 1
	}

	if !result.Ok[int, TestError](1).IsOkWith(f) {
		t.Error("expected `result.Ok(1).IsOkWith(...)` to be true")
	} else if result.Ok[int, TestError](2).IsOkWith(f) {
		t.Error("expected `result.Ok(2).IsOkWith(...)` to be false")
	} else if result.Err[int](TestError{}).IsOkWith(f) {
		t.Error("expected `result.Err(...).IsOkWith(...)` to be false")
	}
}

func TestResultIsErrWith(t *testing.T) {
	f := func(i *TestError) bool {
		return i.Value == 1
	}

	if !result.Err[int](TestError{Value: 1}).IsErrWith(f) {
		t.Error("expected `result.Err(TestError{1}).IsErrWith(...)` to be true")
	} else if result.Err[int](TestError{Value: 2}).IsErrWith(f) {
		t.Error("expected `result.Err(TestError{2}).IsErrWith(...)` to be false")
	} else if result.Ok[int, TestError](1).IsErrWith(f) {
		t.Error("expected `result.Ok(...).IsErrWith(...)` to be false")
	}
}

func TestResultOk(t *testing.T) {
	if result.Ok[int, TestError](1).Ok().Unwrap() != 1 {
		t.Error("expected `result.Ok(1).Ok()` to be Some(1)")
	} else if result.Err[int](TestError{}).Ok().IsSome() {
		t.Error("expected `result.Err(...).Ok()` to be None")
	}
}

func TestResultErr(t *testing.T) {
	err := TestError{}

	if result.Err[int](err).Err().Unwrap() != err {
		t.Error("expected `result.Err(TestError{}).Err()` to be Some(TestError{})")
	} else if result.Ok[int, TestError](1).Err().IsSome() {
		t.Error("expected `result.Ok(...).Err()` to be None")
	}
}

func TestResultIter(t *testing.T) {
	okIter := result.Ok[int, TestError](1).Iter()
	errIter := result.Err[int](TestError{}).Iter()

	if okIter.Next().Unwrap() != 1 {
		t.Error("expected `1st okIter.Next()` to be 1")
	} else if !okIter.Next().IsNone() {
		t.Error("expected `2nd okIter.Next()` to be None")
	} else if !errIter.Next().IsNone() {
		t.Error("expected `result.Err(...).Iter().Next()` to be None")
	}
}

func TestResultExpectSuccess(t *testing.T) {
	if result.Ok[int, TestError](1).Expect("should not panic") != 1 {
		t.Error("expected `result.Ok(1).Expect(...)` to be 1")
	}
}

func TestResultExpectPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected `result.Err(...).Expect` to panic")
		}
	}()

	result.Err[int](TestError{}).Expect("should panic")
}

func TestResultExpectErrSuccess(t *testing.T) {
	err := TestError{}

	if result.Err[int](err).ExpectErr("should not panic") != err {
		t.Error("expected `result.Err(TestError{}).ExpectErr(...)` to be TestError{}")
	}
}

func TestResultExpectErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected `result.Ok(...).ExpectErr` to panic")
		}
	}()

	result.Ok[int, TestError](1).ExpectErr("should panic")
}

func TestResultUnwrapSuccess(t *testing.T) {
	if result.Ok[int, TestError](1).Unwrap() != 1 {
		t.Error("expected `result.Ok(1).Unwrap()` to be 1")
	}
}

func TestResultUnwrapPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected `result.Err(...).Unwrap()` to panic")
		}
	}()

	result.Err[int](TestError{}).Unwrap()
}

func TestResultUnwrapErrSuccess(t *testing.T) {
	err := TestError{}
	if result.Err[int](err).UnwrapErr() != err {
		t.Error("expected `result.Err(TestError{}).UnwrapErr()` to be TestError{}")
	}
}

func TestResultUnwrapErrPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected `result.Ok(...).UnwrapErr()` to panic")
		}
	}()

	result.Ok[int, TestError](1).UnwrapErr()
}

func TestResultUnwrapOr(t *testing.T) {
	if result.Ok[int, TestError](1).UnwrapOr(2) != 1 {
		t.Error("expected `result.Ok(1).UnwrapOr(2)` to be 1")
	} else if result.Err[int](TestError{}).UnwrapOr(2) != 2 {
		t.Error("expected `result.Err(...).UnwrapOr(2)` to be 2")
	}
}

func TestResultUnwrapOrDefault(t *testing.T) {
	if result.Ok[int, TestError](1).UnwrapOrDefault() != 1 {
		t.Error("expected `result.Ok(1).UnwrapOrDefault()` to be 1")
	} else if result.Err[int](TestError{}).UnwrapOrDefault() != 0 {
		t.Error("expected `result.Err(...).UnwrapOrDefault()` to be 0")
	}
}

func TestResultUnwrapOrElse(t *testing.T) {
	if result.Ok[int, TestError](1).UnwrapOrElse(func(*TestError) int { return 2 }) != 1 {
		t.Error("expected `result.Ok(1).UnwrapOrElse(...)` to be 1")
	} else if result.Err[int](TestError{}).UnwrapOrElse(func(*TestError) int { return 2 }) != 2 {
		t.Error("expected `result.Err(...).UnwrapOrElse(...)` to be 2")
	}
}

func MapExample(s *string) int {
	return len(*s)
}

func TestResultMap(t *testing.T) {
	text := "Hello"
	ok := result.Ok[string, TestError](text)
	err := result.Err[string](TestError{})

	if result.Map(ok, MapExample).Unwrap() != len(text) {
		t.Error("expected `result.Map(...).Unwrap()` to be len(text)")
	} else if !result.Map(err, MapExample).IsErr() {
		t.Error("expected `result.Map(...)` to be Err")
	}
}

func TestResultMapOr(t *testing.T) {
	text := "Hello"
	def := 2
	ok := result.Ok[string, TestError](text)
	err := result.Err[string](TestError{})

	if result.MapOr(ok, def, MapExample) != len(text) {
		t.Error("expected `result.MapOr(ok, ...)` to be len(text)")
	} else if result.MapOr(err, def, MapExample) != def {
		t.Error("expected `result.MapOr(err, ...)` to be def")
	}
}

func TestResultMapOrElse(t *testing.T) {
	text := "Hello"
	ok := result.Ok[string, TestError](text)
	err := result.Err[string](TestError{})
	def := 2
	f := func(*TestError) int { return def }

	if result.MapOrElse(ok, f, MapExample) != len(text) {
		t.Error("expected `result.MapOrElse(ok, ...)` to be len(text)")
	} else if result.MapOrElse(err, f, MapExample) != def {
		t.Error("expected `result.MapOrElse(err, ...)` to be {def}")
	}
}

const MapErrExample_CustomError = "custom error"

func MapErrExample(s *TestError) error {
	return fmt.Errorf(MapErrExample_CustomError)
}

func TestResultMapErr(t *testing.T) {
	text := "Hello"
	ok := result.Ok[string, TestError](text)
	err := result.Err[string](TestError{})

	if result.MapErr(ok, MapErrExample).IsErr() {
		t.Error("expected `result.MapErr(ok, ...).IsErr()` to be false")
	} else if result.MapErr(err, MapErrExample).UnwrapErr().Error() != MapErrExample_CustomError {
		t.Error("expected `result.MapErr(err, ...).UnwrapErr().Error()` to be MapErrExample_CustomError")
	}
}

func TestResultAnd(t *testing.T) {
	x := result.Ok[int, TestError](1)
	y := result.Err[int](TestError{Value: 1})
	if result.And(x, y).UnwrapErr().Value != 1 {
		t.Error("expected `result.And(x, y)` to return TestError{1}")
	}

	x = result.Err[int](TestError{Value: 2})
	y = result.Ok[int, TestError](2)
	if result.And(x, y).UnwrapErr().Value != 2 {
		t.Error("expected `result.And(x, y)` to return TestError{2}")
	}

	x = result.Err[int](TestError{Value: 3})
	y = result.Err[int](TestError{Value: 4})
	if result.And(x, y).UnwrapErr().Value != 3 {
		t.Error("expected `result.And(x, y)` to return TestError{3}")
	}

	x = result.Ok[int, TestError](3)
	y = result.Ok[int, TestError](4)
	if result.And(x, y).Unwrap() != 4 {
		t.Error("expected `result.And(x, y)` to return 4")
	}
}

func TestResultAndThen(t *testing.T) {
	x := result.Ok[string, TestError]("Hello")
	y := result.Ok[string, TestError]("World")
	z := result.Err[string](TestError{Value: 1})

	f := func(s *string) result.Result[int, TestError] {
		if *s == "World" {
			return result.Err[int](TestError{Value: 2})
		}

		return result.Ok[int, TestError](len(*s))
	}

	if result.AndThen(x, f).Unwrap() != 5 {
		t.Error("expected `result.AndThen(x, f)` to return 5")
	} else if result.AndThen(y, f).UnwrapErr().Value != 2 {
		t.Error("expected `result.AndThen(y, f)` to return TestError{2}")
	} else if result.AndThen(z, f).UnwrapErr().Value != 1 {
		t.Error("expected `result.AndThen(z, f)` to return TestError{1}")
	}
}

func TestResultOr(t *testing.T) {
	x := result.Ok[int, TestError](1)
	y := result.Err[int](TestError{Value: 1})
	if result.Or(x, y).Unwrap() != 1 {
		t.Error("expected `result.Or(x, y)` to return 1")
	}

	x = result.Err[int](TestError{Value: 2})
	y = result.Ok[int, TestError](2)
	if result.Or(x, y).Unwrap() != 2 {
		t.Error("expected `result.Or(x, y)` to return 2")
	}

	x = result.Err[int](TestError{Value: 3})
	y = result.Err[int](TestError{Value: 4})
	if result.Or(x, y).UnwrapErr().Value != 4 {
		t.Error("expected `result.Or(x, y)` to return TestError{4}")
	}

	x = result.Ok[int, TestError](3)
	y = result.Ok[int, TestError](4)
	if result.Or(x, y).Unwrap() != 3 {
		t.Error("expected `result.Or(x, y)` to return 3")
	}
}

func TestResultOrElse(t *testing.T) {
	sq := func(i *TestError) result.Result[int, TestError] { return result.Ok[int, TestError](i.Value * i.Value) }
	err := func(i *TestError) result.Result[int, TestError] { return result.Err[int](TestError{Value: i.Value}) }

	if result.OrElse(result.Err[int](TestError{Value: 3}), sq).Unwrap() != 9 {
		t.Error("expected `result.OrElse(result.Err[int](TestError{3}), sq)` to return 9")
	} else if result.OrElse(result.Ok[int, TestError](4), sq).Unwrap() != 4 {
		t.Error("expected `result.OrElse(result.Ok[int, TestError](4), sq)` to return 4")
	} else if result.OrElse(result.Err[int](TestError{Value: 5}), err).UnwrapErr().Value != 5 {
		t.Error("expected `result.OrElse(result.Err[int](TestError{5}), err)` to return TestError{5}")
	} else if result.OrElse(result.Ok[int, TestError](6), err).Unwrap() != 6 {
		t.Error("expected `result.OrElse(result.Ok[int, TestError](6), err)` to return 6")
	}
}
