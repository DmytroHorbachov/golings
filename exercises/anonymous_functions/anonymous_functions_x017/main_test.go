// anonymous_functions_x017: Вызов с аргументами
// Make the tests pass!
// I AM NOT DONE
//
// area вычисляется литералом, вызванным сразу с аргументами.
// Тренирует: передачу аргументов при немедленном вызове.
// Сложность: easy
package main_test

import "testing"

func area() int {
	return func(w, h int) int {
		return w * h
	}(3, 3)
}

func TestArea(t *testing.T) {
	if area() != 12 {
		t.Errorf("area = %d, want 12", area())
	}
}
