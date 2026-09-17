// switch_x014: Класс символа
// Make the tests pass!
// I AM NOT DONE
//
// charClass возвращает "digit", "letter" или "other".
// Тренирует: switch без тега с диапазонами символов.
// Сложность: easy
package main_test

import "testing"

func charClass(c byte) string {
	switch {
	case c >= '0' && c <= '9':
		return "digit"
	case c >= 'a' && c <= 'z':
		return "letter"
	default:
		return "other"
	}
}

func TestCharClass(t *testing.T) {
	cases := map[byte]string{'5': "digit", 'q': "letter", 'Q': "letter", '-': "other"}
	for in, want := range cases {
		if got := charClass(in); got != want {
			t.Errorf("charClass(%c) = %s, want %s", in, got, want)
		}
	}
}
