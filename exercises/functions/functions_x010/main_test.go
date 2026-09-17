// functions_x010: strings.Map
// Make the tests pass!
// I AM NOT DONE
//
// Функция rot13 должна кодировать латинские буквы, остальные символы не трогать.
// Функция-отображение передаётся в strings.Map, но передана не та.
// Тренирует: функции высшего порядка из стандартной библиотеки.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func rotRune(r rune) rune {
	if !unicode.IsLetter(r) {
		return r
	}
	switch {
	case r >= 'a' && r <= 'z':
		return 'a' + (r-'a'+13)%26
	case r >= 'A' && r <= 'Z':
		return 'A' + (r-'A'+13)%26
	}
	return r
}

func rot13(s string) string {
	return strings.Map(unicode.ToUpper, s)
}

func TestRot13(t *testing.T) {
	if got := rot13("Hello, Go!"); got != "Uryyb, Tb!" {
		t.Errorf("rot13(Hello, Go!) = %q, want %q", got, "Uryyb, Tb!")
	}
}
