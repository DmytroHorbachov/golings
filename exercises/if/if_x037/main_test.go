// if_x037: Налоговые ставки
// Make the tests pass!
// I AM NOT DONE
//
// tax считает налог: до 10000 — 0%, до 50000 — 10%, выше — 20% от всей суммы.
// Сейчас ставки применяются к неправильным диапазонам.
// Тренирует: цепочку else if с возрастающими порогами.
// Сложность: medium
package main_test

import "testing"

func tax(income int) int {
	if income < 10000 {
		return 0
	} else if income < 50000 {
		return income * 20 / 100
	}
	return income * 10 / 100
}

func TestTax(t *testing.T) {
	cases := map[int]int{5000: 0, 10000: 0, 20000: 2000, 50000: 5000, 100000: 20000}
	for in, want := range cases {
		if got := tax(in); got != want {
			t.Errorf("tax(%d) = %d, want %d", in, got, want)
		}
	}
}
