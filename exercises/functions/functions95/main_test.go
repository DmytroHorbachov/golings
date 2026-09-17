// functions95
// Make the tests pass!

// I AM NOT DONE
//
// Функция power(base, exp) возводит base в степень exp.
// Вызов перепутал аргументы местами.
// Тренирует: позиционную передачу аргументов.
// Сложность: easy
package main_test

import "testing"

func power(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}

func cubeOfTwo() int {
	return power(3, 2)
}

func TestCubeOfTwo(t *testing.T) {
	if got := cubeOfTwo(); got != 8 {
		t.Errorf("cubeOfTwo() = %d, want 8", got)
	}
}
