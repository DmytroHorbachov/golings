// range_x036: continue с меткой
// Make the tests pass!
// I AM NOT DONE
//
// validRows возвращает индексы строк без отрицательных чисел.
// Тренирует: continue с меткой внешнего цикла.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func validRows(g [][]int) []int {
	var out []int
	for i, row := range g {
		for _, v := range row {
			if v < 0 {
				continue
			}
		}
		out = append(out, i)
	}
	return out
}

func TestValidRows(t *testing.T) {
	got := validRows([][]int{{1, 2}, {3, -1}, {}, {-5}})
	if !reflect.DeepEqual(got, []int{0, 2}) {
		t.Errorf("validRows = %v", got)
	}
}
