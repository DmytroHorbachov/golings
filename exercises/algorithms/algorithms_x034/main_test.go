// algorithms_x034: Maximal Rectangle (прямоугольник из единиц)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: гистограммы по строкам и монотонный стек. В бинарной матрице найдите
// площадь наибольшего прямоугольника, состоящего только из '1'.
// Сложность: hard. Ожидаемая асимптотика: O(r·c) по времени, O(c) по памяти
package main_test

import "testing"

func maximalRectangle(m []string) int {
	return 0
}

func TestMaximalRectangle(t *testing.T) {
	cases := []struct {
		m    []string
		want int
	}{
		{[]string{"10100", "10111", "11111", "10010"}, 6},
		{[]string{"0"}, 0},
		{[]string{"1"}, 1},
		{nil, 0},
		{[]string{"111", "111"}, 6},
		{[]string{"01", "10"}, 1},
	}
	for _, c := range cases {
		if got := maximalRectangle(c.m); got != c.want {
			t.Errorf("maximalRectangle(%v) = %d, want %d", c.m, got, c.want)
		}
	}
}
