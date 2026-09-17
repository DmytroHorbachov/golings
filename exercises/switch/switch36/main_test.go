// switch36
// Make the tests pass!

// I AM NOT DONE
//
// money форматирует сумму: "$5", "5 €", "5 ₽"; неизвестная валюта — "5 XXX".
// Тренирует: switch с разным порядком частей в ветках.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func money(amount int, currency string) string {
	n := strconv.Itoa(amount)
	switch currency {
	case "USD":
		return "$" + n
	case "EUR":
		return "€" + n
	}
	return n
}

func TestMoney(t *testing.T) {
	cases := []struct {
		amount   int
		cur, out string
	}{{5, "USD", "$5"}, {5, "EUR", "5 €"}, {7, "RUB", "7 ₽"}, {3, "GBP", "3 GBP"}}
	for _, c := range cases {
		if got := money(c.amount, c.cur); got != c.out {
			t.Errorf("money(%d, %s) = %q, want %q", c.amount, c.cur, got, c.out)
		}
	}
}
