// primitive_types_x083: Не только ASCII-цифры
// Make the tests pass!
// I AM NOT DONE
//
// onlyDigits проверяет, что строка — корректное десятичное число для strconv.Atoi.
// unicode.IsDigit пропускает цифры других письменностей, и Atoi затем падает.
// Тренирует: unicode.IsDigit шире, чем '0'..'9'.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
	"unicode"
)

func onlyDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

func TestOnlyDigits(t *testing.T) {
	_ = unicode.IsDigit
	for _, s := range []string{"123", "0"} {
		if !onlyDigits(s) {
			t.Errorf("onlyDigits(%q) = false", s)
		}
	}
	for _, s := range []string{"١٢٣", "१२", "12a", ""} {
		if onlyDigits(s) {
			if _, err := strconv.Atoi(s); err != nil {
				t.Errorf("onlyDigits(%q) = true, but Atoi fails: %v", s, err)
			} else {
				t.Errorf("onlyDigits(%q) = true", s)
			}
		}
	}
}
