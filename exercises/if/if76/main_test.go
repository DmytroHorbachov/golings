// if76
// Make the tests pass!

// I AM NOT DONE
//
// discount: участники клуба получают 10%, заказы от 5000 — 5%,
// а участники с заказом от 5000 — 15%.
// Тренирует: порядок проверки комбинированных условий.
// Сложность: medium
package main_test

import "testing"

func discount(member bool, total int) int {
	if member {
		return 10
	} else if total >= 5000 {
		return 5
	}
	return 0
}

func TestDiscount(t *testing.T) {
	cases := []struct {
		member bool
		total  int
		want   int
	}{{true, 6000, 15}, {true, 100, 10}, {false, 5000, 5}, {false, 100, 0}}
	for _, c := range cases {
		if got := discount(c.member, c.total); got != c.want {
			t.Errorf("discount(%v, %d) = %d, want %d", c.member, c.total, got, c.want)
		}
	}
}
