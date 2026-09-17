// algorithms118
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: двоичный поиск по разбиению. Найдите медиану объединения двух
// отсортированных срезов (хотя бы один непустой).
// Сложность: hard. Ожидаемая асимптотика: O(log min(m, n)) по времени, O(1) по памяти
package main_test

import (
	"math"
	"testing"
)

func findMedian(a, b []int) float64 {
	return 0
}

func TestFindMedian(t *testing.T) {
	_ = math.MaxInt64
	cases := []struct {
		a, b []int
		want float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
		{nil, []int{1}, 1},
		{[]int{2}, nil, 2},
		{[]int{1, 1, 1}, []int{1, 1}, 1},
		{[]int{-5, 3, 6, 12, 15}, []int{-12, -10, -6, -3, 4, 10}, 3},
	}
	for _, c := range cases {
		if got := findMedian(c.a, c.b); got != c.want {
			t.Errorf("findMedian(%v, %v) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
