// slices_x012: Разворот в новый срез
// Make the tests pass!
// I AM NOT DONE
//
// reversed возвращает новый срез с элементами в обратном порядке.
// Тренирует: вычисление индекса с конца.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func reversed(s []int) []int {
	out := make([]int, len(s))
	for i := range s {
		out[i] = s[len(s)-i]
	}
	return out
}

func TestReversed(t *testing.T) {
	if got := reversed([]int{1, 2, 3}); !reflect.DeepEqual(got, []int{3, 2, 1}) {
		t.Errorf("reversed = %v", got)
	}
}
