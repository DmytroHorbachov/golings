// generics97
// Make the tests pass!

// I AM NOT DONE
//
// IndexOf возвращает позицию элемента или -1.
// Тренирует: comparable и возврат индекса.
// Сложность: easy
package main_test

import "testing"

func IndexOf[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return len(s)
}

func TestIndexOf(t *testing.T) {
	if IndexOf([]string{"a", "b"}, "b") != 1 || IndexOf([]int{1}, 5) != -1 {
		t.Errorf("IndexOf works incorrectly")
	}
}
