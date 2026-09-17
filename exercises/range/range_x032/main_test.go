// range_x032: Пробелы
// Make the tests pass!
// I AM NOT DONE
//
// countSpaces считает пробелы в строке.
// Тренирует: сравнение руны в range.
// Сложность: easy
package main_test

import "testing"

func countSpaces(s string) int {
	n := 0
	for _, r := range s {
		if r == '_' {
			n++
		}
	}
	return n
}

func TestCountSpaces(t *testing.T) {
	if got := countSpaces("a b  c_d"); got != 3 {
		t.Errorf("countSpaces = %d, want 3", got)
	}
}
