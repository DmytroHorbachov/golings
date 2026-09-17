// variables17
// Make the tests pass!

// I AM NOT DONE
//
// Функция reset должна обнулить переданный счётчик.
// Сейчас она работает с копией значения.
// Тренирует: указатели и изменение переменной через них.
// Сложность: medium
package main_test

import "testing"

func reset(counter int) {
	counter = 0
}

func resetAndGet() int {
	hits := 42
	reset(hits)
	return hits
}

func TestReset(t *testing.T) {
	if got := resetAndGet(); got != 0 {
		t.Errorf("after reset hits = %d, want 0", got)
	}
}
