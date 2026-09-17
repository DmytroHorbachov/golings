// anonymous_functions33
// Make the tests pass!

// I AM NOT DONE
//
// evens возвращает замыкание, выдающее 0, 2, 4, ...
// Тренирует: замыкание-генератор.
// Сложность: easy
package main_test

import "testing"

func evens() func() int {
	n := -2
	return func() int {
		n++
		return n
	}
}

func TestEvens(t *testing.T) {
	next := evens()
	next()
	next()
	if got := next(); got != 4 {
		t.Errorf("third = %d, want 4", got)
	}
}
