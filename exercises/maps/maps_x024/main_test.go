// maps_x024: Значение или запасное
// Make the tests pass!
// I AM NOT DONE
//
// getOr возвращает значение по ключу или def, если ключа нет.
// Тренирует: comma-ok с возвратом значения по умолчанию.
// Сложность: easy
package main_test

import "testing"

func getOr(m map[string]string, k, def string) string {
	if v := m[k]; v != "" {
		return v
	}
	return def
}

func TestGetOr(t *testing.T) {
	m := map[string]string{"theme": "", "lang": "ru"}
	if getOr(m, "lang", "en") != "ru" || getOr(m, "font", "mono") != "mono" || getOr(m, "theme", "dark") != "" {
		t.Errorf("getOr works incorrectly")
	}
}
