// algorithms6
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: ДП по числу попыток. Есть k яиц и здание из n этажей. Найдите
// минимальное число бросков, гарантированно определяющее критический этаж.
// Сложность: hard. Ожидаемая асимптотика: O(k·n) по времени, O(k) по памяти
package main_test

import "testing"

func superEggDrop(k, n int) int {
	return 0
}

func TestSuperEggDrop(t *testing.T) {
	cases := [][3]int{{1, 2, 2}, {2, 6, 3}, {3, 14, 4}, {1, 1, 1}, {2, 100, 14}, {10, 10000, 14}}
	for _, c := range cases {
		if got := superEggDrop(c[0], c[1]); got != c[2] {
			t.Errorf("superEggDrop(%d, %d) = %d, want %d", c[0], c[1], got, c[2])
		}
	}
}
