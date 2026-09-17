// algorithms_x020: Is Subsequence (подпоследовательность)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: два указателя. Проверьте, можно ли получить s из t, удалив
// некоторые символы и не меняя порядок остальных.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func isSubsequence(s, t string) bool {
	return false
}

func TestIsSubsequence(t *testing.T) {
	cases := []struct {
		s, t string
		want bool
	}{
		{"abc", "ahbgdc", true},
		{"axc", "ahbgdc", false},
		{"", "abc", true},
		{"a", "", false},
		{"", "", true},
		{"aaa", "aa", false},
	}
	for _, c := range cases {
		if got := isSubsequence(c.s, c.t); got != c.want {
			t.Errorf("isSubsequence(%q, %q) = %v, want %v", c.s, c.t, got, c.want)
		}
	}
}
