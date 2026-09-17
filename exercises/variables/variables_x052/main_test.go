// variables_x052: Середина без переполнения
// Make the tests pass!
// I AM NOT DONE
//
// Функция midpoint должна вернуть середину отрезка [lo, hi] для любых int64.
// На больших значениях результат становится отрицательным.
// Тренирует: переполнение при сложении больших чисел.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func midpoint(lo, hi int64) int64 {
	return (lo + hi) / 2
}

func TestMidpoint(t *testing.T) {
	if got := midpoint(2, 10); got != 6 {
		t.Errorf("midpoint(2, 10) = %d, want 6", got)
	}
	big := int64(math.MaxInt64)
	if got := midpoint(big-10, big); got != big-5 {
		t.Errorf("midpoint(max-10, max) = %d, want %d", got, big-5)
	}
}
