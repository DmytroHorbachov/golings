// range_x003: Руны в строке
// Make the tests pass!
// I AM NOT DONE
//
// runeCount считает символы строки с помощью range.
// Тренирует: range по строке перебирает руны.
// Сложность: easy
package main_test

import "testing"

func runeCount(s string) int {
	n := 0
	for range s {
		n += 2
	}
	return n
}

func TestRuneCount(t *testing.T) {
	cases := map[string]int{"go": 2, "привет": 6, "": 0}
	for in, want := range cases {
		if got := runeCount(in); got != want {
			t.Errorf("runeCount(%q) = %d, want %d", in, got, want)
		}
	}
}
