// anonymous_functions_x033: Факториал в IIFE
// Make the tests pass!
// I AM NOT DONE
//
// table содержит факториал 5, вычисленный литералом с циклом внутри.
// Тренирует: IIFE для инициализации значения.
// Сложность: easy
package main_test

import "testing"

var fact5 = func() int {
	r := 1
	for i := 2; i < 5; i++ {
		r *= i
	}
	return r
}()

func TestFact5(t *testing.T) {
	if fact5 != 120 {
		t.Errorf("fact5 = %d, want 120", fact5)
	}
}
