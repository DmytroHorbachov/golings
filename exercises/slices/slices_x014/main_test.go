// slices_x014: Двумерный срез
// Make the tests pass!
// I AM NOT DONE
//
// at возвращает элемент в строке r и столбце c.
// Тренирует: индексацию [][]T.
// Сложность: easy
package main_test

import "testing"

func at(grid [][]int, r, c int) int {
	return grid[c][r]
}

func TestAt(t *testing.T) {
	g := [][]int{{1, 2, 3}, {4, 5, 6}}
	if at(g, 1, 0) != 4 || at(g, 0, 2) != 3 {
		t.Errorf("at works incorrectly")
	}
}
