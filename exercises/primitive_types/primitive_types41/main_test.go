// primitive_types41
// Make the tests pass!

// I AM NOT DONE
//
// bitsOf возвращает 8 бит числа int8 в дополнительном коде: -1 -> "11111111".
// Тренирует: преобразование int8 в uint8 и форматирование.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

func bitsOf(v int8) string {
	u := v
	return fmt.Sprintf("%b", u)
}

func TestBitsOf(t *testing.T) {
	cases := map[int8]string{-1: "11111111", 5: "00000101", -128: "10000000", 127: "01111111"}
	for in, want := range cases {
		if got := bitsOf(in); got != want {
			t.Errorf("bitsOf(%d) = %s, want %s", in, got, want)
		}
	}
}
