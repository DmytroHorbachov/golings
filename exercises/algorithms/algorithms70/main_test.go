// algorithms70
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: перебор с возвратом и множествами занятых линий. Посчитайте,
// сколькими способами можно расставить n ферзей на доске n×n так,
// чтобы они не били друг друга.
// Сложность: hard. Ожидаемая асимптотика: O(n!) по времени, O(n) по памяти
package main_test

import "testing"

func totalNQueens(n int) int {
	return 0
}

func TestTotalNQueens(t *testing.T) {
	cases := map[int]int{1: 1, 2: 0, 3: 0, 4: 2, 5: 10, 6: 4, 8: 92}
	for n, want := range cases {
		if got := totalNQueens(n); got != want {
			t.Errorf("totalNQueens(%d) = %d, want %d", n, got, want)
		}
	}
}
