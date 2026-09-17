// primitive_types_x044: Проверка IPv4
// Make the tests pass!
// I AM NOT DONE
//
// validIPv4 проверяет адрес вида "192.168.0.1": четыре числа 0–255
// без ведущих нулей (кроме самого "0").
// Тренирует: strings.Split, strconv.Atoi и проверки диапазонов.
// Сложность: medium
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func validIPv4(s string) bool {
	parts := strings.Split(s, ".")
	for _, p := range parts {
		if _, err := strconv.Atoi(p); err != nil {
			return false
		}
	}
	return true
}

func TestValidIPv4(t *testing.T) {
	cases := map[string]bool{
		"192.168.0.1": true, "0.0.0.0": true, "255.255.255.255": true,
		"256.1.1.1": false, "1.2.3": false, "1.2.3.4.5": false, "01.2.3.4": false, "1..2.3": false, "-1.2.3.4": false,
	}
	for in, want := range cases {
		if got := validIPv4(in); got != want {
			t.Errorf("validIPv4(%q) = %v, want %v", in, got, want)
		}
	}
}
