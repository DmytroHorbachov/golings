// algorithms33
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: битовые операции. Каждое число встречается дважды, кроме одного.
// Найдите его, не используя дополнительную память.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func singleNumber(nums []int) int {
	return 0
}

func TestSingleNumber(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{7}, 7},
		{[]int{-3, 5, 5}, -3},
		{nil, 0},
	}
	for _, c := range cases {
		if got := singleNumber(c.nums); got != c.want {
			t.Errorf("singleNumber(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
