// if92
// Make the tests pass!

// I AM NOT DONE
//
// isBlank должна вернуть true для строки, состоящей только из пробелов, или пустой.
// Тренирует: условие с вызовом функции.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func isBlank(s string) bool {
	if strings.TrimSpace(s) != "" {
		return true
	}
	return false
}

func TestIsBlank(t *testing.T) {
	cases := map[string]bool{"": true, "   ": true, " go ": false, "x": false}
	for in, want := range cases {
		if got := isBlank(in); got != want {
			t.Errorf("isBlank(%q) = %v, want %v", in, got, want)
		}
	}
}
