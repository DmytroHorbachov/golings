// algorithms145
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: использование массива как хэш-таблицы. Найдите наименьшее
// положительное целое, которого нет в срезе. Разрешается переставлять элементы.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(1) доп. памяти
package main_test

import "testing"

func firstMissingPositive(nums []int) int {
	return 0
}

func TestFirstMissingPositive(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 0}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
		{nil, 1},
		{[]int{1}, 2},
		{[]int{2, 2, 2}, 1},
		{[]int{1 << 62, -1 << 62, 1}, 2},
	}
	for _, c := range cases {
		in := append([]int(nil), c.nums...)
		if got := firstMissingPositive(in); got != c.want {
			t.Errorf("firstMissingPositive(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
