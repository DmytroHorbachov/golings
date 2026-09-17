// range24
// Make the tests pass!

// I AM NOT DONE
//
// adler считает упрощённую контрольную сумму: a = 1 + сумма байт, b = сумма
// промежуточных a; результат b*65536 + a (все по модулю 65521).
// Тренирует: два аккумулятора в range по []byte.
// Сложность: medium
package main_test

import "testing"

const mod = 65521

func adler(data []byte) uint32 {
	a, b := uint32(1), uint32(0)
	for _, c := range data {
		b += a
		a += uint32(c)
	}
	return b<<16 | a
}

func TestAdler(t *testing.T) {
	if got := adler([]byte("Wikipedia")); got != 0x11E60398 {
		t.Errorf("adler = %#x, want 0x11e60398", got)
	}
}
