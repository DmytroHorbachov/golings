// anonymous_functions_x069: := внутри литерала
// Make the tests pass!
// I AM NOT DONE
//
// Литерал должен накапливать общую сумму, но сумма остаётся нулевой.
// Тренирует: := в теле литерала объявляет новую локальную переменную.
// Сложность: hard
package main_test

import "testing"

func sumAll(groups [][]int) int {
	total := 0
	add := func(vs []int) {
		for _, v := range vs {
			total := total + v
			_ = total
		}
	}
	for _, g := range groups {
		add(g)
	}
	return total
}

func TestSumAll(t *testing.T) {
	if got := sumAll([][]int{{1, 2}, {3}}); got != 6 {
		t.Errorf("sumAll = %d, want 6", got)
	}
}
