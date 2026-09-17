// if_x072: Строки — не числа
// Make the tests pass!
// I AM NOT DONE
//
// newer должна сравнить номера сборок, записанные строками без ведущих нулей.
// "10" должна быть новее, чем "9".
// Тренирует: строки сравниваются лексикографически, байт за байтом.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

func newer(a, b string) bool {
	if a > b {
		return true
	}
	return false
}

func TestNewer(t *testing.T) {
	_ = strconv.Atoi
	cases := []struct {
		a, b string
		want bool
	}{{"10", "9", true}, {"9", "10", false}, {"100", "99", true}, {"2", "1", true}, {"5", "5", false}}
	for _, c := range cases {
		if got := newer(c.a, c.b); got != c.want {
			t.Errorf("newer(%s, %s) = %v, want %v", c.a, c.b, got, c.want)
		}
	}
}
