// primitive_types_x020: Остаток для float
// Make the tests pass!
// I AM NOT DONE
//
// frac360 должна вернуть остаток от деления угла на 360 (угол неотрицательный).
// Код не компилируется: оператор % не определён для float64.
// Тренирует: math.Mod.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func frac360(angle float64) float64 {
	return angle % 360
}

func TestFrac360(t *testing.T) {
	_ = math.Mod
	cases := map[float64]float64{370.5: 10.5, 720: 0, 45: 45}
	for in, want := range cases {
		if got := frac360(in); got != want {
			t.Errorf("frac360(%v) = %v, want %v", in, got, want)
		}
	}
}
