// if77
// Make the tests pass!

// I AM NOT DONE
//
// validNick разрешает ники длиной от 3 до 10 символов (не байт).
// Кириллический ник из 6 букв сейчас отклоняется.
// Тренирует: len(string) считает байты, а не символы.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode/utf8"
)

func validNick(nick string) bool {
	n := len(nick)
	if n < 3 || n > 10 {
		return false
	}
	return true
}

func TestValidNick(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]bool{"gopher": true, "гофер": true, "ёжик_42": true, "ab": false, "очень_длинный": false}
	for in, want := range cases {
		if got := validNick(in); got != want {
			t.Errorf("validNick(%q) = %v, want %v", in, got, want)
		}
	}
}
