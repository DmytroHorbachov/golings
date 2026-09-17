// variables71
// Make the tests pass!

// I AM NOT DONE
//
// Функция должна вернуть значение Big*4/8, где Big = 2^62.
// Промежуточный результат не помещается в int64.
// Тренирует: нетипизированные константы вычисляются с произвольной точностью.
// Сложность: hard
package main_test

import "testing"

const Big = 1 << 62

func halfOfDouble() int {
	x := Big
	return x * 4 / 8
}

func TestHalfOfDouble(t *testing.T) {
	if got := halfOfDouble(); got != 1<<61 {
		t.Errorf("halfOfDouble() = %d, want %d", got, 1<<61)
	}
}
