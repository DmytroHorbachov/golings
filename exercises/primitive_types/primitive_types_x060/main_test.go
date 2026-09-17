// primitive_types_x060: Только буквы
// Make the tests pass!
// I AM NOT DONE
//
// letters считает буквы любых алфавитов, пропуская цифры, пробелы и знаки.
// Тренирует: итерацию по рунам и unicode.IsLetter.
// Сложность: medium
package main_test

import (
	"testing"
	"unicode"
)

func letters(s string) int {
	n := 0
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			n++
		}
	}
	return n
}

func TestLetters(t *testing.T) {
	_ = unicode.IsLetter
	cases := map[string]int{"Go 1.22!": 2, "Привет, мир": 9, "": 0}
	for in, want := range cases {
		if got := letters(in); got != want {
			t.Errorf("letters(%q) = %d, want %d", in, got, want)
		}
	}
}
