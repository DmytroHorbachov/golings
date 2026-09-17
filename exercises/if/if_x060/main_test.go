// if_x060: HEX-цвет
// Make the tests pass!
// I AM NOT DONE
//
// validHex проверяет строку вида "#1a2B3c": решётка и ровно 6 шестнадцатеричных цифр.
// Тренирует: условие на длину и посимвольные проверки.
// Сложность: medium
package main_test

import "testing"

func validHex(s string) bool {
	if s[0] != '#' {
		return false
	}
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			continue
		}
		return false
	}
	return true
}

func TestValidHex(t *testing.T) {
	cases := map[string]bool{"#1a2B3c": true, "#000000": true, "123456": false, "#12345": false, "#12345g": false, "": false}
	for in, want := range cases {
		if got := validHex(in); got != want {
			t.Errorf("validHex(%q) = %v, want %v", in, got, want)
		}
	}
}
