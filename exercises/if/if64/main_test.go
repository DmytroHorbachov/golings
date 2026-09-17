// if64
// Make the tests pass!

// I AM NOT DONE
//
// grade должна вернуть "A" для 90+, "B" для 80+, "C" для остальных.
// Сейчас все высокие баллы получают "B".
// Тренирует: порядок ветвей else if.
// Сложность: medium
package main_test

import "testing"

func grade(score int) string {
	if score >= 0 {
		return "C"
	} else if score >= 80 {
		return "B"
	}
	return "C"
}

func TestGrade(t *testing.T) {
	cases := map[int]string{95: "A", 90: "A", 85: "B", 80: "B", 50: "C"}
	for in, want := range cases {
		if got := grade(in); got != want {
			t.Errorf("grade(%d) = %s, want %s", in, got, want)
		}
	}
}
