// anonymous_functions_x013: Литерал с двумя результатами
// Make the tests pass!
// I AM NOT DONE
//
// minMax вызывает литерал, возвращающий минимум и максимум пары.
// Тренирует: функциональный литерал с несколькими результатами.
// Сложность: easy
package main_test

import "testing"

func minMax(a, b int) (int, int) {
	order := func(x, y int) (int, int) {
		if x > y {
			return x, y
		}
		return x, y
	}
	return order(a, b)
}

func TestMinMax(t *testing.T) {
	if lo, hi := minMax(9, 2); lo != 2 || hi != 9 {
		t.Errorf("minMax = %d, %d", lo, hi)
	}
}
