// primitive_types_x032: Количество вхождений
// Make the tests pass!
// I AM NOT DONE
//
// countA должна посчитать, сколько раз буква 'a' встречается в строке.
// Тренирует: strings.Count.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func countA(s string) int {
	return strings.Count(s, "A")
}

func TestCountA(t *testing.T) {
	cases := map[string]int{"banana": 3, "Apple": 0, "": 0}
	for in, want := range cases {
		if got := countA(in); got != want {
			t.Errorf("countA(%q) = %d, want %d", in, got, want)
		}
	}
}
