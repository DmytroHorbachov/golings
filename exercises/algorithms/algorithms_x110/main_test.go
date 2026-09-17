// algorithms_x110: Russian Doll Envelopes (вложенные конверты)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: сортировка и наибольшая возрастающая подпоследовательность.
// Конверт [w, h] вкладывается в другой, если строго меньше по обеим сторонам.
// Верните максимальное число вложенных друг в друга конвертов.
// Сложность: hard. Ожидаемая асимптотика: O(n·log n) по времени, O(n) по памяти
package main_test

import (
	"sort"
	"testing"
)

func maxEnvelopes(envelopes [][2]int) int {
	_ = sort.SearchInts
	return 0
}

func TestMaxEnvelopes(t *testing.T) {
	cases := []struct {
		env  [][2]int
		want int
	}{
		{[][2]int{{5, 4}, {6, 4}, {6, 7}, {2, 3}}, 3},
		{[][2]int{{1, 1}, {1, 1}, {1, 1}}, 1},
		{nil, 0},
		{[][2]int{{4, 5}}, 1},
		{[][2]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}, 4},
	}
	for _, c := range cases {
		if got := maxEnvelopes(c.env); got != c.want {
			t.Errorf("maxEnvelopes(%v) = %d, want %d", c.env, got, c.want)
		}
	}
}
