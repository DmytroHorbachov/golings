// range_x055: Палиндром по рунам
// Make the tests pass!
// I AM NOT DONE
//
// isPalindrome проверяет строку без учёта регистра, сравнивая руны с двух концов.
// Тренирует: range по []rune с индексом.
// Сложность: medium
package main_test

import (
	"testing"
	"unicode"
)

func isPalindrome(s string) bool {
	rs := []rune(s)
	for i, r := range rs {
		if r != rs[len(rs)-i] {
			return false
		}
	}
	return true
}

func TestIsPalindrome(t *testing.T) {
	_ = unicode.ToLower
	cases := map[string]bool{"Шалаш": true, "Level": true, "go": false, "": true}
	for in, want := range cases {
		if got := isPalindrome(in); got != want {
			t.Errorf("isPalindrome(%q) = %v, want %v", in, got, want)
		}
	}
}
