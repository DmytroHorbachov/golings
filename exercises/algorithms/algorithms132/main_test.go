// algorithms132
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: одномерное ДП. Нельзя грабить два соседних дома. Найдите
// максимальную сумму, которую можно унести.
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func rob(nums []int) int {
	return 0
}

func TestRob(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
		{nil, 0},
		{[]int{5}, 5},
		{[]int{2, 1, 1, 2}, 4},
	}
	for _, c := range cases {
		if got := rob(c.nums); got != c.want {
			t.Errorf("rob(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
