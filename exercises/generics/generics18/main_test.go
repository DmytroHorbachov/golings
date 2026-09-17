// generics18
// Make the tests pass!

// I AM NOT DONE
//
// Contains проверяет наличие элемента в срезе.
// Код не компилируется: для == нужен другой ограничитель.
// Тренирует: ограничение comparable.
// Сложность: easy
package main_test

import "testing"

func Contains[T any](s []T, x T) bool {
	for _, v := range s {
		if v == x {
			return true
		}
	}
	return false
}

func TestContains(t *testing.T) {
	if !Contains([]string{"a", "b"}, "b") || Contains([]int{1, 2}, 3) {
		t.Errorf("Contains works incorrectly")
	}
}
