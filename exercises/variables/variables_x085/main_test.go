// variables_x085: Минимум среди float
// Make the tests pass!
// I AM NOT DONE
//
// Функция minTemp должна вернуть минимальную температуру, а для пустого
// среза — +Inf. Сейчас для положительных температур возвращается 0.
// Тренирует: специальные значения float64 и начальное значение переменной.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

func minTemp(temps []float64) float64 {
	var m float64
	for _, v := range temps {
		if v > m {
			m = v
		}
	}
	return m
}

func TestMinTemp(t *testing.T) {
	if got := minTemp([]float64{12.5, 3.25, 7}); got != 3.25 {
		t.Errorf("minTemp = %v, want 3.25", got)
	}
	if got := minTemp(nil); !math.IsInf(got, 1) {
		t.Errorf("minTemp(nil) = %v, want +Inf", got)
	}
}
