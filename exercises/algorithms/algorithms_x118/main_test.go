// algorithms_x118: Regular Expression Matching (сопоставление с шаблоном)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: двумерное ДП. Шаблон содержит '.' (любой символ) и '*' (ноль или
// более повторений предыдущего символа). Шаблон должен покрывать всю строку.
// Сложность: hard. Ожидаемая асимптотика: O(n·m) по времени, O(n·m) по памяти
package main_test

import "testing"

func isMatch(s, p string) bool {
	return false
}

func TestIsMatch(t *testing.T) {
	cases := []struct {
		s, p string
		want bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"", ".*", true},
		{"", "", true},
		{"abc", "", false},
	}
	for _, c := range cases {
		if got := isMatch(c.s, c.p); got != c.want {
			t.Errorf("isMatch(%q, %q) = %v, want %v", c.s, c.p, got, c.want)
		}
	}
}
