// variables61
// Make the tests pass!

// I AM NOT DONE
//
// Функция reverse должна развернуть строку с любыми символами Unicode.
// Для "привет" получается мусор.
// Тренирует: строка — это байты UTF-8, а символ — это rune.
// Сложность: hard
package main_test

import "testing"

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func TestReverse(t *testing.T) {
	cases := map[string]string{"go": "og", "привет": "тевирп", "": "", "аb": "bа"}
	for in, want := range cases {
		if got := reverse(in); got != want {
			t.Errorf("reverse(%q) = %q, want %q", in, got, want)
		}
	}
}
