// if58
// Make the tests pass!

// I AM NOT DONE
//
// describe должна вернуть строку в верхнем регистре, если значение — строка,
// иначе "not a string". Сейчас на числе функция паникует.
// Тренирует: форма v, ok := x.(T) не паникует при несовпадении типа.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func describe(v interface{}) string {
	if s := v.(string); s != "" {
		return strings.ToUpper(s)
	}
	return "not a string"
}

func TestDescribe(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{"go", "GO"}, {42, "not a string"}, {nil, "not a string"}, {"", ""}}
	for _, c := range cases {
		if got := describe(c.in); got != c.want {
			t.Errorf("describe(%v) = %q, want %q", c.in, got, c.want)
		}
	}
}
