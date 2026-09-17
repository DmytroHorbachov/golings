// if49
// Make the tests pass!

// I AM NOT DONE
//
// ticketPrice: дети до 12 лет — 5, пенсионеры от 65 — 7, остальные — 10.
// Тренирует: цепочку условий по диапазонам.
// Сложность: easy
package main_test

import "testing"

func ticketPrice(age int) int {
	if age < 12 {
		return 5
	} else if age > 65 {
		return 7
	}
	return 10
}

func TestTicketPrice(t *testing.T) {
	cases := map[int]int{5: 5, 12: 10, 40: 10, 65: 7, 80: 7}
	for in, want := range cases {
		if got := ticketPrice(in); got != want {
			t.Errorf("ticketPrice(%d) = %d, want %d", in, got, want)
		}
	}
}
