// switch_x002: Список значений в case
// Make the tests pass!
// I AM NOT DONE
//
// isVowel должна вернуть true для всех пяти гласных a, e, i, o, u.
// Тренирует: несколько значений в одной ветке case.
// Сложность: easy
package main_test

import "testing"

func isVowel(r rune) bool {
	switch r {
	case 'a', 'e', 'i', 'o':
		return true
	}
	return false
}

func TestIsVowel(t *testing.T) {
	for _, r := range "aeiou" {
		if !isVowel(r) {
			t.Errorf("isVowel(%q) = false", r)
		}
	}
	if isVowel('z') {
		t.Errorf("isVowel('z') = true")
	}
}
