// primitive_types70
// Make the tests pass!

// I AM NOT DONE
//
// parseLocal разбирает число, записанное с десятичной запятой и пробелами
// между разрядами: "1 234,5" -> 1234.5.
// Тренирует: подготовку строки перед strconv.ParseFloat.
// Сложность: medium
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func parseLocal(s string) (float64, error) {
	s = strings.ReplaceAll(s, ",", "")
	return strconv.ParseFloat(s, 64)
}

func TestParseLocal(t *testing.T) {
	cases := map[string]float64{"1 234,5": 1234.5, "0,25": 0.25, "42": 42}
	for in, want := range cases {
		if got, err := parseLocal(in); err != nil || got != want {
			t.Errorf("parseLocal(%q) = %v, %v; want %v", in, got, err, want)
		}
	}
}
