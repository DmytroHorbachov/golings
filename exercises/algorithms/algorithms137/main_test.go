// algorithms137
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: голосование Бойера–Мура. Гарантируется, что в непустом срезе есть
// элемент, встречающийся больше n/2 раз. Найдите его без дополнительной памяти.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func majorityElement(nums []int) int {
	return 0
}

func TestMajorityElement(t *testing.T) {
	big := make([]int, 100001)
	for i := range big {
		if i%2 == 0 {
			big[i] = -7
		} else {
			big[i] = i
		}
	}
	cases := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 3}, 3},
		{[]int{2, 2, 1, 1, 1, 2, 2}, 2},
		{[]int{9}, 9},
		{big, -7},
	}
	for _, c := range cases {
		if got := majorityElement(c.nums); got != c.want {
			t.Errorf("majorityElement(len=%d) = %d, want %d", len(c.nums), got, c.want)
		}
	}
}
