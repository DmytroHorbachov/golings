// structs37
// Make the tests pass!

// I AM NOT DONE
//
// newCounter возвращает указатель на счётчик, начинающий с 1.
// Тренирует: &T{...}.
// Сложность: easy
package main_test

import "testing"

type Counter struct{ N int }

func newCounter() *Counter {
	return &Counter{}
}

func TestNewCounter(t *testing.T) {
	if c := newCounter(); c.N != 1 {
		t.Errorf("N = %d", c.N)
	}
}
