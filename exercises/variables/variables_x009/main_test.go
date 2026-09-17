// variables_x009: Секунд в сутках
// Make the tests pass!
// I AM NOT DONE
//
// Константа SecondsPerDay должна содержать число секунд в сутках.
// Тренирует: константные выражения, вычисляемые при компиляции.
// Сложность: easy
package main_test

import "testing"

const (
	SecondsPerMinute = 60
	MinutesPerHour   = 60
	HoursPerDay      = 24
)

const SecondsPerDay = SecondsPerMinute * MinutesPerHour

func TestSecondsPerDay(t *testing.T) {
	if SecondsPerDay != 86400 {
		t.Errorf("SecondsPerDay = %d, want 86400", SecondsPerDay)
	}
}
