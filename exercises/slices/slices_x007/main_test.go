// slices_x007: Результат copy
// Make the tests pass!
// I AM NOT DONE
//
// copied должна вернуть количество скопированных элементов.
// Тренирует: copy возвращает число скопированных элементов.
// Сложность: easy
package main_test

import "testing"

func copied(dst, src []int) int {
	n := copy(dst, src)
	return len(src)
}

func TestCopied(t *testing.T) {
	if got := copied(make([]int, 2), []int{1, 2, 3}); got != 2 {
		t.Errorf("copied = %d, want 2", got)
	}
}
