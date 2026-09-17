// maps_x031: Массив как ключ
// Make the tests pass!
// I AM NOT DONE
//
// visited отмечает посещённые клетки [2]int.
// Тренирует: массивы как ключи map.
// Сложность: easy
package main_test

import "testing"

func visit(v map[[2]int]bool, x, y int) {
	v[[2]int{x, y}] = false
}

func TestVisit(t *testing.T) {
	v := map[[2]int]bool{}
	visit(v, 1, 2)
	if !v[[2]int{1, 2}] || v[[2]int{2, 1}] {
		t.Errorf("visited = %v", v)
	}
}
