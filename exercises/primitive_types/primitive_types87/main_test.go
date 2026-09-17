// primitive_types87
// Make the tests pass!

// I AM NOT DONE
//
// rgb раскладывает цвет 0xRRGGBB на три компоненты.
// Тренирует: сдвиги и маски над uint32.
// Сложность: medium
package main_test

import "testing"

func rgb(c uint32) (r, g, b uint8) {
	r = uint8(c >> 24)
	g = uint8(c >> 16)
	b = uint8(c >> 8)
	return
}

func TestRGB(t *testing.T) {
	r, g, b := rgb(0x1E90FF)
	if r != 0x1E || g != 0x90 || b != 0xFF {
		t.Errorf("rgb(0x1E90FF) = %x %x %x", r, g, b)
	}
}
