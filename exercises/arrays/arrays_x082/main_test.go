// arrays_x082: Неквадратная матрица
// Make the tests pass!
// I AM NOT DONE
//
// sumColumns суммирует столбцы матрицы 2×3 (2 строки, 3 столбца).
// Индексы строк и столбцов перепутаны.
// Тренирует: len(m) и len(m[0]) у неквадратных массивов.
// Сложность: hard
package main_test

import "testing"

func sumColumns(m [2][3]int) [3]int {
	var out [3]int
	for i := 0; i < len(m[0]); i++ {
		for j := 0; j < len(m); j++ {
			out[j] += m[i][j]
		}
	}
	return out
}

func TestSumColumns(t *testing.T) {
	_ = sumColumns
	m := [2][3]int{{1, 2, 3}, {10, 20, 30}}
	if got := sumColumns(m); got != [3]int{11, 22, 33} {
		t.Errorf("sumColumns = %v", got)
	}
}
