// if80
// Make the tests pass!

// I AM NOT DONE
//
// winner возвращает "draw", "first" или "second" для ходов двух игроков.
// Тренирует: сложные логические условия.
// Сложность: medium
package main_test

import "testing"

func winner(a, b string) string {
	if a == "rock" && b == "scissors" {
		return "first"
	}
	return "second"
}

func TestWinner(t *testing.T) {
	cases := []struct{ a, b, want string }{
		{"rock", "rock", "draw"}, {"rock", "scissors", "first"}, {"scissors", "paper", "first"},
		{"paper", "rock", "first"}, {"rock", "paper", "second"}, {"paper", "scissors", "second"},
	}
	for _, c := range cases {
		if got := winner(c.a, c.b); got != c.want {
			t.Errorf("winner(%s, %s) = %s, want %s", c.a, c.b, got, c.want)
		}
	}
}
