// slices_x037: Уникальные с сохранением порядка
// Make the tests pass!
// I AM NOT DONE
//
// unique удаляет повторы, сохраняя порядок первых вхождений.
// Тренирует: срез вместе с map-множеством.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func unique(s []string) []string {
	seen := map[string]bool{}
	var out []string
	for _, v := range s {
		if seen[v] {
			out = append(out, v)
		}
	}
	return out
}

func TestUnique(t *testing.T) {
	if got := unique([]string{"b", "a", "b", "c", "a"}); !reflect.DeepEqual(got, []string{"b", "a", "c"}) {
		t.Errorf("unique = %v", got)
	}
}
