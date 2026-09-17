// functions59
// Make the tests pass!

// I AM NOT DONE
//
// Функция safeRatio должна вернуть 0, если знаменатель равен нулю.
// Сейчас проверка есть, но функция продолжает выполнение.
// Тренирует: ранний return (guard clause).
// Сложность: easy
package main_test

import "testing"

func safeRatio(a, b int) int {
	result := 0
	if b == 0 {
		result = 0
	}
	result = a / b
	return result
}

func TestSafeRatio(t *testing.T) {
	if got := safeRatio(10, 2); got != 5 {
		t.Errorf("safeRatio(10, 2) = %d, want 5", got)
	}
	if got := safeRatio(1, 0); got != 0 {
		t.Errorf("safeRatio(1, 0) = %d, want 0", got)
	}
}
