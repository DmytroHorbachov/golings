// generics_x027: Обмен по указателям
// Make the tests pass!
// I AM NOT DONE
//
// SwapPtr меняет местами значения по двум указателям.
// Тренирует: обобщённые указатели *T.
// Сложность: easy
package main_test

import "testing"

func SwapPtr[T any](a, b *T) {
	a, b = b, a
}

func TestSwapPtr(t *testing.T) {
	x, y := "left", "right"
	SwapPtr(&x, &y)
	if x != "right" || y != "left" {
		t.Errorf("x=%s y=%s", x, y)
	}
}
