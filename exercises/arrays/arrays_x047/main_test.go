// arrays_x047: Соседи в «Жизни»
// Make the tests pass!
// I AM NOT DONE
//
// neighbors считает живых соседей клетки на поле 4×4 (края не зациклены).
// Тренирует: обход соседних индексов с проверкой границ.
// Сложность: medium
package main_test

import "testing"

func neighbors(g [4][4]bool, r, c int) int {
	n := 0
	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {
			rr, cc := r+dr, c+dc
			if g[rr][cc] {
				n++
			}
		}
	}
	return n
}

func TestNeighbors(t *testing.T) {
	var g [4][4]bool
	g[0][1], g[1][0], g[1][1], g[3][3] = true, true, true, true
	if got := neighbors(g, 0, 0); got != 3 {
		t.Errorf("neighbors(0,0) = %d, want 3", got)
	}
	if got := neighbors(g, 1, 1); got != 2 {
		t.Errorf("neighbors(1,1) = %d, want 2", got)
	}
	if got := neighbors(g, 3, 3); got != 0 {
		t.Errorf("neighbors(3,3) = %d, want 0", got)
	}
}
