// slices62
// Make the tests pass!

// I AM NOT DONE
//
// csv склеивает поля через запятую.
// Тренирует: strings.Join для []string.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func csv(fields []string) string {
	return strings.Join(fields, ", ")
}

func TestCSV(t *testing.T) {
	if got := csv([]string{"a", "b", "c"}); got != "a,b,c" {
		t.Errorf("csv = %q", got)
	}
}
