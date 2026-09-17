// range90
// Make the tests pass!

// I AM NOT DONE
//
// missing находит единственное пропущенное число в перестановке 0..n без одного элемента.
// Тренирует: range с индексом и суммированием.
// Сложность: medium
package main_test

import "testing"

func missing(nums []int) int {
	result := 1
	for i, v := range nums {
		result += i - v
	}
	return result
}

func TestMissing(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{3, 0, 1}, 2}, {[]int{0, 1}, 2}, {[]int{1}, 0}, {nil, 0}}
	for _, c := range cases {
		if got := missing(c.in); got != c.want {
			t.Errorf("missing(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
