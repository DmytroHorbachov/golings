// anonymous_functions_x012: Первая заглавная
// Make the tests pass!
// I AM NOT DONE
//
// firstUpper возвращает индекс первой заглавной буквы через strings.IndexFunc.
// Тренирует: литерал-предикат.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
	"unicode"
)

func firstUpper(s string) int {
	return strings.IndexFunc(s, func(r rune) bool {
		return unicode.IsLower(r)
	})
}

func TestFirstUpper(t *testing.T) {
	if firstUpper("helloWorld") != 5 || firstUpper("abc") != -1 {
		t.Errorf("firstUpper works incorrectly")
	}
}
