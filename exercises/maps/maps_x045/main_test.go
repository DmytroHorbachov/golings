// maps_x045: Ключи по убыванию значения
// Make the tests pass!
// I AM NOT DONE
//
// rank возвращает ключи, отсортированные по убыванию значения,
// при равенстве — по алфавиту.
// Тренирует: сортировку ключей по значениям map.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func rank(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

func TestRank(t *testing.T) {
	got := rank(map[string]int{"ann": 3, "bob": 5, "cid": 3, "dan": 1})
	if !reflect.DeepEqual(got, []string{"bob", "ann", "cid", "dan"}) {
		t.Errorf("rank = %v", got)
	}
}
