// algorithms5
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: скользящее окно фиксированного размера. Верните наибольшую сумму
// k подряд идущих элементов; если элементов меньше k — 0.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func maxSumK(nums []int, k int) int {
	return 0
}

func TestMaxSumK(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{2, 1, 5, 1, 3, 2}, 3, 9},
		{[]int{-1, -2, -3}, 2, -3},
		{[]int{4}, 1, 4},
		{[]int{1, 2}, 3, 0},
		{nil, 1, 0},
		{[]int{1 << 40, 1 << 40}, 2, 1 << 41},
	}
	for _, c := range cases {
		if got := maxSumK(c.nums, c.k); got != c.want {
			t.Errorf("maxSumK(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
