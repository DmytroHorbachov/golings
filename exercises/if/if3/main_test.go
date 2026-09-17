// if3
// Make the tests pass!

// I AM NOT DONE
//
// isUpper должна вернуть true для заглавных латинских букв.
// Тренирует: сравнение рун с диапазоном символов.
// Сложность: easy
package main_test

import "testing"

func isUpper(r rune) bool {
	if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func TestIsUpper(t *testing.T) {
	cases := map[rune]bool{'A': true, 'Z': true, 'a': false, '1': false, 'm': false}
	for in, want := range cases {
		if got := isUpper(in); got != want {
			t.Errorf("isUpper(%q) = %v, want %v", in, got, want)
		}
	}
}
