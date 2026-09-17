// primitive_types_x014: Двоичная строка в число
// Make the tests pass!
// I AM NOT DONE
//
// fromBinary должна разобрать строку из нулей и единиц.
// Тренирует: strconv.ParseInt с основанием.
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func fromBinary(s string) int64 {
	n, _ := strconv.ParseInt(s, 10, 64)
	return n
}

func TestFromBinary(t *testing.T) {
	cases := map[string]int64{"101": 5, "1111": 15, "0": 0}
	for in, want := range cases {
		if got := fromBinary(in); got != want {
			t.Errorf("fromBinary(%s) = %d, want %d", in, got, want)
		}
	}
}
