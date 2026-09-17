// slices_x003: Склейка срезов
// Make the tests pass!
// I AM NOT DONE
//
// concat объединяет два среза.
// Тренирует: append(a, b...).
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func concat(a, b []int) []int {
	return append(a, b)
}

func TestConcat(t *testing.T) {
	if got := concat([]int{1, 2}, []int{3, 4}); !reflect.DeepEqual(got, []int{1, 2, 3, 4}) {
		t.Errorf("concat = %v", got)
	}
}
