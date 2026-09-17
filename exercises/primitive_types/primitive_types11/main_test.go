// primitive_types11
// Make the tests pass!

// I AM NOT DONE
//
// roundCents округляет сумму до целых копеек по банковскому правилу:
// половина округляется до чётного (0.5 -> 0, 1.5 -> 2, 2.5 -> 2).
// Тренирует: math.Round округляет половину от нуля; есть math.RoundToEven.
// Сложность: hard
package main_test

import (
	"math"
	"testing"
)

func roundCents(x float64) float64 {
	return math.Round(x)
}

func TestRoundCents(t *testing.T) {
	cases := map[float64]float64{0.5: 0, 1.5: 2, 2.5: 2, 3.5: 4, -2.5: -2, 2.6: 3}
	for in, want := range cases {
		if got := roundCents(in); got != want {
			t.Errorf("roundCents(%v) = %v, want %v", in, got, want)
		}
	}
}
