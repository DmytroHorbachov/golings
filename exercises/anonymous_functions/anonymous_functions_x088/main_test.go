// anonymous_functions_x088: 0 вместо -1 в strings.Map
// Make the tests pass!
// I AM NOT DONE
//
// stripVowels должна удалить гласные. Литерал возвращает 0, и в строке
// появляются нулевые символы.
// Тренирует: strings.Map удаляет символ только при отрицательном результате.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func stripVowels(s string) string {
	return strings.Map(func(r rune) rune {
		if strings.ContainsRune("aeiou", r) {
			return 0
		}
		return r
	}, s)
}

func TestStripVowels(t *testing.T) {
	if got := stripVowels("gopher"); got != "gphr" {
		t.Errorf("stripVowels = %q", got)
	}
}
