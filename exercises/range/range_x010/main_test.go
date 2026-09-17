// range_x010: Разворот строки
// Make the tests pass!
// I AM NOT DONE
//
// reverse разворачивает строку, добавляя каждую руну в начало результата.
// Тренирует: range по рунам строки.
// Сложность: easy
package main_test

import "testing"

func reverse(s string) string {
	out := ""
	for _, r := range s {
		out = out + string(r)
	}
	return out
}

func TestReverse(t *testing.T) {
	cases := map[string]string{"abc": "cba", "мир": "рим", "": ""}
	for in, want := range cases {
		if got := reverse(in); got != want {
			t.Errorf("reverse(%q) = %q, want %q", in, got, want)
		}
	}
}
