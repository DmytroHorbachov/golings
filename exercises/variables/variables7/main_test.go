// variables7
// Make the tests pass!

// I AM NOT DONE
//
// Функция minValue должна вернуть минимальный элемент непустого среза.
// Для положительных чисел всегда возвращается 0.
// Тренирует: правильный выбор начального значения переменной.
// Сложность: easy
package main_test

import "testing"

func minValue(nums []int) int {
	var m int
	for _, v := range nums {
		if v < m {
			m = v
		}
	}
	return m
}

func TestMinValue(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{5, 3, 8}, 3}, {[]int{-1, -7}, -7}, {[]int{42}, 42}}
	for _, c := range cases {
		if got := minValue(c.in); got != c.want {
			t.Errorf("minValue(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
