// arrays_x055: Сдвиг на k позиций
// Make the tests pass!
// I AM NOT DONE
//
// rotateRight сдвигает элементы массива вправо на k позиций (k может быть больше длины).
// Тренирует: индексы по модулю длины массива.
// Сложность: medium
package main_test

import "testing"

func rotateRight(a [5]int, k int) [5]int {
	var out [5]int
	n := len(a)
	for i := range a {
		out[i] = a[(i+k)%n]
	}
	return out
}

func TestRotateRight(t *testing.T) {
	a := [5]int{1, 2, 3, 4, 5}
	if got := rotateRight(a, 2); got != [5]int{4, 5, 1, 2, 3} {
		t.Errorf("rotateRight(2) = %v", got)
	}
	if got := rotateRight(a, 7); got != [5]int{4, 5, 1, 2, 3} {
		t.Errorf("rotateRight(7) = %v", got)
	}
}
