// if33
// Make the tests pass!

// I AM NOT DONE
//
// oddSum должна вернуть true, если сумма двух чисел нечётна.
// Тренирует: условие с арифметическим выражением.
// Сложность: easy
package main_test

import "testing"

func oddSum(a, b int) bool {
	if (a+b)%2 == 0 {
		return true
	}
	return false
}

func TestOddSum(t *testing.T) {
	cases := [][3]int{{1, 2, 1}, {2, 2, 0}, {-3, 0, 1}, {5, 5, 0}}
	for _, c := range cases {
		if got := oddSum(c[0], c[1]); got != (c[2] == 1) {
			t.Errorf("oddSum(%d, %d) = %v", c[0], c[1], got)
		}
	}
}
