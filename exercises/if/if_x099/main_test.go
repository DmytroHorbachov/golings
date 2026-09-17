// if_x099: Первый символ — байт или руна
// Make the tests pass!
// I AM NOT DONE
//
// startsWithE проверяет, что строка начинается с буквы 'é'.
// Сравнение первого байта никогда не срабатывает.
// Тренирует: s[0] — байт UTF-8, а не символ.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode/utf8"
)

func startsWithE(s string) bool {
	if s == "" {
		return false
	}
	if s[0] == 'é' {
		return true
	}
	return false
}

func TestStartsWithE(t *testing.T) {
	_ = utf8.RuneLen
	cases := map[string]bool{"école": true, "été": true, "ecole": false, "": false, "Ã©": false}
	for in, want := range cases {
		if got := startsWithE(in); got != want {
			t.Errorf("startsWithE(%q) = %v, want %v", in, got, want)
		}
	}
}
