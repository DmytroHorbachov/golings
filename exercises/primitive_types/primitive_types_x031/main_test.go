// primitive_types_x031: Расстояние между точками
// Make the tests pass!
// I AM NOT DONE
//
// distance должна вернуть расстояние между (x1, y1) и (x2, y2).
// Тренирует: math.Hypot.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func distance(x1, y1, x2, y2 float64) float64 {
	return math.Hypot(x2-x1, y2-x1)
}

func TestDistance(t *testing.T) {
	if got := distance(1, 1, 4, 5); got != 5 {
		t.Errorf("distance = %v, want 5", got)
	}
	if got := distance(2, 0, 2, 3); got != 3 {
		t.Errorf("distance = %v, want 3", got)
	}
}
