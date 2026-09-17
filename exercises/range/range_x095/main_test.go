// range_x095: append к копии строки
// Make the tests pass!
// I AM NOT DONE
//
// padRows дополняет каждую строку таблицы нулями до ширины w.
// Строки таблицы не меняются: append выполняется к переменной цикла.
// Тренирует: v в range по [][]T — копия заголовка среза.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func padRows(grid [][]int, w int) {
	for _, row := range grid {
		for len(row) < w {
			row = append(row, 0)
		}
	}
}

func TestPadRows(t *testing.T) {
	g := [][]int{{1}, {2, 3, 4}, {}}
	padRows(g, 3)
	if !reflect.DeepEqual(g, [][]int{{1, 0, 0}, {2, 3, 4}, {0, 0, 0}}) {
		t.Errorf("padRows = %v", g)
	}
}
