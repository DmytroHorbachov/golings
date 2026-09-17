// arrays91
// Make the tests pass!

// I AM NOT DONE
//
// blend смешивает два цвета [3]uint8 с весом w (0..1) для второго цвета.
// Тренирует: поэлементную арифметику над массивами.
// Сложность: medium
package main_test

import (
	"math"
	"testing"
)

func blend(a, b [3]uint8, w float64) [3]uint8 {
	var out [3]uint8
	for i := range out {
		out[i] = a[i] + uint8(float64(b[i])*w)
	}
	return out
}

func TestBlend(t *testing.T) {
	_ = math.Round
	red, blue := [3]uint8{255, 0, 0}, [3]uint8{0, 0, 255}
	if got := blend(red, blue, 0.5); got != [3]uint8{128, 0, 128} {
		t.Errorf("blend 50%% = %v", got)
	}
	if got := blend(red, blue, 0); got != red {
		t.Errorf("blend 0%% = %v", got)
	}
}
