// maps_x028: Флаги по умолчанию
// Make the tests pass!
// I AM NOT DONE
//
// enabled сообщает, включена ли функция; незаданные функции выключены.
// Тренирует: нулевое значение bool в map.
// Сложность: easy
package main_test

import "testing"

func enabled(flags map[string]bool, name string) bool {
	v, ok := flags[name]
	return ok
}

func TestEnabled(t *testing.T) {
	flags := map[string]bool{"beta": true, "legacy": false}
	if !enabled(flags, "beta") || enabled(flags, "legacy") || enabled(flags, "dark") {
		t.Errorf("enabled works incorrectly")
	}
}
