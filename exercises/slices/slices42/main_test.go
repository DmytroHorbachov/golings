// slices42
// Make the tests pass!

// I AM NOT DONE
//
// removeAll возвращает новый срез без всех вхождений x, не меняя исходный.
// Тренирует: фильтрацию в новый срез.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func removeAll(s []int, x int) []int {
	for i, v := range s {
		if v == x {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func TestRemoveAll(t *testing.T) {
	in := []int{1, 2, 1, 3, 1}
	if got := removeAll(in, 1); !reflect.DeepEqual(got, []int{2, 3}) {
		t.Errorf("removeAll = %v", got)
	}
	if !reflect.DeepEqual(in, []int{1, 2, 1, 3, 1}) {
		t.Errorf("input modified: %v", in)
	}
}
