// maps_x030: Обращение map
// Make the tests pass!
// I AM NOT DONE
//
// invert меняет местами ключи и значения (значения уникальны).
// Тренирует: запись в новую map в цикле.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func invert(m map[string]int) map[int]string {
	out := make(map[int]string, len(m))
	for k, v := range m {
		out[len(k)] = k
	}
	return out
}

func TestInvert(t *testing.T) {
	got := invert(map[string]int{"one": 1, "two": 2})
	if !reflect.DeepEqual(got, map[int]string{1: "one", 2: "two"}) {
		t.Errorf("invert = %v", got)
	}
}
