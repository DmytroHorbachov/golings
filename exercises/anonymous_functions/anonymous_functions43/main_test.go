// anonymous_functions43
// Make the tests pass!

// I AM NOT DONE
//
// strings.Map требует функцию func(rune) rune. Литерал объявлен с другим типом.
// Тренирует: сигнатура литерала должна точно совпадать с ожидаемой.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func upperA(s string) string {
	return strings.Map(func(r rune) byte {
		if r == 'a' {
			return 'A'
		}
		return r
	}, s)
}

func TestUpperA(t *testing.T) {
	if upperA("banana") != "bAnAnA" {
		t.Errorf("upperA = %q", upperA("banana"))
	}
}
