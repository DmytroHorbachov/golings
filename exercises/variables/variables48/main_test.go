// variables48
// Make the tests pass!

// I AM NOT DONE
//
// Функция wrapIndex должна переводить любой индекс в диапазон [0, n).
// Для отрицательных индексов результат выходит за диапазон.
// Тренирует: знак результата оператора % в Go.
// Сложность: hard
package main_test

import "testing"

func wrapIndex(i, n int) int {
	return i % n
}

func TestWrapIndex(t *testing.T) {
	cases := []struct{ i, n, want int }{
		{0, 5, 0}, {7, 5, 2}, {-1, 5, 4}, {-5, 5, 0}, {-12, 5, 3},
	}
	for _, c := range cases {
		if got := wrapIndex(c.i, c.n); got != c.want {
			t.Errorf("wrapIndex(%d, %d) = %d, want %d", c.i, c.n, got, c.want)
		}
	}
}
