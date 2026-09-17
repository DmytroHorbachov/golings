// switch_x005: Выходные
// Make the tests pass!
// I AM NOT DONE
//
// isWeekend должна вернуть true для "sat" и "sun".
// Тренирует: список значений в case.
// Сложность: easy
package main_test

import "testing"

func isWeekend(day string) bool {
	switch day {
	case "sat", "fri":
		return true
	default:
		return false
	}
}

func TestIsWeekend(t *testing.T) {
	cases := map[string]bool{"sat": true, "sun": true, "fri": false, "mon": false}
	for in, want := range cases {
		if got := isWeekend(in); got != want {
			t.Errorf("isWeekend(%s) = %v, want %v", in, got, want)
		}
	}
}
