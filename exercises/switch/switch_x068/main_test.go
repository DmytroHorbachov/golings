// switch_x068: break внутри switch
// Make the tests pass!
// I AM NOT DONE
//
// firstNegative должна вернуть индекс первого отрицательного числа.
// break в switch не останавливает цикл, и находится последнее отрицательное.
// Тренирует: break внутри switch выходит только из switch.
// Сложность: hard
package main_test

import "testing"

func firstNegative(nums []int) int {
	idx := -1
	for i, n := range nums {
		switch {
		case n < 0:
			idx = i
			break
		}
	}
	return idx
}

func TestFirstNegative(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{{[]int{3, -1, 4, -5}, 1}, {[]int{1, 2}, -1}, {[]int{-7, -8}, 0}}
	for _, c := range cases {
		if got := firstNegative(c.in); got != c.want {
			t.Errorf("firstNegative(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
