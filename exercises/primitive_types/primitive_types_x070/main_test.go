// primitive_types_x070: Обрезка по символам
// Make the tests pass!
// I AM NOT DONE
//
// truncate должна оставить не более n символов строки и добавить "…", если строка
// была длиннее. Сейчас кириллица превращается в «битые» символы.
// Тренирует: срез строки режет по байтам и может разрезать символ UTF-8.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode/utf8"
)

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "…"
}

func TestTruncate(t *testing.T) {
	cases := []struct {
		in   string
		n    int
		want string
	}{{"привет мир", 4, "прив…"}, {"go", 4, "go"}, {"hello!", 5, "hello…"}, {"мир", 3, "мир"}}
	for _, c := range cases {
		got := truncate(c.in, c.n)
		if got != c.want || !utf8.ValidString(got) {
			t.Errorf("truncate(%q, %d) = %q, want %q", c.in, c.n, got, c.want)
		}
	}
}
