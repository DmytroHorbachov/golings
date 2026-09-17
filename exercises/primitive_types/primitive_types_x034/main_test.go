// primitive_types_x034: Размер алфавита
// Make the tests pass!
// I AM NOT DONE
//
// alphabetSize должна вернуть число букв от 'a' до 'z' включительно.
// Тренирует: арифметику над rune-константами.
// Сложность: easy
package main_test

import "testing"

func alphabetSize() int {
	return int('z' - 'a')
}

func TestAlphabetSize(t *testing.T) {
	if got := alphabetSize(); got != 26 {
		t.Errorf("alphabetSize() = %d, want 26", got)
	}
}
