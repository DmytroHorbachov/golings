// arrays_x068: range копирует массив
// Make the tests pass!
// I AM NOT DONE
//
// propagate проходит по массиву и переносит «заражение» вперёд: если элемент
// заражён, следующий тоже становится заражённым. Сейчас распространяется только на один шаг.
// Тренирует: range по массиву-значению вычисляет копию массива один раз.
// Сложность: hard
package main_test

import "testing"

func propagate(a [5]bool) [5]bool {
	for i, v := range a {
		if v && i+1 < len(a) {
			a[i+1] = true
		}
	}
	return a
}

func TestPropagate(t *testing.T) {
	if got := propagate([5]bool{false, true, false, false, false}); got != [5]bool{false, true, true, true, true} {
		t.Errorf("propagate = %v", got)
	}
}
