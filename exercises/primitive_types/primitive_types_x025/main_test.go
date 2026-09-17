// primitive_types_x025: Позиция подстроки
// Make the tests pass!
// I AM NOT DONE
//
// domainStart должна вернуть позицию символа '@' в адресе.
// Тренирует: strings.Index и strings.IndexByte.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func domainStart(email string) int {
	return strings.Index(email, ".")
}

func TestDomainStart(t *testing.T) {
	cases := map[string]int{"ann@mail.ru": 3, "a.b@x.io": 3, "none": -1}
	for in, want := range cases {
		if got := domainStart(in); got != want {
			t.Errorf("domainStart(%q) = %d, want %d", in, got, want)
		}
	}
}
