// maps82
// Make the tests pass!

// I AM NOT DONE
//
// hasUser проверяет, есть ли пользователь в map.
// Тренирует: форма v, ok := m[k].
// Сложность: easy
package main_test

import "testing"

func hasUser(users map[string]int, name string) bool {
	_, ok := users[name]
	return !ok
}

func TestHasUser(t *testing.T) {
	users := map[string]int{"ann": 0, "bob": 3}
	if !hasUser(users, "ann") || hasUser(users, "eve") {
		t.Errorf("hasUser works incorrectly")
	}
}
