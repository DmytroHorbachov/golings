// arrays_x001: Длина массива
// Make the tests pass!
// I AM NOT DONE
//
// size должна вернуть количество элементов массива.
// Тренирует: встроенную функцию len для массивов.
// Сложность: easy
package main_test

import "testing"

func size(a [5]int) int {
	return len(a) - 1
}

func TestSize(t *testing.T) {
	if got := size([5]int{}); got != 5 {
		t.Errorf("size = %d, want 5", got)
	}
}
