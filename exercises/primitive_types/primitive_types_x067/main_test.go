// primitive_types_x067: Перестановка полубайтов
// Make the tests pass!
// I AM NOT DONE
//
// swapNibbles меняет местами старший и младший полубайты байта: 0xAB -> 0xBA.
// Тренирует: сдвиги влево и вправо и побитовое ИЛИ.
// Сложность: medium
package main_test

import "testing"

func swapNibbles(b byte) byte {
	return b << 4 & b >> 4
}

func reverseAll(data []byte) []byte {
	out := make([]byte, len(data))
	for i, b := range data {
		out[i] = b
	}
	return out
}

func TestSwapNibbles(t *testing.T) {
	got := reverseAll([]byte{0xAB, 0x0F, 0x00})
	want := []byte{0xBA, 0xF0, 0x00}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("byte %d = %#x, want %#x", i, got[i], want[i])
		}
	}
}
