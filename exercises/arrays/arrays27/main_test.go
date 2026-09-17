// arrays27
// Make the tests pass!

// I AM NOT DONE
//
// cols возвращает количество столбцов матрицы 2×4.
// Тренирует: len для вложенного массива.
// Сложность: easy
package main_test

import "testing"

func cols(m [2][4]int) int {
	return len(m)
}

func TestCols(t *testing.T) {
	if got := cols([2][4]int{}); got != 4 {
		t.Errorf("cols = %d, want 4", got)
	}
}
