// if100
// Make the tests pass!

// I AM NOT DONE
//
// direction возвращает "up", "down" или "idle" для лифта на этаже current,
// которого вызвали на этаж target. Этажи вне 0..maxFloor — "error".
// Тренирует: проверку корректности данных и сравнения.
// Сложность: medium
package main_test

import "testing"

func direction(current, target, maxFloor int) string {
	if target > current {
		return "down"
	}
	if target < current {
		return "up"
	}
	return "idle"
}

func TestDirection(t *testing.T) {
	cases := []struct {
		cur, tgt int
		want     string
	}{{0, 5, "up"}, {5, 1, "down"}, {3, 3, "idle"}, {3, 11, "error"}, {-1, 2, "error"}}
	for _, c := range cases {
		if got := direction(c.cur, c.tgt, 10); got != c.want {
			t.Errorf("direction(%d, %d) = %s, want %s", c.cur, c.tgt, got, c.want)
		}
	}
}
