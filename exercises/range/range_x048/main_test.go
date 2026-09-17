// range_x048: Слияние с приоритетом
// Make the tests pass!
// I AM NOT DONE
//
// mergeMax объединяет map, оставляя для совпадающих ключей большее значение.
// Тренирует: range по нескольким map и comma-ok.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func mergeMax(a, b map[string]int) map[string]int {
	out := map[string]int{}
	for k, v := range a {
		out[k] = v
	}
	for k, v := range b {
		out[k] = v
	}
	return out
}

func TestMergeMax(t *testing.T) {
	got := mergeMax(map[string]int{"x": 5, "y": 1}, map[string]int{"x": 2, "y": 3, "z": -1})
	if !reflect.DeepEqual(got, map[string]int{"x": 5, "y": 3, "z": -1}) {
		t.Errorf("mergeMax = %v", got)
	}
}
