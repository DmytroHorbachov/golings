// primitive_types_x085: Отрицательный ноль
// Make the tests pass!
// I AM NOT DONE
//
// formatTemp округляет температуру до целого и выводит её; значение "-0"
// выглядит странно и должно печататься как "0".
// Тренирует: в float64 существует -0, и он печатается со знаком.
// Сложность: hard
package main_test

import (
	"math"
	"strconv"
	"testing"
)

func formatTemp(t float64) string {
	r := math.Round(t)
	return strconv.FormatFloat(r, 'f', 0, 64) + "°"
}

func TestFormatTemp(t *testing.T) {
	cases := map[float64]string{-0.4: "0°", 0.4: "0°", -1.6: "-2°", 21.5: "22°"}
	for in, want := range cases {
		if got := formatTemp(in); got != want {
			t.Errorf("formatTemp(%v) = %q, want %q", in, got, want)
		}
	}
}
