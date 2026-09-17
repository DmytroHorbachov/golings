// algorithms87
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: скользящее окно со счётчиками. Найдите самую короткую подстроку s,
// содержащую все символы t (с учётом кратности). Если такой нет — "".
// При нескольких минимальных окнах верните самое левое.
// Сложность: hard. Ожидаемая асимптотика: O(n + m) по времени, O(k) по памяти
package main_test

import "testing"

func minWindow(s, t string) string {
	return ""
}

func TestMinWindow(t *testing.T) {
	cases := []struct{ s, t, want string }{
		{"ADOBECODEBANC", "ABC", "BANC"},
		{"a", "a", "a"},
		{"a", "aa", ""},
		{"", "a", ""},
		{"abc", "", ""},
		{"aaflslflsldkalskaaa", "aaa", "aaa"},
		{"bba", "ab", "ba"},
	}
	for _, c := range cases {
		if got := minWindow(c.s, c.t); got != c.want {
			t.Errorf("minWindow(%q, %q) = %q, want %q", c.s, c.t, got, c.want)
		}
	}
}
