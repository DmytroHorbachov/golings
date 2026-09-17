// primitive_types45
// Make the tests pass!

// I AM NOT DONE
//
// upper turns lowercase russian letters into capitals by subtracting 32.
// For the letter "ё" the result is wrong.
// The offset between the cases is not the same for every character.
package main_test

import (
	"testing"
	"unicode"
)

func upper(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] >= 'а' && r[i] <= 'я' || r[i] == 'ё' {
			r[i] -= 32
		}
	}
	return string(r)
}

func TestUpper(t *testing.T) {
	_ = unicode.ToUpper
	cases := map[string]string{"мир": "МИР", "ёж": "ЁЖ", "ещё": "ЕЩЁ"}
	for in, want := range cases {
		if got := upper(in); got != want {
			t.Errorf("upper(%q) = %q, want %q", in, got, want)
		}
	}
}
