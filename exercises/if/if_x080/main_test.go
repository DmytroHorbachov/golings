// if_x080: Двойное сравнение
// Make the tests pass!
// I AM NOT DONE
//
// isDigitCode должна проверить, что код символа между '0' и '9'.
// Код не компилируется: запись вида a < x < b в Go не работает.
// Тренирует: сравнения не объединяются в цепочки, как в математике.
// Сложность: hard
package main_test

import "testing"

func isDigitCode(c byte) bool {
	if '0' <= c <= '9' {
		return true
	}
	return false
}

func TestIsDigitCode(t *testing.T) {
	for _, c := range []byte("0459") {
		if !isDigitCode(c) {
			t.Errorf("isDigitCode(%q) = false", c)
		}
	}
	for _, c := range []byte("a/:") {
		if isDigitCode(c) {
			t.Errorf("isDigitCode(%q) = true", c)
		}
	}
}
