// algorithms45
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: стек. Проверьте, что скобки ()[]{} в строке правильно вложены
// и закрыты. Другие символы не встречаются.
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import "testing"

func isValid(s string) bool {
	return false
}

func TestIsValid(t *testing.T) {
	cases := map[string]bool{"()": true, "()[]{}": true, "(]": false, "([)]": false, "{[]}": true, "": true, "(": false, "]": false}
	for in, want := range cases {
		if got := isValid(in); got != want {
			t.Errorf("isValid(%q) = %v, want %v", in, got, want)
		}
	}
}
