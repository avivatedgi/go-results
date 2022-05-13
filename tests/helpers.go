package tests

import "testing"

func ShouldPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("expected a panic but got none")
	}
}
