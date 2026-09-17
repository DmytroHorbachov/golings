// algorithms63
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: обход в глубину по сетке. Посчитайте количество областей из '1',
// соединённых по горизонтали и вертикали.
// Сложность: medium. Ожидаемая асимптотика: O(r·c) по времени, O(r·c) по памяти
package main_test

import "testing"

func numIslands(grid []string) int {
	return 0
}

func TestNumIslands(t *testing.T) {
	cases := []struct {
		grid []string
		want int
	}{
		{[]string{"11110", "11010", "11000", "00000"}, 1},
		{[]string{"11000", "11000", "00100", "00011"}, 3},
		{[]string{"000"}, 0},
		{nil, 0},
		{[]string{"1"}, 1},
	}
	for _, c := range cases {
		if got := numIslands(c.grid); got != c.want {
			t.Errorf("numIslands(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}
