// primitive_types85
// Make the tests pass!

// I AM NOT DONE
//
// highBitSet проверяет, установлен ли старший бит у значения int8.
// Код не компилируется: маска 0x80 не помещается в int8.
// Тренирует: типизированные операнды и диапазон констант.
// Сложность: hard
package main_test

import "testing"

func highBitSet(v int8) bool {
	return v&0x80 != 0
}

func TestHighBitSet(t *testing.T) {
	cases := map[int8]bool{-1: true, -128: true, 127: false, 0: false, 64: false}
	for in, want := range cases {
		if got := highBitSet(in); got != want {
			t.Errorf("highBitSet(%d) = %v, want %v", in, got, want)
		}
	}
}
