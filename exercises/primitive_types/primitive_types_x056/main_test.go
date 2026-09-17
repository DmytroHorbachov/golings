// primitive_types_x056: Формат чч:мм
// Make the tests pass!
// I AM NOT DONE
//
// hhmm переводит минуты от полуночи в строку "07:05".
// Тренирует: целочисленное деление и заполнение нулями в fmt.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

func hhmm(minutes int) string {
	h, m := minutes/100, minutes%100
	return fmt.Sprintf("%d:%d", h, m)
}

func TestHHMM(t *testing.T) {
	cases := map[int]string{425: "07:05", 0: "00:00", 1439: "23:59", 600: "10:00"}
	for in, want := range cases {
		if got := hhmm(in); got != want {
			t.Errorf("hhmm(%d) = %s, want %s", in, got, want)
		}
	}
}
