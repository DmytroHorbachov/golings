// algorithms13
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: подсчёт частот. Проверьте, что строки состоят из одних и тех же
// символов Unicode в том же количестве.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(k) по памяти (k — размер алфавита)
package main_test

import "testing"

func isAnagram(s, t string) bool {
	return false
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
		{"", "", true},
		{"a", "", false},
		{"ρόδο", "δόρο", true},
		{"aab", "abb", false},
	}
	for _, c := range cases {
		if got := isAnagram(c.s, c.t); got != c.want {
			t.Errorf("isAnagram(%q, %q) = %v, want %v", c.s, c.t, got, c.want)
		}
	}
}
