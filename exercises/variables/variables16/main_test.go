// variables16
// Make the tests pass!

// I AM NOT DONE
//
// Права доступа rw-r--r-- в восьмеричной записи — 644, в десятичной — 420.
// Сейчас число записано в десятичной системе.
// Тренирует: восьмеричные литералы 0o.
// Сложность: easy
package main_test

import "testing"

func defaultPerm() int {
	perm := 644
	return perm
}

func TestDefaultPerm(t *testing.T) {
	if got := defaultPerm(); got != 420 {
		t.Errorf("defaultPerm() = %d, want 420 (0o644)", got)
	}
}
