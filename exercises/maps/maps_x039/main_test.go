// maps_x039: Обращение с повторами
// Make the tests pass!
// I AM NOT DONE
//
// byValue группирует ключи по значению (значения могут повторяться);
// списки ключей отсортированы.
// Тренирует: map[V][]K и сортировку значений map.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byValue(m map[string]int) map[int][]string {
	out := map[int][]string{}
	for k, v := range m {
		out[v] = []string{k}
	}
	return out
}

func TestByValue(t *testing.T) {
	got := byValue(map[string]int{"a": 1, "b": 2, "c": 1, "d": 1})
	want := map[int][]string{1: {"a", "c", "d"}, 2: {"b"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("byValue = %v", got)
	}
}
