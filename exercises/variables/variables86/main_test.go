// variables86
// Make the tests pass!

// I AM NOT DONE
//
// Функция pageSize принимает *int: nil означает «использовать 20».
// Сейчас при nil функция паникует.
// Тренирует: нулевое значение указателя и проверку на nil.
// Сложность: medium
package main_test

import "testing"

func pageSize(size *int) int {
	return *size
}

func TestPageSize(t *testing.T) {
	if got := pageSize(nil); got != 20 {
		t.Errorf("pageSize(nil) = %d, want 20", got)
	}
	n := 50
	if got := pageSize(&n); got != 50 {
		t.Errorf("pageSize(&50) = %d, want 50", got)
	}
}
