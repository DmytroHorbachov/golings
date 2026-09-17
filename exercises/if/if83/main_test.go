// if83
// Make the tests pass!

// I AM NOT DONE
//
// greeting: до 12 часов — "morning", до 18 — "afternoon", иначе "evening".
// Тренирует: последовательные проверки с ранним return.
// Сложность: easy
package main_test

import "testing"

func greeting(hour int) string {
	if hour < 12 {
		return "morning"
	}
	if hour <= 18 {
		return "afternoon"
	}
	return "evening"
}

func TestGreeting(t *testing.T) {
	cases := map[int]string{0: "morning", 11: "morning", 12: "afternoon", 17: "afternoon", 18: "evening", 23: "evening"}
	for in, want := range cases {
		if got := greeting(in); got != want {
			t.Errorf("greeting(%d) = %s, want %s", in, got, want)
		}
	}
}
