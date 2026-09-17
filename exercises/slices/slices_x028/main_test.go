// slices_x028: Очистка с сохранением ёмкости
// Make the tests pass!
// I AM NOT DONE
//
// reset очищает буфер, оставляя выделенную память для повторного использования.
// Тренирует: s[:0].
// Сложность: easy
package main_test

import "testing"

func reset(s []int) []int {
	return nil
}

func TestReset(t *testing.T) {
	s := reset(make([]int, 5, 8))
	if len(s) != 0 || cap(s) != 8 {
		t.Errorf("reset: len=%d cap=%d, want 0 and 8", len(s), cap(s))
	}
}
