// primitive_types_x013: Заглавная ли буква
// Make the tests pass!
// I AM NOT DONE
//
// startsUpper должна вернуть true, если первая руна — заглавная буква (любой алфавит).
// Тренирует: пакет unicode.
// Сложность: easy
package main_test

import (
	"testing"
	"unicode"
)

func startsUpper(s string) bool {
	for _, r := range s {
		return unicode.IsLower(r)
	}
	return false
}

func TestStartsUpper(t *testing.T) {
	cases := map[string]bool{"Go": true, "Ёж": true, "go": false, "1st": false, "": false}
	for in, want := range cases {
		if got := startsUpper(in); got != want {
			t.Errorf("startsUpper(%q) = %v, want %v", in, got, want)
		}
	}
}
