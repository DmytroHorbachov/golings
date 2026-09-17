// arrays_x019: Уникальные буквы
// Make the tests pass!
// I AM NOT DONE
//
// allUnique проверяет, что в строке из строчных латинских букв нет повторов.
// Тренирует: массив [26]bool как множество.
// Сложность: easy
package main_test

import "testing"

func allUnique(s string) bool {
	var seen [26]bool
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if seen[c] {
			return false
		}
		seen[c] = false
	}
	return true
}

func TestAllUnique(t *testing.T) {
	cases := map[string]bool{"abc": true, "hello": false, "": true, "gopher": true}
	for in, want := range cases {
		if got := allUnique(in); got != want {
			t.Errorf("allUnique(%q) = %v, want %v", in, got, want)
		}
	}
}
