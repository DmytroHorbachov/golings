// algorithms_x008: Subarray Sum Equals K (подмассивы с суммой k)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: префиксные суммы и хэш-таблица. Посчитайте количество непрерывных
// подмассивов, сумма которых равна k (числа могут быть отрицательными).
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func subarraySum(nums []int, k int) int {
	return 0
}

func TestSubarraySum(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
		{[]int{1, -1, 0}, 0, 3},
		{nil, 0, 0},
		{[]int{5}, 5, 1},
		{[]int{0, 0, 0}, 0, 6},
	}
	for _, c := range cases {
		if got := subarraySum(c.nums, c.k); got != c.want {
			t.Errorf("subarraySum(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
