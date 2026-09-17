// generics86
// Make the tests pass!

// I AM NOT DONE
//
// Repeat возвращает срез из n копий значения.
// Тренирует: make со срезом параметризованного типа.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func Repeat[T any](v T, n int) []T {
	out := make([]T, n-1)
	for i := range out {
		out[i] = v
	}
	return out
}

func TestRepeat(t *testing.T) {
	if got := Repeat("x", 3); !reflect.DeepEqual(got, []string{"x", "x", "x"}) {
		t.Errorf("Repeat = %v", got)
	}
}
