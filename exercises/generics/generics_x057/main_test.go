// generics_x057: Слияние map
// Make the tests pass!
// I AM NOT DONE
//
// Merge объединяет map; при совпадении ключей используется resolve.
// Тренирует: обобщённые map и функция разрешения конфликтов.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func Merge[K comparable, V any](a, b map[K]V, resolve func(V, V) V) map[K]V {
	out := make(map[K]V, len(a))
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func TestMerge(t *testing.T) {
	got := Merge(map[string]int{"a": 1, "b": 2}, map[string]int{"b": 10, "c": 3}, func(x, y int) int { return x + y })
	if !reflect.DeepEqual(got, map[string]int{"a": 1, "b": 12, "c": 3}) {
		t.Errorf("Merge = %v", got)
	}
}
