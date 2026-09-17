// slices58
// Make the tests pass!

// I AM NOT DONE
//
// sortedUnique возвращает отсортированный срез уникальных строк.
// Тренирует: sort.Strings и удаление соседних дубликатов.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func sortedUnique(s []string) []string {
	c := append([]string(nil), s...)
	sort.Strings(c)
	var out []string
	for _, v := range c {
		if len(out) == 0 || out[0] != v {
			out = append(out, v)
		}
	}
	return c
}

func TestSortedUnique(t *testing.T) {
	got := sortedUnique([]string{"b", "a", "b", "c", "a", "c"})
	if !reflect.DeepEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedUnique = %v", got)
	}
}
