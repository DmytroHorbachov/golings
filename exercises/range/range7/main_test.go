// range7
// Make the tests pass!

// I AM NOT DONE
//
// squeeze заменяет подряд идущие пробелы одним пробелом.
// Тренирует: range по строке с запоминанием предыдущего символа.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func squeeze(s string) string {
	var b strings.Builder
	var prev rune
	for _, r := range s {
		if r == ' ' {
			continue
		}
		b.WriteRune(r)
	}
	return b.String()
}

func TestSqueeze(t *testing.T) {
	cases := map[string]string{"a   b  c": "a b c", "  x": " x", "no": "no"}
	for in, want := range cases {
		if got := squeeze(in); got != want {
			t.Errorf("squeeze(%q) = %q, want %q", in, got, want)
		}
	}
}
