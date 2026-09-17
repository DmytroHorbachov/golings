// variables_x078: Двоичная маска
// Make the tests pass!
// I AM NOT DONE
//
// Функция lowNibble должна вернуть младшие 4 бита числа.
// Маска записана двоичным литералом, но в ней ошибка.
// Тренирует: двоичные литералы 0b и побитовое И.
// Сложность: easy
package main_test

import "testing"

func lowNibble(v uint8) uint8 {
	const mask = 0b1010
	return v & mask
}

func TestLowNibble(t *testing.T) {
	cases := map[uint8]uint8{0xAB: 0xB, 0x0F: 0xF, 0xF0: 0, 0x37: 7}
	for in, want := range cases {
		if got := lowNibble(in); got != want {
			t.Errorf("lowNibble(%#x) = %#x, want %#x", in, got, want)
		}
	}
}
