// algorithms_x031: Decode String (распаковка строки)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: стек. Строка вида "3[a2[c]]" раскрывается в "accaccacc".
// Числа — положительные, скобки корректны, внутри — латинские буквы.
// Сложность: medium. Ожидаемая асимптотика: O(n·k) по времени, O(n) по памяти
package main_test

import (
	"strings"
	"testing"
)

func decodeString(s string) string {
	return ""
}

func TestDecodeString(t *testing.T) {
	_ = strings.Repeat
	cases := map[string]string{
		"3[a]2[bc]":     "aaabcbc",
		"3[a2[c]]":      "accaccacc",
		"2[abc]3[cd]ef": "abcabccdcdcdef",
		"abc":           "abc",
		"":              "",
		"10[x]":         "xxxxxxxxxx",
	}
	for in, want := range cases {
		if got := decodeString(in); got != want {
			t.Errorf("decodeString(%q) = %q, want %q", in, got, want)
		}
	}
}
