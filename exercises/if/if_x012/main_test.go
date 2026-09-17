// if_x012: Агрегатное состояние воды
// Make the tests pass!
// I AM NOT DONE
//
// waterState должна вернуть "ice" ниже 0, "steam" от 100, иначе "water".
// Тренирует: цепочку условий с граничными значениями.
// Сложность: easy
package main_test

import "testing"

func waterState(t float64) string {
	if t < 0 {
		return "ice"
	}
	if t > 100 {
		return "steam"
	}
	return "water"
}

func TestWaterState(t *testing.T) {
	cases := map[float64]string{-5: "ice", 0: "water", 99.9: "water", 100: "steam", 120: "steam"}
	for in, want := range cases {
		if got := waterState(in); got != want {
			t.Errorf("waterState(%v) = %s, want %s", in, got, want)
		}
	}
}
