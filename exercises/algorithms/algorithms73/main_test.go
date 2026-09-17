// algorithms73
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: сортировка слиянием с подсчётом. Посчитайте пары индексов i < j,
// для которых nums[i] > nums[j]. Исходный срез менять нельзя.
// Сложность: hard. Ожидаемая асимптотика: O(n·log n) по времени, O(n) по памяти
package main_test

import "testing"

func countInversions(nums []int) int {
	return 0
}

func TestCountInversions(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 4, 1, 3, 5}, 3},
		{[]int{5, 4, 3, 2, 1}, 10},
		{[]int{1, 2, 3}, 0},
		{nil, 0},
		{[]int{1}, 0},
		{[]int{2, 2, 2}, 0},
	}
	for _, c := range cases {
		in := append([]int{}, c.nums...)
		if got := countInversions(in); got != c.want {
			t.Errorf("countInversions(%v) = %d, want %d", c.nums, got, c.want)
		}
		for i := range in {
			if in[i] != c.nums[i] {
				t.Errorf("input modified: %v", in)
				break
			}
		}
	}
}
