// functions_x027: Обрезка по функции
// Make the tests pass!
// I AM NOT DONE
//
// cleanTag должна убрать по краям строки все символы, кроме букв и цифр.
// Тренирует: strings.TrimFunc с собственным предикатом.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func notAlnum(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func cleanTag(s string) string {
	return strings.TrimFunc(s, notAlnum)
}

func TestCleanTag(t *testing.T) {
	cases := map[string]string{"  #golang! ": "golang", "go1.22": "go1.22", "--": ""}
	for in, want := range cases {
		if got := cleanTag(in); got != want {
			t.Errorf("cleanTag(%q) = %q, want %q", in, got, want)
		}
	}
}
