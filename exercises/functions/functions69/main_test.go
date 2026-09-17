// functions69
// Make the tests pass!

// I AM NOT DONE
//
// paths(r, c) должна посчитать число путей из левого верхнего угла сетки r×c
// в правый нижний, двигаясь только вправо и вниз.
// Тренирует: рекурсию с двумя ветвями и базовыми случаями.
// Сложность: medium
package main_test

import "testing"

func paths(r, c int) int {
	if r == 0 || c == 0 {
		return 1
	}
	return paths(r-1, c) * paths(r, c-1)
}

func TestPaths(t *testing.T) {
	cases := []struct{ r, c, want int }{{1, 1, 1}, {2, 2, 2}, {3, 3, 6}, {3, 7, 28}}
	for _, cs := range cases {
		if got := paths(cs.r, cs.c); got != cs.want {
			t.Errorf("paths(%d, %d) = %d, want %d", cs.r, cs.c, got, cs.want)
		}
	}
}
