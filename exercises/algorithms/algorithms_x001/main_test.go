// algorithms_x001: Two Sum (сумма двух чисел)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: массивы и хэш-таблицы. Найдите индексы i < j двух разных элементов,
// сумма которых равна target, и верните их. Если пары нет — (-1, -1).
// Если пар несколько, верните пару с наименьшим j.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func twoSum(nums []int, target int) (int, int) {
	return 0, 0
}

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums         []int
		target, i, j int
	}{
		{[]int{2, 7, 11, 15}, 9, 0, 1},
		{[]int{3, 2, 4}, 6, 1, 2},
		{[]int{3, 3}, 6, 0, 1},
		{[]int{-1, -2, -3, -4}, -7, 2, 3},
		{[]int{1, 2}, 10, -1, -1},
		{nil, 0, -1, -1},
		{[]int{5}, 10, -1, -1},
		{[]int{1 << 40, 1 << 40, 3}, 1 << 41, 0, 1},
	}
	for _, c := range cases {
		if i, j := twoSum(c.nums, c.target); i != c.i || j != c.j {
			t.Errorf("twoSum(%v, %d) = (%d, %d), want (%d, %d)", c.nums, c.target, i, j, c.i, c.j)
		}
	}
}
