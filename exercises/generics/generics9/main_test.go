// generics9
// Make the tests pass!

// I AM NOT DONE
//
// As пытается привести значение к типу T. Код не компилируется:
// утверждение типа применимо только к интерфейсному значению.
// Тренирует: v.(T) требует, чтобы v было интерфейсом.
// Сложность: hard
package main_test

import "testing"

func As[T any](v T) (T, bool) {
	r, ok := v.(T)
	return r, ok
}

func TestAs(t *testing.T) {
	if s, ok := As[string]("go"); !ok || s != "go" {
		t.Errorf("As[string] = %q, %v", s, ok)
	}
	if _, ok := As[int]("go"); ok {
		t.Errorf("As[int](\"go\") should fail")
	}
}
