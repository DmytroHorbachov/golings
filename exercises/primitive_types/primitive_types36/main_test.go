// primitive_types36
// Make the tests pass!

// I AM NOT DONE
//
// hasGo должна вернуть true, если в строке встречается "go".
// Тренирует: strings.Contains.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func hasGo(s string) bool {
	return strings.Contains("go", s)
}

func TestHasGo(t *testing.T) {
	cases := map[string]bool{"golang": true, "ergo": true, "rust": false}
	for in, want := range cases {
		if got := hasGo(in); got != want {
			t.Errorf("hasGo(%q) = %v, want %v", in, got, want)
		}
	}
}
