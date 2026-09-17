// maps_x014: Отсортированные ключи
// Make the tests pass!
// I AM NOT DONE
//
// sortedKeys возвращает ключи map в алфавитном порядке.
// Тренирует: сбор ключей и sort.Strings.
// Сложность: easy
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortedKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	return keys
}

func TestSortedKeys(t *testing.T) {
	if got := sortedKeys(map[string]int{"b": 1, "c": 2, "a": 3}); !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedKeys = %v", got)
	}
}
