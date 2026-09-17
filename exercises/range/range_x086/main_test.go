// range_x086: Руна в []byte
// Make the tests pass!
// I AM NOT DONE
//
// onlyLetters собирает буквы строки в []byte. Для кириллицы вместо букв
// получается мусор.
// Тренирует: byte(r) обрезает руну до одного байта.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode"
	"unicode/utf8"
)

func onlyLetters(s string) []byte {
	var b []byte
	for _, r := range s {
		if unicode.IsLetter(r) {
			b = append(b, byte(r))
		}
	}
	return b
}

func TestOnlyLetters(t *testing.T) {
	_ = utf8.RuneLen
	if got := string(onlyLetters("Go-1 и Ё!")); got != "GoиЁ" {
		t.Errorf("onlyLetters = %q, want GoиЁ", got)
	}
}
