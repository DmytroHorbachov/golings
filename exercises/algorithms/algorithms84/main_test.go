// algorithms84
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: скользящее окно с упорядоченной структурой. Для каждого окна
// размера k верните медиану (для чётного k — среднее двух центральных).
// Сложность: hard. Ожидаемая асимптотика: O(n·k) по времени (или O(n log k) с кучами), O(k) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func medianSlidingWindow(nums []int, k int) []float64 {
	return nil
}

func TestMedianSlidingWindow(t *testing.T) {
	_ = sort.Ints
	cases := []struct {
		nums []int
		k    int
		want []float64
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []float64{1, -1, -1, 3, 5, 6}},
		{[]int{1, 2, 3, 4}, 2, []float64{1.5, 2.5, 3.5}},
		{[]int{5}, 1, []float64{5}},
		{[]int{2147483647, 2147483647}, 2, []float64{2147483647}},
		{[]int{1}, 2, nil},
	}
	for _, c := range cases {
		if got := medianSlidingWindow(c.nums, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("medianSlidingWindow(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
}
