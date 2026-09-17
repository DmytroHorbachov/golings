// range36
// Make the tests pass!

// I AM NOT DONE
//
// shiftRussian сдвигает строчные русские буквы (без ё) на k позиций по кругу.
// Тренирует: range по рунам и арифметику над rune.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func shiftRussian(s string, k int) string {
	var b strings.Builder
	for _, r := range s {
		if r >= 'а' && r < 'я' {
			r = r + rune(k)
		}
		b.WriteRune(r)
	}
	return b.String()
}

func TestShiftRussian(t *testing.T) {
	cases := map[string]string{"абв": "бвг", "яма!": "анб!"}
	for in, want := range cases {
		if got := shiftRussian(in, 1); got != want {
			t.Errorf("shiftRussian(%q) = %q, want %q", in, got, want)
		}
	}
}
