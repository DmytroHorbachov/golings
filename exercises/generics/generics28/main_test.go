// generics28
// Make the tests pass!

// I AM NOT DONE
//
// Flatten превращает [][]T в []T.
// Тренирует: обобщённые вложенные срезы.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func Flatten[T any](parts [][]T) []T {
	n := 0
	for _, p := range parts {
		n++
	}
	out := make([]T, n)
	for _, p := range parts {
		out = append(out, p...)
	}
	return out
}

func TestFlatten(t *testing.T) {
	if got := Flatten([][]int{{1}, {}, {2, 3}}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("Flatten = %v", got)
	}
}
