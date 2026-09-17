// algorithms124
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: сортировка интервалов. Определите, можно ли посетить все встречи
// [начало, конец), то есть нет ли пересечений.
// Сложность: easy. Ожидаемая асимптотика: O(n·log n) по времени, O(n) по памяти
package main_test

import (
	"sort"
	"testing"
)

func canAttendMeetings(intervals [][2]int) bool {
	_ = sort.Ints
	return false
}

func TestCanAttendMeetings(t *testing.T) {
	cases := []struct {
		in   [][2]int
		want bool
	}{
		{[][2]int{{0, 30}, {5, 10}, {15, 20}}, false},
		{[][2]int{{7, 10}, {2, 4}}, true},
		{nil, true},
		{[][2]int{{1, 2}, {2, 3}}, true},
		{[][2]int{{1, 5}, {4, 6}}, false},
	}
	for _, c := range cases {
		if got := canAttendMeetings(c.in); got != c.want {
			t.Errorf("canAttendMeetings(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
