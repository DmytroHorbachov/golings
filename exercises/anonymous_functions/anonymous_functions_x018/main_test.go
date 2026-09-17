// anonymous_functions_x018: Счётчик вызовов
// Make the tests pass!
// I AM NOT DONE
//
// track оборачивает функцию и считает её вызовы в захваченной переменной.
// Тренирует: изменение внешней переменной из литерала.
// Сложность: easy
package main_test

import "testing"

func track(f func(), calls *int) func() {
	return func() {
		*calls = 1
		f()
	}
}

func TestTrack(t *testing.T) {
	n := 0
	g := track(func() {}, &n)
	g()
	g()
	g()
	if n != 3 {
		t.Errorf("calls = %d, want 3", n)
	}
}
