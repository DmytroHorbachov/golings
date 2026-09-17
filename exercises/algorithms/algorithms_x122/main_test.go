// algorithms_x122: Dungeon Game (подземелье)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двумерное ДП с конца. Рыцарь идёт из левого верхнего в правый нижний
// угол (вправо и вниз); в клетках он теряет или получает здоровье. Найдите
// минимальное стартовое здоровье, при котором оно всегда остаётся положительным.
// Сложность: hard. Ожидаемая асимптотика: O(r·c) по времени, O(c) по памяти
package main_test

import "testing"

func calculateMinimumHP(dungeon [][]int) int {
	return 0
}

func TestCalculateMinimumHP(t *testing.T) {
	cases := []struct {
		d    [][]int
		want int
	}{
		{[][]int{{-2, -3, 3}, {-5, -10, 1}, {10, 30, -5}}, 7},
		{[][]int{{0}}, 1},
		{[][]int{{-5}}, 6},
		{[][]int{{100}}, 1},
		{nil, 1},
		{[][]int{{1, -3, 3}, {0, -2, 0}, {-3, -3, -3}}, 3},
	}
	for _, c := range cases {
		if got := calculateMinimumHP(c.d); got != c.want {
			t.Errorf("calculateMinimumHP(%v) = %d, want %d", c.d, got, c.want)
		}
	}
}
