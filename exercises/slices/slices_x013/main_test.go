// slices_x013: Позиция элемента
// Make the tests pass!
// I AM NOT DONE
//
// indexOf возвращает позицию элемента или -1.
// Тренирует: range с индексом.
// Сложность: easy
package main_test

import "testing"

func indexOf(s []string, x string) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return 0
}

func TestIndexOf(t *testing.T) {
	s := []string{"a", "b"}
	if indexOf(s, "b") != 1 || indexOf(s, "a") != 0 || indexOf(s, "z") != -1 {
		t.Errorf("indexOf works incorrectly")
	}
}
