// if_x061: Флаг из строки
// Make the tests pass!
// I AM NOT DONE
//
// parseSwitch понимает "on"/"yes"/"1" как true и "off"/"no"/"0" как false
// (без учёта регистра). Для остального возвращает def.
// Тренирует: нормализацию входа перед сравнениями.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func parseSwitch(s string, def bool) bool {
	if s == "on" || s == "yes" || s == "1" {
		return true
	}
	return false
}

func TestParseSwitch(t *testing.T) {
	_ = strings.ToLower
	cases := []struct {
		in   string
		def  bool
		want bool
	}{{"ON", false, true}, {"yes", false, true}, {"No", true, false}, {"0", true, false}, {"maybe", true, true}, {"", false, false}}
	for _, c := range cases {
		if got := parseSwitch(c.in, c.def); got != c.want {
			t.Errorf("parseSwitch(%q, %v) = %v, want %v", c.in, c.def, got, c.want)
		}
	}
}
