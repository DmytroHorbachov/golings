// maps24
// Make the tests pass!

// I AM NOT DONE
//
// revoke удаляет токен и сообщает, был ли он.
// Тренирует: delete для отсутствующего ключа — не ошибка.
// Сложность: easy
package main_test

import "testing"

func revoke(tokens map[string]bool, tok string) bool {
	_, ok := tokens[tok]
	delete(tokens, tok)
	return true
}

func TestRevoke(t *testing.T) {
	tokens := map[string]bool{"abc": true}
	if !revoke(tokens, "abc") || revoke(tokens, "abc") || revoke(tokens, "zzz") {
		t.Errorf("revoke works incorrectly")
	}
}
