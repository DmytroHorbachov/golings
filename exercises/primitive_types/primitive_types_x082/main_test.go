// primitive_types_x082: WriteByte и Unicode
// Make the tests pass!
// I AM NOT DONE
//
// mask должна заменить каждый символ строки, кроме первого, на '•'.
// Сейчас результат содержит мусор вместо точек.
// Тренирует: '•' — руна вне диапазона byte; её нужно писать через WriteRune.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func mask(s string) string {
	var b strings.Builder
	for i, r := range []rune(s) {
		if i == 0 {
			b.WriteRune(r)
			continue
		}
		b.WriteByte(byte('•'))
	}
	return b.String()
}

func TestMask(t *testing.T) {
	cases := map[string]string{"secret": "s•••••", "пароль": "п•••••", "x": "x"}
	for in, want := range cases {
		if got := mask(in); got != want {
			t.Errorf("mask(%q) = %q, want %q", in, got, want)
		}
	}
}
