// range_x075: break во вложенном цикле
// Make the tests pass!
// I AM NOT DONE
//
// contains2D ищет значение в таблице и считает, сколько клеток просмотрено.
// После находки обход должен полностью прекратиться.
// Тренирует: break выходит только из ближайшего цикла.
// Сложность: hard
package main_test

import "testing"

func contains2D(g [][]int, x int) (bool, int) {
	found, visited := false, 0
	for _, row := range g {
		for _, v := range row {
			visited++
			if v == x {
				found = true
				break
			}
		}
	}
	return found, visited
}

func TestContains2D(t *testing.T) {
	found, visited := contains2D([][]int{{1, 2}, {3, 4}, {5, 6}}, 2)
	if !found || visited != 2 {
		t.Errorf("contains2D = %v, %d; want true, 2", found, visited)
	}
}
