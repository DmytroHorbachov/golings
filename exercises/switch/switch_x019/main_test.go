// switch_x019: Стороны света в градусах
// Make the tests pass!
// I AM NOT DONE
//
// degrees возвращает азимут направления: N=0, E=90, S=180, W=270.
// Тренирует: switch по строке.
// Сложность: easy
package main_test

import "testing"

func degrees(dir string) int {
	switch dir {
	case "N":
		return 0
	case "E":
		return 90
	case "S":
		return 180
	case "W":
		return 360
	}
	return -1
}

func TestDegrees(t *testing.T) {
	cases := map[string]int{"N": 0, "E": 90, "S": 180, "W": 270, "?": -1}
	for in, want := range cases {
		if got := degrees(in); got != want {
			t.Errorf("degrees(%s) = %d, want %d", in, got, want)
		}
	}
}
