// variables_x098: Нулевой strings.Builder
// Make the tests pass!
// I AM NOT DONE
//
// Функция csvLine должна склеить поля через запятую.
// Нулевое значение strings.Builder готово к работе без инициализации.
// Тренирует: переменные с полезным нулевым значением.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func csvLine(fields []string) string {
	var b strings.Builder
	for _, f := range fields {
		b.WriteString(f)
		b.WriteString(",")
	}
	return b.String()
}

func TestCSVLine(t *testing.T) {
	cases := map[string][]string{"a,b,c": {"a", "b", "c"}, "x": {"x"}, "": nil}
	for want, in := range cases {
		if got := csvLine(in); got != want {
			t.Errorf("csvLine(%v) = %q, want %q", in, got, want)
		}
	}
}
