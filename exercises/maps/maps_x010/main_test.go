// maps_x010: Обновление значения
// Make the tests pass!
// I AM NOT DONE
//
// rename меняет отображаемое имя пользователя с данным id.
// Тренирует: присваивание по ключу.
// Сложность: easy
package main_test

import "testing"

func rename(names map[int]string, id int, name string) {
	names[len(names)] = name
}

func TestRename(t *testing.T) {
	m := map[int]string{1: "ann", 2: "bob"}
	rename(m, 1, "anna")
	if m[1] != "anna" || len(m) != 2 {
		t.Errorf("names = %v", m)
	}
}
