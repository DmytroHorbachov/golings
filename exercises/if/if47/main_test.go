// if47
// Make the tests pass!

// I AM NOT DONE
//
// shipping должна вернуть 0 при сумме заказа от 3000, иначе 300.
// Тренирует: простое ветвление по порогу.
// Сложность: easy
package main_test

import "testing"

func shipping(total int) int {
	if total >= 3000 {
		return 300
	}
	return 300
}

func TestShipping(t *testing.T) {
	cases := map[int]int{2999: 300, 3000: 0, 10000: 0, 0: 300}
	for in, want := range cases {
		if got := shipping(in); got != want {
			t.Errorf("shipping(%d) = %d, want %d", in, got, want)
		}
	}
}
