// variables_x071: Переменная-функция
// Make the tests pass!
// I AM NOT DONE
//
// Переменная op должна хранить функцию умножения.
// Сейчас она не инициализирована, и вызов вызывает панику.
// Тренирует: функции как значения и нулевое значение func-переменной.
// Сложность: easy
package main_test

import "testing"

func multiply(a, b int) int { return a * b }

func apply(a, b int) int {
	var op func(int, int) int
	return op(a, b)
}

func TestApply(t *testing.T) {
	if got := apply(6, 7); got != 42 {
		t.Errorf("apply(6, 7) = %d, want 42", got)
	}
}
