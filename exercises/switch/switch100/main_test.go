// switch100
// Make the tests pass!

// I AM NOT DONE
//
// delta переводит клавишу в смещение (dx, dy). Поддерживаются WASD и стрелки.
// Ось y направлена вниз.
// Тренирует: объединение альтернативных значений в case.
// Сложность: medium
package main_test

import "testing"

func delta(key string) (int, int) {
	switch key {
	case "w":
		return 0, 1
	case "s", "down":
		return 0, 1
	case "a", "left":
		return -1, 0
	case "d":
		return 1, 0
	}
	return 0, 0
}

func TestDelta(t *testing.T) {
	cases := map[string][2]int{"w": {0, -1}, "up": {0, -1}, "s": {0, 1}, "left": {-1, 0}, "right": {1, 0}, "x": {0, 0}}
	for in, want := range cases {
		dx, dy := delta(in)
		if dx != want[0] || dy != want[1] {
			t.Errorf("delta(%s) = (%d, %d), want %v", in, dx, dy, want)
		}
	}
}
