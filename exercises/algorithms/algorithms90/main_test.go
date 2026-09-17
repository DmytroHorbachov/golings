// algorithms90
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: два указателя. По высотам столбцов посчитайте, сколько воды
// задержится между ними после дождя.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func trap(h []int) int {
	return 0
}

func TestTrap(t *testing.T) {
	cases := []struct {
		h    []int
		want int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
		{nil, 0},
		{[]int{3}, 0},
		{[]int{1, 2, 3}, 0},
		{[]int{5, 0, 0, 0, 5}, 15},
	}
	for _, c := range cases {
		if got := trap(c.h); got != c.want {
			t.Errorf("trap(%v) = %d, want %d", c.h, got, c.want)
		}
	}
}
