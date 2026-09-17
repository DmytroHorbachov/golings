// slices_x019: Удаление по индексу
// Make the tests pass!
// I AM NOT DONE
//
// removeAt возвращает срез без элемента с индексом i (исходный срез можно менять).
// Тренирует: append(s[:i], s[i+1:]...).
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func removeAt(s []string, i int) []string {
	return append(s[:i], s[i:]...)
}

func TestRemoveAt(t *testing.T) {
	if got := removeAt([]string{"a", "b", "c"}, 1); !reflect.DeepEqual(got, []string{"a", "c"}) {
		t.Errorf("removeAt = %v", got)
	}
}
