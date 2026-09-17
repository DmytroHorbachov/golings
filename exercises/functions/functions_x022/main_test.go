// functions_x022: Возведение в степень рекурсивно
// Make the tests pass!
// I AM NOT DONE
//
// pow(x, n) должна вычислять x^n для n >= 0 рекурсивно.
// Тренирует: рекурсию с уменьшением аргумента.
// Сложность: easy
package main_test

import "testing"

func pow(x, n int) int {
	if n == 0 {
		return 1
	}
	return x * pow(x-1, n)
}

func TestPow(t *testing.T) {
	cases := []struct{ x, n, want int }{{2, 10, 1024}, {3, 0, 1}, {5, 3, 125}, {-2, 3, -8}}
	for _, c := range cases {
		if got := pow(c.x, c.n); got != c.want {
			t.Errorf("pow(%d, %d) = %d, want %d", c.x, c.n, got, c.want)
		}
	}
}
