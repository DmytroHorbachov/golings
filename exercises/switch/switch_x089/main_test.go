// switch_x089: Кириллица и байт
// Make the tests pass!
// I AM NOT DONE
//
// greetingLang определяет язык по первой букве: 'п' — "ru", 'h' — "en".
// Код не компилируется: руна 'п' не помещается в byte.
// Тренирует: s[0] — это байт, а case-константы должны помещаться в тип тега.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode/utf8"
)

func greetingLang(s string) string {
	if s == "" {
		return "unknown"
	}
	switch s[0] {
	case 'п':
		return "ru"
	case 'h':
		return "en"
	}
	return "unknown"
}

func TestGreetingLang(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]string{"привет": "ru", "hello": "en", "hola": "en", "bonjour": "unknown", "": "unknown"}
	for in, want := range cases {
		if got := greetingLang(in); got != want {
			t.Errorf("greetingLang(%q) = %s, want %s", in, got, want)
		}
	}
}
