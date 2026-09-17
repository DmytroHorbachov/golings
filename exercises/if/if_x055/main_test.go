// if_x055: Ход коня
// Make the tests pass!
// I AM NOT DONE
//
// knightMove проверяет, может ли шахматный конь перейти из (r1, c1) в (r2, c2)
// на доске 8×8 (клетки 0..7).
// Тренирует: условия с модулем разности и проверкой границ.
// Сложность: medium
package main_test

import "testing"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func knightMove(r1, c1, r2, c2 int) bool {
	if r2 < 0 || r2 > 8 || c2 < 0 || c2 > 8 {
		return false
	}
	dr, dc := abs(r1-r2), abs(c1-c2)
	if dr == 1 && dc == 2 {
		return true
	}
	return false
}

func TestKnightMove(t *testing.T) {
	cases := []struct {
		r1, c1, r2, c2 int
		want           bool
	}{{0, 0, 1, 2, true}, {0, 0, 2, 1, true}, {4, 4, 2, 3, true}, {0, 0, 1, 1, false}, {7, 7, 9, 8, false}, {3, 3, 3, 5, false}, {6, 6, 8, 7, false}}
	for _, c := range cases {
		if got := knightMove(c.r1, c.c1, c.r2, c.c2); got != c.want {
			t.Errorf("knightMove(%d,%d -> %d,%d) = %v, want %v", c.r1, c.c1, c.r2, c.c2, got, c.want)
		}
	}
}
