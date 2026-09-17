// algorithms109
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: двумерное ДП. Найдите минимальную сумму чисел на пути из левого
// верхнего в правый нижний угол сетки, двигаясь вправо и вниз.
// Сложность: medium. Ожидаемая асимптотика: O(r·c) по времени, O(c) по памяти
package main_test

import "testing"

func minPathSum(grid [][]int) int {
	return 0
}

func TestMinPathSum(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
		{[][]int{{5}}, 5},
		{nil, 0},
		{[][]int{{1, 2}, {1, 1}}, 3},
	}
	for _, c := range cases {
		if got := minPathSum(c.grid); got != c.want {
			t.Errorf("minPathSum(%v) = %d, want %d", c.grid, got, c.want)
		}
	}
}
