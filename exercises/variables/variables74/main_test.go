// variables74
// Make the tests pass!

// I AM NOT DONE
//
// Функция swapPtr должна поменять местами значения, на которые указывают a и b.
// Сейчас меняются только локальные копии указателей.
// Тренирует: разыменование указателей при присваивании.
// Сложность: easy
package main_test

import "testing"

func swapPtr(a, b *int) {
	a, b = b, a
}

func TestSwapPtr(t *testing.T) {
	x, y := 1, 2
	swapPtr(&x, &y)
	if x != 2 || y != 1 {
		t.Errorf("after swapPtr x=%d y=%d, want x=2 y=1", x, y)
	}
}
