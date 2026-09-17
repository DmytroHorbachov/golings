// variables_x018: Константа из пакета math
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть наибольшее значение int32.
// Тренирует: предопределённые константы пакета math.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func maxInt32() int32 {
	var limit int32 = math.MaxInt16
	return limit
}

func TestMaxInt32(t *testing.T) {
	if got := maxInt32(); got != 2147483647 {
		t.Errorf("maxInt32() = %d, want 2147483647", got)
	}
}
