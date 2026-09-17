// primitive_types105
// Make the tests pass!

// I AM NOT DONE
//
// parseQty разбирает количество из пользовательского ввода вроде " 42\n".
// strconv.Atoi не прощает пробелов.
// Тренирует: строгость функций strconv.
// Сложность: hard
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func parseQty(s string) (int, error) {
	return strconv.Atoi(s)
}

func TestParseQty(t *testing.T) {
	_ = strings.TrimSpace
	cases := map[string]int{" 42\n": 42, "7": 7, "\t-3 ": -3}
	for in, want := range cases {
		if got, err := parseQty(in); err != nil || got != want {
			t.Errorf("parseQty(%q) = %d, %v", in, got, err)
		}
	}
	if _, err := parseQty("4 2"); err == nil {
		t.Errorf("parseQty(\"4 2\") should fail")
	}
}
