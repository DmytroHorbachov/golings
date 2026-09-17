// primitive_types_x015: Число в двоичную строку
// Make the tests pass!
// I AM NOT DONE
//
// toBinary должна вернуть двоичное представление числа.
// Тренирует: strconv.FormatInt.
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func toBinary(n int64) string {
	return strconv.FormatInt(n, 8)
}

func TestToBinary(t *testing.T) {
	cases := map[int64]string{5: "101", 8: "1000", 0: "0", -3: "-11"}
	for in, want := range cases {
		if got := toBinary(in); got != want {
			t.Errorf("toBinary(%d) = %s, want %s", in, got, want)
		}
	}
}
