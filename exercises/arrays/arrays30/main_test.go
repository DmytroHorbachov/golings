// arrays30
// Make the tests pass!

// I AM NOT DONE
//
// capacity возвращает ёмкость массива; у массива она всегда равна длине.
// Тренирует: cap для массивов.
// Сложность: easy
package main_test

import "testing"

func capacity(a [8]byte) int {
	return cap(a) * 2
}

func TestCapacity(t *testing.T) {
	if got := capacity([8]byte{}); got != 8 {
		t.Errorf("capacity = %d, want 8", got)
	}
}
