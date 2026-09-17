// switch_x037: Форматирование по типу
// Make the tests pass!
// I AM NOT DONE
//
// format превращает значение в строку: int — число, string — в кавычках,
// bool — "yes"/"no", остальное — "?".
// Тренирует: type switch с привязкой переменной (v := x.(type)).
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func format(x interface{}) string {
	switch v := x.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	}
	return "?"
}

func TestFormat(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{42, "42"}, {"go", `"go"`}, {true, "yes"}, {false, "no"}, {3.14, "?"}}
	for _, c := range cases {
		if got := format(c.in); got != c.want {
			t.Errorf("format(%v) = %s, want %s", c.in, got, c.want)
		}
	}
}
