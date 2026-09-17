// anonymous_functions75
// Make the tests pass!

// I AM NOT DONE
//
// noDigits удаляет цифры из строки с помощью strings.Map.
// Тренирует: литерал, возвращающий -1 для удаления символа.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func noDigits(s string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return ' '
		}
		return r
	}, s)
}

func TestNoDigits(t *testing.T) {
	if got := noDigits("a1b22c"); got != "abc" {
		t.Errorf("noDigits = %q", got)
	}
}
