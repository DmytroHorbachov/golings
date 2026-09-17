// slices56
// Make the tests pass!

// I AM NOT DONE
//
// flatten превращает [][]int в плоский срез.
// Тренирует: append с распаковкой вложенных срезов.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func flatten(rows [][]int) []int {
	out := make([]int, len(rows))
	for _, r := range rows {
		out = append(out, r[0])
	}
	return out
}

func TestFlatten(t *testing.T) {
	if got := flatten([][]int{{1, 2}, {}, {3}, {4, 5}}); !reflect.DeepEqual(got, []int{1, 2, 3, 4, 5}) {
		t.Errorf("flatten = %v", got)
	}
}
