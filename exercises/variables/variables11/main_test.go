// variables11
// Make the tests pass!

// I AM NOT DONE
//
// Функция parseLimit должна вернуть число из строки, а при ошибке — 10.
// Тренирует: обработку второго возвращаемого значения (error).
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

const defaultLimit = 10

func parseLimit(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func TestParseLimit(t *testing.T) {
	cases := map[string]int{"25": 25, "0": 0, "abc": 10, "": 10}
	for in, want := range cases {
		if got := parseLimit(in); got != want {
			t.Errorf("parseLimit(%q) = %d, want %d", in, got, want)
		}
	}
}
