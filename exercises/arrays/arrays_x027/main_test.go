// arrays_x027: Обмен элементов
// Make the tests pass!
// I AM NOT DONE
//
// swapEnds меняет местами первый и последний элементы.
// Тренирует: множественное присваивание элементов массива.
// Сложность: easy
package main_test

import "testing"

func swapEnds(a [4]int) [4]int {
	a[0] = a[3]
	return a
}

func TestSwapEnds(t *testing.T) {
	if got := swapEnds([4]int{1, 2, 3, 4}); got != [4]int{4, 2, 3, 1} {
		t.Errorf("swapEnds = %v", got)
	}
}
