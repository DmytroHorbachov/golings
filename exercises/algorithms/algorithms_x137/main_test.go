// algorithms_x137: Plus One (прибавить единицу)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: работа с числом-массивом. Цифры числа хранятся в срезе
// от старшей к младшей. Прибавьте 1 и верните результат.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) доп. памяти
package main_test

import (
	"reflect"
	"testing"
)

func plusOne(digits []int) []int {
	return nil
}

func TestPlusOne(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 9}, []int{4, 4, 0}},
		{[]int{9, 9}, []int{1, 0, 0}},
		{[]int{0}, []int{1}},
		{nil, []int{1}},
	}
	for _, c := range cases {
		if got := plusOne(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("plusOne(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
