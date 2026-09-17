// if68
// Make the tests pass!

// I AM NOT DONE
//
// isAdult должна вернуть true для возраста 18 и старше.
// Тренирует: операторы сравнения в условии if.
// Сложность: easy
package main_test

import "testing"

func isAdult(age int) bool {
	if age > 18 {
		return true
	}
	return false
}

func TestIsAdult(t *testing.T) {
	cases := map[int]bool{17: false, 18: true, 30: true, 0: false}
	for in, want := range cases {
		if got := isAdult(in); got != want {
			t.Errorf("isAdult(%d) = %v, want %v", in, got, want)
		}
	}
}
