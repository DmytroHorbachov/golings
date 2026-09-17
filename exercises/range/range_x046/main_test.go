// range_x046: Подсчёт слов по переходам
// Make the tests pass!
// I AM NOT DONE
//
// countWords считает слова как переходы «разделитель -> не разделитель».
// Тренирует: range по строке с флагом состояния.
// Сложность: medium
package main_test

import (
	"testing"
	"unicode"
)

func countWords(s string) int {
	n := 0
	inWord := false
	for _, r := range s {
		if unicode.IsSpace(r) {
			n++
		}
	}
	return n
}

func TestCountWords(t *testing.T) {
	cases := map[string]int{"hello world": 2, "  lots   of\tspace ": 3, "": 0, "one": 1}
	for in, want := range cases {
		if got := countWords(in); got != want {
			t.Errorf("countWords(%q) = %d, want %d", in, got, want)
		}
	}
}
