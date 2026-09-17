// algorithms108
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: монотонный стек. Найдите площадь наибольшего прямоугольника,
// который помещается в гистограмму со столбцами ширины 1.
// Сложность: hard. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func largestRectangle(h []int) int {
	return 0
}

func TestLargestRectangle(t *testing.T) {
	cases := []struct {
		h    []int
		want int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
		{nil, 0},
		{[]int{7}, 7},
		{[]int{3, 3, 3, 3}, 12},
		{[]int{1, 2, 3, 4, 5}, 9},
	}
	for _, c := range cases {
		if got := largestRectangle(c.h); got != c.want {
			t.Errorf("largestRectangle(%v) = %d, want %d", c.h, got, c.want)
		}
	}
}
