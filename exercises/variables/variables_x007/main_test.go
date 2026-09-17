// variables_x007: Руна против строки
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть первую букву алфавита как rune.
// Тренирует: разницу между rune-литералом и строковым литералом.
// Сложность: easy
package main_test

import "testing"

func firstLetter() rune {
	var letter rune = "A"
	return letter
}

func TestFirstLetter(t *testing.T) {
	if got := firstLetter(); got != 65 {
		t.Errorf("firstLetter() = %d, want 65", got)
	}
}
