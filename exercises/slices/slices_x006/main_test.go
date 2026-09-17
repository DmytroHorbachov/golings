// slices_x006: Без последнего элемента
// Make the tests pass!
// I AM NOT DONE
//
// pop возвращает срез без последнего элемента.
// Тренирует: выражение среза s[:high].
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func pop(s []int) []int {
	return s[:len(s)]
}

func TestPop(t *testing.T) {
	if got := pop([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{1, 2}) {
		t.Errorf("pop = %v", got)
	}
}
