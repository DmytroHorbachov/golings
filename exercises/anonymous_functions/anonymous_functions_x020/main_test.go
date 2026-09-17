// anonymous_functions_x020: Обрезка по условию
// Make the tests pass!
// I AM NOT DONE
//
// trimPunct убирает знаки препинания по краям строки.
// Тренирует: литерал для strings.TrimFunc.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func trimPunct(s string) string {
	return strings.TrimFunc(s, func(r rune) bool {
		return !unicode.IsPunct(r)
	})
}

func TestTrimPunct(t *testing.T) {
	if got := trimPunct("...hi, there!!"); got != "hi, there" {
		t.Errorf("trimPunct = %q", got)
	}
}
