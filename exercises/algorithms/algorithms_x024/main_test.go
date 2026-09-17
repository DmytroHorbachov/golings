// algorithms_x024: Sliding Window Maximum (максимумы окон)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: монотонная очередь (дек). Для каждого окна размера k верните
// его максимум. Решение должно работать за линейное время.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(k) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func maxSlidingWindow(nums []int, k int) []int {
	return nil
}

func TestMaxSlidingWindow(t *testing.T) {
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{9, 8, 7, 6}, 2, []int{9, 8, 7}},
		{[]int{1, 2}, 3, nil},
	}
	for _, c := range cases {
		if got := maxSlidingWindow(c.nums, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("maxSlidingWindow(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
	big := make([]int, 200000)
	for i := range big {
		big[i] = i
	}
	got := maxSlidingWindow(big, 1000)
	if len(got) != len(big)-999 || got[len(got)-1] != len(big)-1 {
		t.Errorf("big window result is wrong")
	}
}
