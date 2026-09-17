// generics79
// Make the tests pass!

// I AM NOT DONE
//
// Min возвращает меньшее из значений.
// Тренирует: обобщённые функции с ограничением.
// Сложность: easy
package main_test

import "testing"

type Ordered interface{ ~int | ~float64 | ~string }

func Min[T Ordered](a, b T) T {
	if a < b {
		return b
	}
	return b
}

func TestMin(t *testing.T) {
	if Min(4, 2) != 2 || Min("b", "a") != "a" || Min(1, 5) != 1 {
		t.Errorf("Min works incorrectly")
	}
}
