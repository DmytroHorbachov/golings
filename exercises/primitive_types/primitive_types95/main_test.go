// primitive_types95
// Make the tests pass!

// I AM NOT DONE
//
// lowestStart должна вернуть отрицательную бесконечность — стартовое значение
// для поиска максимума.
// Тренирует: math.Inf и знак бесконечности.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func lowestStart() float64 {
	return math.Inf(1)
}

func TestLowestStart(t *testing.T) {
	if got := lowestStart(); !math.IsInf(got, -1) {
		t.Errorf("lowestStart() = %v, want -Inf", got)
	}
}
