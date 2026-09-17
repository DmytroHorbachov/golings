// primitive_types_x048: Шифр Атбаш
// Make the tests pass!
// I AM NOT DONE
//
// atbash заменяет латинскую букву зеркальной: a<->z, b<->y; регистр сохраняется,
// прочие символы не меняются.
// Тренирует: арифметику над рунами для двух диапазонов.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

func mirror(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return 'a' + (r - 'z')
	}
	return r
}

func atbash(s string) string {
	return strings.Map(mirror, s)
}

func TestAtbash(t *testing.T) {
	cases := map[string]string{"abc": "zyx", "Hello!": "Svool!", "Zz 9": "Aa 9"}
	for in, want := range cases {
		if got := atbash(in); got != want {
			t.Errorf("atbash(%q) = %q, want %q", in, got, want)
		}
	}
}
