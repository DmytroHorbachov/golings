// slices_x061: Вставка с сохранением порядка
// Make the tests pass!
// I AM NOT DONE
//
// insertSorted вставляет число в отсортированный срез, сохраняя порядок.
// Тренирует: sort.SearchInts и вставку по индексу.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func insertSorted(s []int, x int) []int {
	i := sort.SearchInts(s, x)
	s = append(s, x)
	_ = i
	return s
}

func TestInsertSorted(t *testing.T) {
	s := []int{}
	for _, v := range []int{5, 1, 3, 4, 2} {
		s = insertSorted(s, v)
	}
	if !reflect.DeepEqual(s, []int{1, 2, 3, 4, 5}) {
		t.Errorf("insertSorted = %v", s)
	}
}
