// switch_x018: Единицы размера
// Make the tests pass!
// I AM NOT DONE
//
// humanSize форматирует размер: байты, KB или MB.
// Тренирует: switch без тега с порогами.
// Сложность: easy
package main_test

import (
	"strconv"
	"testing"
)

func humanSize(n int) string {
	switch {
	case n >= 1024*1000:
		return strconv.Itoa(n/(1024*1024)) + "MB"
	case n >= 1024:
		return strconv.Itoa(n/1024) + "KB"
	default:
		return strconv.Itoa(n) + "B"
	}
}

func TestHumanSize(t *testing.T) {
	cases := map[int]string{512: "512B", 2048: "2KB", 1048575: "1023KB", 3 * 1048576: "3MB"}
	for in, want := range cases {
		if got := humanSize(in); got != want {
			t.Errorf("humanSize(%d) = %s, want %s", in, got, want)
		}
	}
}
