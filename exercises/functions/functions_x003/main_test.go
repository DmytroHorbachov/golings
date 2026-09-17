// functions_x003: Порядок результатов
// Make the tests pass!
// I AM NOT DONE
//
// Функция minMax должна вернуть сначала минимум, потом максимум.
// Тренирует: возврат нескольких значений.
// Сложность: easy
package main_test

import "testing"

func minMax(a, b int) (int, int) {
	if a > b {
		a, b = b, a
	}
	return b, a
}

func TestMinMax(t *testing.T) {
	lo, hi := minMax(9, 4)
	if lo != 4 || hi != 9 {
		t.Errorf("minMax(9, 4) = %d, %d; want 4, 9", lo, hi)
	}
}
