// anonymous_functions_x001: Немедленный вызов
// Make the tests pass!
// I AM NOT DONE
//
// limit вычисляется функциональным литералом, вызываемым сразу.
// Тренирует: IIFE (immediately invoked function expression).
// Сложность: easy
package main_test

import "testing"

func limit() int {
	l := func() int {
		base := 10
		return base + 10
	}()
	return l
}

func TestLimit(t *testing.T) {
	if limit() != 100 {
		t.Errorf("limit = %d, want 100", limit())
	}
}
