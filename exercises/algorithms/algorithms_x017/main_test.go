// algorithms_x017: Remove Duplicates from Sorted Array (удаление повторов)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: два указателя. В отсортированном срезе оставьте на месте только
// уникальные значения в начале и верните их количество.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func removeDuplicates(nums []int) int {
	return 0
}

func TestRemoveDuplicates(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{1, 1, 2}, []int{1, 2}},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, []int{0, 1, 2, 3, 4}},
		{nil, []int{}},
		{[]int{7}, []int{7}},
		{[]int{-3, -3, -3}, []int{-3}},
	}
	for _, c := range cases {
		in := append([]int{}, c.nums...)
		n := removeDuplicates(in)
		if !reflect.DeepEqual(in[:n], c.want) {
			t.Errorf("removeDuplicates(%v) -> %v, want %v", c.nums, in[:n], c.want)
		}
	}
}
