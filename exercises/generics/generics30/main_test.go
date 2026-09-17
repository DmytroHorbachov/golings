// generics30
// Make the tests pass!

// I AM NOT DONE
//
// Box[T] хранит значение; Get его возвращает.
// Тренирует: методы обобщённого типа.
// Сложность: easy
package main_test

import "testing"

type Box[T any] struct{ v T }

func (b Box[T]) Get() T {
	return *new(T)
}

func TestBox(t *testing.T) {
	if (Box[string]{"x"}).Get() != "x" || (Box[int]{7}).Get() != 7 {
		t.Errorf("Box works incorrectly")
	}
}
