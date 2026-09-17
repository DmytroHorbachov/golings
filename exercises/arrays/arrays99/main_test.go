// arrays99
// Make the tests pass!

// I AM NOT DONE
//
// squares возвращает массив квадратов индексов: [0 1 4 9 16].
// Тренирует: запись в элементы массива в цикле.
// Сложность: easy
package main_test

import "testing"

func squares() [5]int {
	var a [5]int
	for i := range a {
		a[i] = i * 2
	}
	return a
}

func TestSquares(t *testing.T) {
	if got := squares(); got != [5]int{0, 1, 4, 9, 16} {
		t.Errorf("squares = %v", got)
	}
}
