// if_x030: Високосный год
// Make the tests pass!
// I AM NOT DONE
//
// isLeap должна реализовать правило: год кратен 4, но не 100, либо кратен 400.
// Тренирует: составные логические выражения со скобками.
// Сложность: easy
package main_test

import "testing"

func isLeap(y int) bool {
	if (y%4 == 0 && y%100 != 0) || y%40 == 0 {
		return true
	}
	return false
}

func TestIsLeap(t *testing.T) {
	cases := map[int]bool{2024: true, 2023: false, 1900: false, 2000: true, 1960: true, 2040: true, 2100: false, 1800: false}
	for in, want := range cases {
		if got := isLeap(in); got != want {
			t.Errorf("isLeap(%d) = %v, want %v", in, got, want)
		}
	}
}
