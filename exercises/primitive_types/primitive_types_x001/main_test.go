// primitive_types_x001: Булево выражение
// Make the tests pass!
// I AM NOT DONE
//
// isPositive должна вернуть результат сравнения напрямую.
// Тренирует: тип bool как результат выражения сравнения.
// Сложность: easy
package main_test

import "testing"

func isPositive(n int) bool {
	return n >= 0
}

func TestIsPositive(t *testing.T) {
	cases := map[int]bool{1: true, 0: false, -1: false}
	for in, want := range cases {
		if got := isPositive(in); got != want {
			t.Errorf("isPositive(%d) = %v, want %v", in, got, want)
		}
	}
}
