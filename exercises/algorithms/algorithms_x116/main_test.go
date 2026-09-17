// algorithms_x116: Longest Palindromic Substring (наибольшая подстрока-палиндром)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: расширение от центра. Найдите самую длинную подстроку-палиндром
// (в байтах); при нескольких одинаковой длины — самую левую.
// Сложность: medium. Ожидаемая асимптотика: O(n²) по времени, O(1) по памяти
package main_test

import "testing"

func longestPalindrome(s string) string {
	return ""
}

func TestLongestPalindrome(t *testing.T) {
	cases := map[string]string{
		"babad": "bab",
		"cbbd":  "bb",
		"a":     "a",
		"":      "",
		"ac":    "a",
		"aaaa":  "aaaa",
	}
	for in, want := range cases {
		if got := longestPalindrome(in); got != want {
			t.Errorf("longestPalindrome(%q) = %q, want %q", in, got, want)
		}
	}
}
