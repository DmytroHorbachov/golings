// primitive_types_x058: Упаковка двух чисел
// Make the tests pass!
// I AM NOT DONE
//
// pack кладёт два uint16 в один uint32 (hi — старшая половина), unpack — обратно.
// Тренирует: расширение типа перед сдвигом.
// Сложность: medium
package main_test

import "testing"

func pack(hi, lo uint16) uint32 {
	return uint32(hi<<16) | uint32(lo)
}

func unpack(v uint32) (hi, lo uint16) {
	return uint16(v), uint16(v >> 16)
}

func TestPack(t *testing.T) {
	v := pack(0xABCD, 0x1234)
	if v != 0xABCD1234 {
		t.Errorf("pack = %#x, want 0xabcd1234", v)
	}
	hi, lo := unpack(v)
	if hi != 0xABCD || lo != 0x1234 {
		t.Errorf("unpack = %#x %#x", hi, lo)
	}
}
