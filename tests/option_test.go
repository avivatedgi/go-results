package tests

import (
	"testing"

	"github.com/avivatedgi/go-rust-std/option"
)

func TestOptionNone(t *testing.T) {
	option := option.None[int]()
	if !option.IsNone() {
		t.Error("expected `option.IsNone()` to be true")
	} else if option.IsSome() {
		t.Error("expected `option.IsSome()` to be false")
	}
}

func TestOptionSome(t *testing.T) {
	option := option.Some(1)
	if option.IsNone() {
		t.Error("expected `option.IsNone()` to be false")
	} else if !option.IsSome() {
		t.Error("expected `option.IsSome()` to be true")
	}
}

func TestOptionIsSomeWith(t *testing.T) {
	a := option.Some(1)
	b := option.Some(2)
	checker := func(option *int) bool {
		return *option%2 == 0
	}

	if a.IsSomeWith(checker) {
		t.Error("expected `a.IsSomeWith(checker)` to be false")
	} else if !b.IsSomeWith(checker) {
		t.Error("expected `b.IsSomeWith(checker)` to be true")
	}
}

func TestOptionExpectSuccess(t *testing.T) {
	a := option.Some(1)

	if a.Expect("should not panic") != 1 {
		t.Error("expected `a.Expect` to be 1")
	}
}

func TestOptionExpectPanic(t *testing.T) {
	defer ShouldPanic(t)

	option.None[int]().Expect("should panic")
}

func TestOptionUnwrapSuccess(t *testing.T) {
	if option.Some(1).Unwrap() != 1 {
		t.Error("expected `a.Unwrap` to be 1")
	}
}

func TestOptionUnwrapPanic(t *testing.T) {
	defer ShouldPanic(t)

	option.None[int]().Unwrap()
}

func TestOptionUnwrapOr(t *testing.T) {
	if option.Some(5).UnwrapOr(4) != 5 {
		t.Error("expected `UnwrapOr` to be 5")
	} else if option.None[int]().UnwrapOr(4) != 4 {
		t.Error("expected `UnwrapOr` to be 4")
	}
}

func TestOptionUnwrapOrElse(t *testing.T) {
	value := 35
	function := func() int {
		return value + 1
	}

	if option.Some(5).UnwrapOrElse(function) != 5 {
		t.Error("expected `UnwrapOr` to be 5")
	} else if option.None[int]().UnwrapOrElse(function) != 36 {
		t.Error("expected `UnwrapOr` to be 36")
	}
}

func TestOptionUnwrapOrDefault(t *testing.T) {
	type Example struct {
		A string
		B int
		C float32
		D float64
		E *string
	}

	example := Example{}

	if option.Some(5).UnwrapOrDefault() != 5 {
		t.Error("expected `UnwrapOrDefault` to be 5")
	} else if option.None[int]().UnwrapOrDefault() != 0 {
		t.Error("expected `UnwrapOrDefault` to be 0")
	} else if option.None[string]().UnwrapOrDefault() != "" {
		t.Error("expected `UnwrapOrDefault` to be \"\"")
	} else if option.None[float32]().UnwrapOrDefault() != 0.0 {
		t.Error("expected `UnwrapOrDefault` to be 0.0")
	} else if option.None[Example]().UnwrapOrDefault() != example {
		t.Error("expected `UnwrapOrDefault` to be Example{}")
	}
}

func OptionMapExample(data *string) int {
	return len(*data)
}

func TestOptionMap(t *testing.T) {
	if option.Map(option.None[string](), OptionMapExample).IsSome() {
		t.Error("expected `option.Map` to be None")
	}

	if option.Map(option.Some("hello"), OptionMapExample).Unwrap() != 5 {
		t.Error("expected `option.Map` to be Some(5)")
	}
}

func TestOptionMapOr(t *testing.T) {
	if option.MapOr(option.None[string](), 6, OptionMapExample) != 6 {
		t.Error("expected `option.MapOr` to be 6")
	} else if option.MapOr(option.Some("hello"), 6, OptionMapExample) != 5 {
		t.Error("expected `option.MapOr` to be 5")
	}
}

func TestOptionMapOrElse(t *testing.T) {
	value := 35
	defaultFunction := func() int {
		return value + 1
	}

	if option.MapOrElse(option.None[string](), defaultFunction, OptionMapExample) != 36 {
		t.Error("expected `option.MapOrElse` to be 36")
	} else if option.MapOrElse(option.Some("hello"), defaultFunction, OptionMapExample) != 5 {
		t.Error("expected `option.MapOrElse` to be 5")
	}
}
