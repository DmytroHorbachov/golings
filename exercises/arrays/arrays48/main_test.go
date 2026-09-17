// arrays48
// Make the tests pass!

// I AM NOT DONE
//
// allOn возвращает true, только если все переключатели включены.
// Тренирует: проверку всех элементов массива.
// Сложность: easy
package main_test

import "testing"

func allOn(sw [4]bool) bool {
	for _, on := range sw {
		if on {
			return false
		}
	}
	return true
}

func TestAllOn(t *testing.T) {
	if !allOn([4]bool{true, true, true, true}) || allOn([4]bool{true, false, true, true}) {
		t.Errorf("allOn works incorrectly")
	}
}
