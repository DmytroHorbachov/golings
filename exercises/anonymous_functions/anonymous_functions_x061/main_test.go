// anonymous_functions_x061: Поэлементная операция над матрицей
// Make the tests pass!
// I AM NOT DONE
//
// apply применяет литерал к каждому элементу матрицы, передавая координаты,
// и возвращает новую матрицу.
// Тренирует: литерал с несколькими параметрами и новые срезы.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func apply(m [][]int, f func(v, r, c int) int) [][]int {
	for r := range m {
		for c := range m[r] {
			m[r][c] = f(m[r][c], c, r)
		}
	}
	return m
}

func TestApply(t *testing.T) {
	m := [][]int{{1, 2}, {3, 4}, {5, 6}}
	got := apply(m, func(v, r, c int) int { return v * (r + 1) })
	if !reflect.DeepEqual(got, [][]int{{1, 2}, {6, 8}, {15, 18}}) {
		t.Errorf("apply = %v", got)
	}
	if m[1][0] != 3 {
		t.Errorf("input modified: %v", m)
	}
}
