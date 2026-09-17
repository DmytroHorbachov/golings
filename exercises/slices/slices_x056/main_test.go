// slices_x056: Топ-K без изменения входа
// Make the tests pass!
// I AM NOT DONE
//
// topK возвращает k наибольших чисел по убыванию, не меняя входной срез.
// Тренирует: копирование перед сортировкой и срез результата.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func topK(s []int, k int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	return s[:k]
}

func TestTopK(t *testing.T) {
	in := []int{5, 1, 9, 3, 7}
	if got := topK(in, 3); !reflect.DeepEqual(got, []int{9, 7, 5}) {
		t.Errorf("topK = %v", got)
	}
	if !reflect.DeepEqual(in, []int{5, 1, 9, 3, 7}) {
		t.Errorf("input modified: %v", in)
	}
	if got := topK([]int{2}, 5); !reflect.DeepEqual(got, []int{2}) {
		t.Errorf("topK(k>len) = %v", got)
	}
}
