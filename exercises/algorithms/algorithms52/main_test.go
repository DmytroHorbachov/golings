// algorithms52
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: разделяй и властвуй со слиянием. Верните новый отсортированный срез,
// не меняя исходный и не используя пакет sort.
// Сложность: medium. Ожидаемая асимптотика: O(n·log n) по времени, O(n) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func mergeSort(nums []int) []int {
	return nil
}

func TestMergeSort(t *testing.T) {
	_ = sort.Ints
	cases := [][]int{
		{5, 2, 9, 1, 5, 6},
		{},
		{7},
		{2, 1},
		{-3, -1, -2},
	}
	for _, c := range cases {
		in := append([]int{}, c...)
		got := mergeSort(in)
		want := append([]int{}, c...)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("mergeSort(%v) = %v, want %v", c, got, want)
		}
		if !reflect.DeepEqual(in, c) {
			t.Errorf("input modified: %v", in)
		}
	}
}
