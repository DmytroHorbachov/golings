// primitive_types62
// Make the tests pass!

// I AM NOT DONE
//
// mean возвращает среднее значение; для пустого среза должна вернуть 0.
// Во float деление 0/0 не паникует, а даёт NaN.
// Тренирует: деление на ноль для целых и для float64 ведёт себя по-разному.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func mean(xs []float64) float64 {
	var sum float64
	for _, x := range xs {
		sum += x
	}
	return sum / float64(len(xs))
}

func TestMean(t *testing.T) {
	if got := mean([]float64{1, 2, 3}); got != 2 {
		t.Errorf("mean(1,2,3) = %v", got)
	}
	if got := mean(nil); got != 0 || math.IsNaN(got) {
		t.Errorf("mean(nil) = %v, want 0", got)
	}
}
