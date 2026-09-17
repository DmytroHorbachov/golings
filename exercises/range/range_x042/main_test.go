// range_x042: Нарастающий максимум
// Make the tests pass!
// I AM NOT DONE
//
// runningMax возвращает для каждой позиции максимум среди элементов до неё включительно.
// Тренирует: аккумулятор в range и запись по индексу.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func runningMax(s []int) []int {
	out := make([]int, len(s))
	m := 0
	for i, v := range s {
		if v > m {
			out[i] = v
		}
	}
	return out
}

func TestRunningMax(t *testing.T) {
	if got := runningMax([]int{-3, -5, 2, 1, 4}); !reflect.DeepEqual(got, []int{-3, -3, 2, 2, 4}) {
		t.Errorf("runningMax = %v", got)
	}
}
