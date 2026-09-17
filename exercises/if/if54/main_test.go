// if54
// Make the tests pass!

// I AM NOT DONE
//
// portOrDefault разбирает порт; если строка не число или порт вне 1..65535 — 8080.
// Тренирует: if с инициализацией и несколькими условиями.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func portOrDefault(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		return 8080
	}
	return n
}

func TestPortOrDefault(t *testing.T) {
	cases := map[string]int{"80": 80, "65535": 65535, "0": 8080, "70000": 8080, "http": 8080}
	for in, want := range cases {
		if got := portOrDefault(in); got != want {
			t.Errorf("portOrDefault(%q) = %d, want %d", in, got, want)
		}
	}
}
