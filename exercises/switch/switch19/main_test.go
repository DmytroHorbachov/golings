// switch19
// Make the tests pass!

// I AM NOT DONE
//
// key переводит клавишу в команду; регистр букв не важен.
// Тренирует: нормализацию значения прямо в выражении switch.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func key(k string) string {
	switch k {
	case "w":
		return "up"
	case "s":
		return "down"
	case "q":
		return "quit"
	}
	return "noop"
}

func TestKey(t *testing.T) {
	_ = strings.ToLower
	cases := map[string]string{"w": "up", "W": "up", "S": "down", "q": "quit", "x": "noop"}
	for in, want := range cases {
		if got := key(in); got != want {
			t.Errorf("key(%s) = %s, want %s", in, got, want)
		}
	}
}
