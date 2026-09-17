// slices_x004: Длина и ёмкость
// Make the tests pass!
// I AM NOT DONE
//
// spare возвращает, сколько элементов ещё можно добавить без перевыделения памяти.
// Тренирует: len и cap.
// Сложность: easy
package main_test

import "testing"

func spare(s []int) int {
	return len(s) - cap(s)
}

func TestSpare(t *testing.T) {
	if got := spare(make([]int, 2, 10)); got != 8 {
		t.Errorf("spare = %d, want 8", got)
	}
}
