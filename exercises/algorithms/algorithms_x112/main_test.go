// algorithms_x112: Unique Paths (число путей в сетке)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двумерное ДП. Сколько путей из левого верхнего угла сетки r×c
// в правый нижний, если ходить можно только вправо и вниз?
// Сложность: medium. Ожидаемая асимптотика: O(r·c) по времени, O(c) по памяти
package main_test

import "testing"

func uniquePaths(rows, cols int) int {
	return 0
}

func TestUniquePaths(t *testing.T) {
	cases := [][3]int{{3, 7, 28}, {3, 2, 3}, {1, 1, 1}, {1, 10, 1}, {0, 5, 0}, {10, 10, 48620}}
	for _, c := range cases {
		if got := uniquePaths(c[0], c[1]); got != c[2] {
			t.Errorf("uniquePaths(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
