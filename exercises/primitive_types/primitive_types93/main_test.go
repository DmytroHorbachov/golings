// primitive_types93
// Make the tests pass!

// I AM NOT DONE
//
// edge должна вернуть длину ребра куба по его объёму.
// Тренирует: функции пакета math для корней.
// Сложность: easy
package main_test

import (
	"math"
	"testing"
)

func edge(volume float64) float64 {
	return math.Sqrt(volume)
}

func TestEdge(t *testing.T) {
	cases := map[float64]float64{27: 3, 8: 2, 1: 1, 0: 0}
	for in, want := range cases {
		if got := edge(in); got != want {
			t.Errorf("edge(%v) = %v, want %v", in, got, want)
		}
	}
}
