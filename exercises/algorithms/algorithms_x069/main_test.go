// algorithms_x069: Longest Common Prefix (общий префикс)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: посимвольное сравнение. Верните самый длинный общий префикс
// всех строк; если его нет — пустую строку.
// Сложность: easy. Ожидаемая асимптотика: O(S) по времени (S — сумма длин), O(1) по памяти
package main_test

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	return ""
}

func TestLongestCommonPrefix(t *testing.T) {
	_ = strings.HasPrefix
	cases := []struct {
		in   []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"same", "same"}, "same"},
		{[]string{"a"}, "a"},
		{nil, ""},
		{[]string{"", "abc"}, ""},
	}
	for _, c := range cases {
		if got := longestCommonPrefix(c.in); got != c.want {
			t.Errorf("longestCommonPrefix(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
