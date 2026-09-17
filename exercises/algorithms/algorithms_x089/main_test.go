// algorithms_x089: Rotting Oranges (гниющие апельсины)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: обход в ширину по уровням. Клетки: 0 — пусто, 1 — свежий апельсин,
// 2 — гнилой. Каждую минуту гниль переходит на соседей. Верните минуты,
// за которые сгниют все, или -1, если это невозможно.
// Сложность: medium. Ожидаемая асимптотика: O(r·c) по времени, O(r·c) по памяти
package main_test

import "testing"

func orangesRotting(grid [][]int) int {
	return 0
}

func TestOrangesRotting(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4},
		{[][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1},
		{[][]int{{0, 2}}, 0},
		{[][]int{{0}}, 0},
		{[][]int{{1}}, -1},
	}
	for _, c := range cases {
		if got := orangesRotting(c.grid); got != c.want {
			t.Errorf("orangesRotting = %d, want %d", got, c.want)
		}
	}
}
