// primitive_types_x068: Деление с округлением вниз
// Make the tests pass!
// I AM NOT DONE
//
// floorDiv должна делить нацело с округлением к минус бесконечности: -7/2 = -4.
// Тренирует: целочисленное деление в Go округляет к нулю.
// Сложность: hard
package main_test

import "testing"

func floorDiv(a, b int) int {
	q := a / b
	return q
}

func TestFloorDiv(t *testing.T) {
	cases := [][3]int{{7, 2, 3}, {-7, 2, -4}, {7, -2, -4}, {-7, -2, 3}, {-8, 2, -4}, {0, 5, 0}}
	for _, c := range cases {
		if got := floorDiv(c[0], c[1]); got != c[2] {
			t.Errorf("floorDiv(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
