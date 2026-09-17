// if52
// Make the tests pass!

// I AM NOT DONE
//
// isTriangle должна проверить неравенство треугольника для сторон a, b, c.
// Тренирует: составные условия из нескольких сравнений.
// Сложность: easy
package main_test

import "testing"

func isTriangle(a, b, c int) bool {
	if a < b+c || b < a+c || c < a+b {
		return true
	}
	return false
}

func TestIsTriangle(t *testing.T) {
	cases := []struct {
		a, b, c int
		want    bool
	}{{3, 4, 5, true}, {1, 2, 3, false}, {1, 1, 10, false}, {2, 2, 2, true}}
	for _, c := range cases {
		if got := isTriangle(c.a, c.b, c.c); got != c.want {
			t.Errorf("isTriangle(%d,%d,%d) = %v, want %v", c.a, c.b, c.c, got, c.want)
		}
	}
}
