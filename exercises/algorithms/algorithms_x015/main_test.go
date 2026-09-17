// algorithms_x015: Container With Most Water (самый вместительный контейнер)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: два указателя. Высоты стенок заданы в срезе. Найдите максимальную
// площадь воды между двумя стенками (ширина × меньшая высота).
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func maxArea(h []int) int {
	return 0
}

func TestMaxArea(t *testing.T) {
	cases := []struct {
		h    []int
		want int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{5}, 0},
		{nil, 0},
		{[]int{1000000, 1, 1000000}, 2000000},
	}
	for _, c := range cases {
		if got := maxArea(c.h); got != c.want {
			t.Errorf("maxArea(%v) = %d, want %d", c.h, got, c.want)
		}
	}
}
