// algorithms_x043: Split Array Largest Sum (разбиение на k частей)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двоичный поиск по ответу и жадная проверка. Разбейте срез
// неотрицательных чисел на k непустых непрерывных частей так, чтобы
// наибольшая сумма части была минимальной. Верните эту сумму.
// Сложность: hard. Ожидаемая асимптотика: O(n·log S) по времени, O(1) по памяти
package main_test

import "testing"

func splitArray(nums []int, k int) int {
	return 0
}

func TestSplitArray(t *testing.T) {
	cases := []struct {
		nums    []int
		k, want int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},
		{[]int{5}, 1, 5},
		{[]int{0, 0, 0}, 2, 0},
		{[]int{1000000000, 1000000000}, 1, 2000000000},
	}
	for _, c := range cases {
		if got := splitArray(c.nums, c.k); got != c.want {
			t.Errorf("splitArray(%v, %d) = %d, want %d", c.nums, c.k, got, c.want)
		}
	}
}
