// if72
// Make the tests pass!

// I AM NOT DONE
//
// inside должна проверить, лежит ли точка (x, y) внутри прямоугольника
// от (0, 0) до (w, h), включая границы.
// Тренирует: составные условия по двум координатам.
// Сложность: easy
package main_test

import "testing"

func inside(x, y, w, h int) bool {
	if x >= 0 && x <= w && y >= 0 && y <= w {
		return true
	}
	return false
}

func TestInside(t *testing.T) {
	cases := []struct {
		x, y int
		want bool
	}{{0, 0, true}, {10, 5, true}, {5, 6, false}, {-1, 2, false}, {3, 5, true}}
	for _, c := range cases {
		if got := inside(c.x, c.y, 10, 5); got != c.want {
			t.Errorf("inside(%d, %d) = %v, want %v", c.x, c.y, got, c.want)
		}
	}
}
