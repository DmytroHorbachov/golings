// maps_x033: Поиск по значению
// Make the tests pass!
// I AM NOT DONE
//
// hasOwner проверяет, принадлежит ли хоть один файл пользователю.
// Тренирует: перебор значений map.
// Сложность: easy
package main_test

import "testing"

func hasOwner(files map[string]string, user string) bool {
	for name, owner := range files {
		_ = name
		if name == user {
			return true
		}
	}
	return false
}

func TestHasOwner(t *testing.T) {
	files := map[string]string{"a.txt": "ann", "b.txt": "bob"}
	if !hasOwner(files, "bob") || hasOwner(files, "a.txt") {
		t.Errorf("hasOwner works incorrectly")
	}
}
