// primitive_types_x019: Наибольшее из float
// Make the tests pass!
// I AM NOT DONE
//
// highest должна вернуть большее из двух чисел.
// Тренирует: функции math.Max и math.Min.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func highest(a, b float64) float64 {
	return math.Min(a, b)
}

func TestHighest(t *testing.T) {
	if got := highest(2.5, -1); got != 2.5 {
		t.Errorf("highest(2.5, -1) = %v", got)
	}
	if got := highest(-3, -2); got != -2 {
		t.Errorf("highest(-3, -2) = %v", got)
	}
}
