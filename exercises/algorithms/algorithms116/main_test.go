// algorithms116
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: жадность по «границе достижимости». nums[i] — максимальная длина
// прыжка из позиции i. Верните минимальное число прыжков до конца
// (достижимость гарантирована).
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func jump(nums []int) int {
	return 0
}

func TestJump(t *testing.T) {
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, 1, 1, 4}, 2},
		{[]int{2, 3, 0, 1, 4}, 2},
		{[]int{0}, 0},
		{nil, 0},
		{[]int{1, 1, 1, 1}, 3},
		{[]int{10, 1, 1}, 1},
	}
	for _, c := range cases {
		if got := jump(c.nums); got != c.want {
			t.Errorf("jump(%v) = %d, want %d", c.nums, got, c.want)
		}
	}
}
