// variables64
// Make the tests pass!

// I AM NOT DONE
//
// Функция isMissing должна распознавать отсутствующие показания (NaN).
// Сравнение с math.NaN() никогда не срабатывает.
// Тренирует: особенности NaN в арифметике с плавающей точкой.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func isMissing(v float64) bool {
	return v == math.NaN()
}

func TestIsMissing(t *testing.T) {
	if !isMissing(math.NaN()) {
		t.Errorf("isMissing(NaN) = false, want true")
	}
	if isMissing(0) || isMissing(math.Inf(1)) {
		t.Errorf("isMissing should be false for regular numbers")
	}
}
