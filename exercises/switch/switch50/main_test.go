// switch50
// Make the tests pass!

// I AM NOT DONE
//
// isWorkday использует time.Weekday. Суббота и воскресенье — выходные.
// Тренирует: switch по значениям именованного типа из стандартной библиотеки.
// Сложность: easy
package main_test

import (
	"testing"
	"time"
)

func isWorkday(d time.Weekday) bool {
	switch d {
	case time.Saturday, time.Monday:
		return false
	}
	return true
}

func TestIsWorkday(t *testing.T) {
	cases := map[time.Weekday]bool{time.Monday: true, time.Friday: true, time.Saturday: false, time.Sunday: false}
	for in, want := range cases {
		if got := isWorkday(in); got != want {
			t.Errorf("isWorkday(%s) = %v, want %v", in, got, want)
		}
	}
}
