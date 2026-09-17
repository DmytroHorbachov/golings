// range_x020: Цифры в строке
// Make the tests pass!
// I AM NOT DONE
//
// countDigits считает цифры в строке.
// Тренирует: range по строке и unicode.IsDigit.
// Сложность: easy
package main_test

import (
	"testing"
	"unicode"
)

func countDigits(s string) int {
	n := 0
	for _, r := range s {
		if unicode.IsLetter(r) {
			n++
		}
	}
	return n
}

func TestCountDigits(t *testing.T) {
	if got := countDigits("room 42, floor 7"); got != 3 {
		t.Errorf("countDigits = %d, want 3", got)
	}
}
