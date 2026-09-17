// slices47
// Make the tests pass!

// I AM NOT DONE
//
// firstOr возвращает первый аргумент командной строки или значение по умолчанию.
// Без аргументов функция паникует.
// Тренирует: индексация пустого или nil-среза — паника.
// Сложность: hard
package main_test

import "testing"

func firstOr(args []string, def string) string {
	if args[0] == "" {
		return def
	}
	return args[0]
}

func TestFirstOr(t *testing.T) {
	if firstOr([]string{"run"}, "help") != "run" {
		t.Errorf("firstOr(run) failed")
	}
	if firstOr(nil, "help") != "help" || firstOr([]string{""}, "help") != "help" {
		t.Errorf("firstOr default failed")
	}
}
