// if39
// Make the tests pass!

// I AM NOT DONE
//
// validateWeight должна отклонять отрицательные значения и NaN.
// Сейчас NaN принимается как корректный вес.
// Тренирует: любое сравнение с NaN ложно.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func validateWeight(w float64) bool {
	if w < 0 {
		return false
	}
	return true
}

func TestValidateWeight(t *testing.T) {
	if !validateWeight(70.5) || !validateWeight(0) {
		t.Errorf("70.5 and 0 are valid weights")
	}
	if validateWeight(-1) {
		t.Errorf("-1 is not a valid weight")
	}
	if validateWeight(math.NaN()) {
		t.Errorf("NaN is not a valid weight")
	}
}
