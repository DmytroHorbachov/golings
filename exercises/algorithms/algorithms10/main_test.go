// algorithms10
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: монотонный стек. Для каждого элемента верните первый элемент
// правее него, который больше; если такого нет — -1.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func nextGreater(nums []int) []int {
	return nil
}

func TestNextGreater(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{2, 1, 2, 4, 3}, []int{4, 2, 4, -1, -1}},
		{[]int{5, 4, 3}, []int{-1, -1, -1}},
		{[]int{1}, []int{-1}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := nextGreater(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("nextGreater(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
