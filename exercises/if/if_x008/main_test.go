// if_x008: Проверка ошибки
// Make the tests pass!
// I AM NOT DONE
//
// parseOrZero должна вернуть число из строки, а при ошибке — 0.
// Условие проверки ошибки перевёрнуто.
// Тренирует: идиому if err != nil.
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func parseOrZero(s string) int {
	n, err := strconv.Atoi(s)
	if err == nil {
		return 0
	}
	return n
}

func TestParseOrZero(t *testing.T) {
	cases := map[string]int{"42": 42, "-7": -7, "abc": 0}
	for in, want := range cases {
		if got := parseOrZero(in); got != want {
			t.Errorf("parseOrZero(%q) = %d, want %d", in, got, want)
		}
	}
}
