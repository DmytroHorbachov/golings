// primitive_types_x100: Нормализация угла
// Make the tests pass!
// I AM NOT DONE
//
// normalize переводит угол в диапазон [0, 360).
// Для отрицательных углов результат отрицательный.
// Тренирует: math.Mod сохраняет знак делимого.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func normalize(deg float64) float64 {
	r := math.Mod(deg, 360)
	return r
}

func TestNormalize(t *testing.T) {
	cases := map[float64]float64{30: 30, 370: 10, -30: 330, -720: 0, 359.5: 359.5}
	for in, want := range cases {
		if got := normalize(in); got != want {
			t.Errorf("normalize(%v) = %v, want %v", in, got, want)
		}
	}
}
