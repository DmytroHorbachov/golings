// algorithms_x019: Squares of a Sorted Array (квадраты отсортированного)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: два указателя. Срез отсортирован и может содержать отрицательные числа.
// Верните отсортированные квадраты элементов без повторной сортировки.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func sortedSquares(nums []int) []int {
	return nil
}

func TestSortedSquares(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
		{[]int{}, []int{}},
		{[]int{-5}, []int{25}},
		{[]int{-3, -2, -1}, []int{1, 4, 9}},
	}
	for _, c := range cases {
		if got := sortedSquares(c.nums); !reflect.DeepEqual(got, c.want) {
			t.Errorf("sortedSquares(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
