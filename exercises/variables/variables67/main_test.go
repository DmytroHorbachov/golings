// variables67
// Make the tests pass!

// I AM NOT DONE
//
// Функция toByte должна «прижимать» значение к диапазону 0..255.
// Сейчас 300 превращается в 44, а -5 — в 251.
// Тренирует: преобразование целых типов не проверяет диапазон.
// Сложность: hard
package main_test

import "testing"

func toByte(v int) byte {
	return byte(v)
}

func TestToByte(t *testing.T) {
	cases := map[int]byte{-5: 0, 0: 0, 128: 128, 255: 255, 300: 255}
	for in, want := range cases {
		if got := toByte(in); got != want {
			t.Errorf("toByte(%d) = %d, want %d", in, got, want)
		}
	}
}
