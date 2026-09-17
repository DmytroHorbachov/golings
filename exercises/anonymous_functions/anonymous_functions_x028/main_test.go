// anonymous_functions_x028: Последняя цифра
// Make the tests pass!
// I AM NOT DONE
//
// lastDigit возвращает индекс последней цифры через strings.LastIndexFunc.
// Тренирует: литерал для функций поиска с конца.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func lastDigit(s string) int {
	return strings.LastIndexFunc(s, func(r rune) bool {
		return unicode.IsLetter(r)
	})
}

func TestLastDigit(t *testing.T) {
	if lastDigit("a1b2c") != 3 || lastDigit("abc") != -1 {
		t.Errorf("lastDigit works incorrectly")
	}
}
