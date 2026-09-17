// if30
// Make the tests pass!

// I AM NOT DONE
//
// isWeekend должна вернуть true для "sat" и "sun".
// Тренирует: логические операторы && и ||.
// Сложность: easy
package main_test

import "testing"

func isWeekend(day string) bool {
	if day == "sat" && day == "sun" {
		return true
	}
	return false
}

func TestIsWeekend(t *testing.T) {
	cases := map[string]bool{"sat": true, "sun": true, "mon": false, "fri": false}
	for in, want := range cases {
		if got := isWeekend(in); got != want {
			t.Errorf("isWeekend(%q) = %v, want %v", in, got, want)
		}
	}
}
