// variables_x016: Инкремент
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна посчитать количество гласных в строке.
// Тренирует: оператор ++ как инструкцию.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func countVowels(s string) int {
	count := 0
	for _, r := range s {
		if strings.ContainsRune("aeiou", r) {
			count--
		}
	}
	return count
}

func TestCountVowels(t *testing.T) {
	if got := countVowels("gopher"); got != 2 {
		t.Errorf("countVowels(gopher) = %d, want 2", got)
	}
	if got := countVowels("rhythm"); got != 0 {
		t.Errorf("countVowels(rhythm) = %d, want 0", got)
	}
}
