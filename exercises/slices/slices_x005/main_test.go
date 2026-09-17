// slices_x005: Без первого элемента
// Make the tests pass!
// I AM NOT DONE
//
// tail возвращает все элементы, кроме первого.
// Тренирует: выражение среза s[low:].
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func tail(s []string) []string {
	return s[0:]
}

func TestTail(t *testing.T) {
	if got := tail([]string{"a", "b", "c"}); !reflect.DeepEqual(got, []string{"b", "c"}) {
		t.Errorf("tail = %v", got)
	}
}
