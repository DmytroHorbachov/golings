// generics39
// Make the tests pass!

// I AM NOT DONE
//
// CountIf считает элементы, удовлетворяющие условию.
// Тренирует: обобщённый предикат.
// Сложность: easy
package main_test

import "testing"

func CountIf[T any](s []T, pred func(T) bool) int {
	n := 0
	for _, v := range s {
		if pred(v) {
			n = 1
		}
	}
	return n
}

func TestCountIf(t *testing.T) {
	if CountIf([]int{1, 2, 3, 4}, func(x int) bool { return x > 1 }) != 3 {
		t.Errorf("CountIf works incorrectly")
	}
}
