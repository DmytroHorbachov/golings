// slices_x048: Создание сетки
// Make the tests pass!
// I AM NOT DONE
//
// newGrid создаёт сетку rows×cols, заполненную значением v.
// Тренирует: make для внешнего среза и для каждой строки.
// Сложность: medium
package main_test

import "testing"

func newGrid(rows, cols, v int) [][]int {
	g := make([][]int, rows)
	for r := range g {
		for c := 0; c < cols; c++ {
			g[r] = append(g[r], c)
		}
	}
	return g
}

func TestNewGrid(t *testing.T) {
	g := newGrid(2, 3, 7)
	if len(g) != 2 || len(g[1]) != 3 || g[1][2] != 7 || g[0][0] != 7 {
		t.Errorf("newGrid = %v", g)
	}
}
