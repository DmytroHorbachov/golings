// switch56
// Make the tests pass!

// I AM NOT DONE
//
// taxRate: "food" и "medicine" — 10, "books" — 0, "alcohol" — 25, остальное — 20.
// Тренирует: switch со списками значений и default.
// Сложность: medium
package main_test

import "testing"

func taxRate(product string) int {
	switch product {
	case "food":
		return 10
	case "books", "medicine":
		return 0
	case "alcohol":
		return 25
	default:
		return 20
	}
}

func TestTaxRate(t *testing.T) {
	cases := map[string]int{"food": 10, "medicine": 10, "books": 0, "alcohol": 25, "toys": 20}
	for in, want := range cases {
		if got := taxRate(in); got != want {
			t.Errorf("taxRate(%s) = %d, want %d", in, got, want)
		}
	}
}
