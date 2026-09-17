// algorithms_x011: Move Zeroes (нули в конец)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: два указателя. Переместите все нули в конец среза на месте,
// сохранив порядок остальных элементов.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func moveZeroes(nums []int) {
}

func TestMoveZeroes(t *testing.T) {
	cases := []struct {
		nums, want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{0, 0, -1}, []int{-1, 0, 0}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		got := append([]int{}, c.nums...)
		moveZeroes(got)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("moveZeroes(%v) = %v, want %v", c.nums, got, c.want)
		}
	}
}
