// primitive_types_x017: Целая часть
// Make the tests pass!
// I AM NOT DONE
//
// wholePart должна отбросить дробную часть, сохранив знак: -2.7 -> -2.
// Тренирует: math.Trunc в отличие от math.Floor.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func wholePart(x float64) float64 {
	return math.Floor(x)
}

func TestWholePart(t *testing.T) {
	cases := map[float64]float64{2.7: 2, -2.7: -2, 5: 5}
	for in, want := range cases {
		if got := wholePart(in); got != want {
			t.Errorf("wholePart(%v) = %v, want %v", in, got, want)
		}
	}
}
