// algorithms76
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: битовые трюки. Посчитайте единичные биты в беззнаковом числе
// без math/bits.
// Сложность: easy. Ожидаемая асимптотика: O(k) по времени (k — число единиц), O(1) по памяти
package main_test

import "testing"

func hammingWeight(x uint64) int {
	return 0
}

func TestHammingWeight(t *testing.T) {
	cases := map[uint64]int{0: 0, 1: 1, 11: 3, 128: 1, 1 << 63: 1, ^uint64(0): 64}
	for in, want := range cases {
		if got := hammingWeight(in); got != want {
			t.Errorf("hammingWeight(%d) = %d, want %d", in, got, want)
		}
	}
}
