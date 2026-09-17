// variables_x002: Обмен значений
// Make the tests pass!
// I AM NOT DONE
//
// Функция swap должна поменять местами два значения и вернуть их.
// Тренирует: множественное присваивание в Go.
// Сложность: easy
package main_test

import "testing"

func swap(a, b string) (string, string) {
	a = b
	return a, b
}

func TestSwap(t *testing.T) {
	x, y := swap("left", "right")
	if x != "right" || y != "left" {
		t.Errorf("swap(left, right) = (%s, %s), want (right, left)", x, y)
	}
}
