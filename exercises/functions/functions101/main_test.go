// functions101
// Make the tests pass!

// I AM NOT DONE
//
// fibGen должна вернуть функцию, которая при каждом вызове выдаёт следующее
// число Фибоначчи: 0, 1, 1, 2, 3, 5...
// Тренирует: замыкания, хранящие состояние между вызовами.
// Сложность: medium
package main_test

import "testing"

func fibGen() func() int {
	return func() int {
		a, b := 0, 1
		a, b = b, a+b
		return a
	}
}

func TestFibGen(t *testing.T) {
	next := fibGen()
	for i, want := range []int{0, 1, 1, 2, 3, 5, 8} {
		if got := next(); got != want {
			t.Errorf("call %d = %d, want %d", i, got, want)
		}
	}
}
