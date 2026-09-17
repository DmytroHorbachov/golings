// primitive_types104
// Make the tests pass!

// I AM NOT DONE
//
// boxes должна вернуть число коробок для items предметов, если в коробку
// помещается 6 штук.
// Тренирует: math.Ceil и преобразования float64/int.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func boxes(items int) int {
	return int(math.Floor(float64(items) / 6))
}

func TestBoxes(t *testing.T) {
	cases := map[int]int{0: 0, 1: 1, 6: 1, 7: 2, 12: 2, 13: 3}
	for in, want := range cases {
		if got := boxes(in); got != want {
			t.Errorf("boxes(%d) = %d, want %d", in, got, want)
		}
	}
}
