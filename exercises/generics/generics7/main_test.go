// generics7
// Make the tests pass!

// I AM NOT DONE
//
// Clear заменяет все элементы среза нулевыми значениями их типа.
// Тренирует: var zero T.
// Сложность: easy
package main_test

import "testing"

func Clear[T any](s []T) {
	var zero T
	for i := range s {
		s[i] = s[0]
	}
}

func TestClear(t *testing.T) {
	a := []string{"x", "y"}
	Clear(a)
	b := []*int{new(int)}
	Clear(b)
	if a[0] != "" || a[1] != "" || b[0] != nil {
		t.Errorf("Clear works incorrectly: %q %v", a, b)
	}
}
