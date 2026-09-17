// switch_x060: Поворот
// Make the tests pass!
// I AM NOT DONE
//
// turn поворачивает направление N/E/S/W налево или направо.
// Тренирует: вложенный выбор по двум параметрам.
// Сложность: medium
package main_test

import "testing"

func turn(dir, side string) string {
	order := "NESW"
	i := -1
	switch dir {
	case "N":
		i = 0
	case "E":
		i = 1
	case "S":
		i = 2
	case "W":
		i = 3
	default:
		return dir
	}
	switch side {
	case "right":
		i = (i + 3) % 4
	case "left":
		i = (i + 1) % 4
	}
	return string(order[i])
}

func TestTurn(t *testing.T) {
	cases := []struct{ dir, side, want string }{
		{"N", "right", "E"}, {"W", "right", "N"}, {"N", "left", "W"}, {"E", "left", "N"}, {"S", "back", "S"},
	}
	for _, c := range cases {
		if got := turn(c.dir, c.side); got != c.want {
			t.Errorf("turn(%s, %s) = %s, want %s", c.dir, c.side, got, c.want)
		}
	}
}
