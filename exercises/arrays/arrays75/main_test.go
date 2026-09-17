// arrays75
// Make the tests pass!

// I AM NOT DONE
//
// indexOfMax должна вернуть индекс максимального элемента.
// Тренирует: две переменные в range по массиву.
// Сложность: easy
package main_test

import "testing"

func indexOfMax(a [5]int) int {
	best := 0
	for v, i := range a {
		if v > a[best] {
			best = i
		}
	}
	return best
}

func TestIndexOfMax(t *testing.T) {
	if got := indexOfMax([5]int{3, 9, 2, 9, 0}); got != 1 {
		t.Errorf("indexOfMax = %d, want 1", got)
	}
}
