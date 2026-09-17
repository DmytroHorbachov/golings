// primitive_types_x011: Замена всех вхождений
// Make the tests pass!
// I AM NOT DONE
//
// dashes должна заменить все пробелы на дефисы.
// Тренирует: strings.Replace и strings.ReplaceAll.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func dashes(s string) string {
	return strings.Replace(s, " ", "-", 1)
}

func TestDashes(t *testing.T) {
	if got := dashes("a b c d"); got != "a-b-c-d" {
		t.Errorf("dashes = %q, want a-b-c-d", got)
	}
}
