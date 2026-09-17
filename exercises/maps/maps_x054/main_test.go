// maps_x054: Настройки по умолчанию
// Make the tests pass!
// I AM NOT DONE
//
// effective накладывает пользовательские настройки на значения по умолчанию,
// не меняя ни одну из входных map.
// Тренирует: копирование map и перезапись ключей.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func effective(defaults, user map[string]string) map[string]string {
	out := defaults
	for k, v := range user {
		out[k] = v
	}
	return out
}

func TestEffective(t *testing.T) {
	def := map[string]string{"theme": "light", "lang": "en"}
	got := effective(def, map[string]string{"theme": "dark"})
	if !reflect.DeepEqual(got, map[string]string{"theme": "dark", "lang": "en"}) {
		t.Errorf("effective = %v", got)
	}
	if def["theme"] != "light" {
		t.Errorf("defaults modified: %v", def)
	}
}
