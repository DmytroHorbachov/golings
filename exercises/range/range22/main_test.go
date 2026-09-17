// range22
// Make the tests pass!

// I AM NOT DONE
//
// findAll возвращает координаты (строка, столбец) всех клеток со значением x.
// Тренирует: два уровня range с индексами.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func findAll(g [][]int, x int) [][2]int {
	var out [][2]int
	for r, row := range g {
		for c := range row {
			if row[c] == x {
				return append(out, [2]int{c, r})
			}
		}
	}
	return out
}

func TestFindAll(t *testing.T) {
	got := findAll([][]int{{1, 0}, {0, 1, 1}}, 1)
	if !reflect.DeepEqual(got, [][2]int{{0, 0}, {1, 1}, {1, 2}}) {
		t.Errorf("findAll = %v", got)
	}
}
