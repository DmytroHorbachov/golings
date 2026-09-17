// functions_x044: Взаимная рекурсия
// Make the tests pass!
// I AM NOT DONE
//
// isEven и isOdd определены друг через друга для неотрицательных n.
// Сейчас базовые случаи и вызовы перепутаны.
// Тренирует: взаимно рекурсивные функции.
// Сложность: medium
package main_test

import "testing"

func isEven(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func isOdd(n int) bool {
	if n == 0 {
		return false
	}
	return isEven(n - 1)
}

func TestParity(t *testing.T) {
	for n := 0; n < 10; n++ {
		if isEven(n) != (n%2 == 0) || isOdd(n) != (n%2 == 1) {
			t.Errorf("parity of %d: isEven=%v isOdd=%v", n, isEven(n), isOdd(n))
		}
	}
}
