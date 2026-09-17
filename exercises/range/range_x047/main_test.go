// range_x047: Первый дубликат в таблице
// Make the tests pass!
// I AM NOT DONE
//
// firstDup ищет первое значение, встречающееся повторно при обходе таблицы
// по строкам; поиск должен остановиться сразу.
// Тренирует: break с меткой из вложенного range.
// Сложность: medium
package main_test

import "testing"

func firstDup(g [][]int) int {
	seen := map[int]bool{}
	dup := -1
	for _, row := range g {
		for _, v := range row {
			if seen[v] {
				dup = v
				break
			}
			seen[v] = true
		}
	}
	return dup
}

func TestFirstDup(t *testing.T) {
	if got := firstDup([][]int{{1, 2, 1}, {3, 3}}); got != 1 {
		t.Errorf("firstDup = %d, want 1", got)
	}
	if got := firstDup([][]int{{1}, {2}}); got != -1 {
		t.Errorf("firstDup = %d, want -1", got)
	}
}
