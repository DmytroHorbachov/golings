// algorithms_x005: Top K Frequent Elements (k самых частых)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: подсчёт частот и bucket sort. Верните k самых частых элементов,
// отсортированных по убыванию частоты; при равной частоте — по возрастанию значения.
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени (блочная сортировка), O(n) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	return nil
}

func TestTopKFrequent(t *testing.T) {
	_ = sort.Ints
	cases := []struct {
		nums []int
		k    int
		want []int
	}{
		{[]int{1, 1, 1, 2, 2, 3}, 2, []int{1, 2}},
		{[]int{1}, 1, []int{1}},
		{[]int{4, 4, 5, 5, 6}, 2, []int{4, 5}},
		{[]int{3, 3, 2, 2, 1, 1}, 3, []int{1, 2, 3}},
		{[]int{-1, -1, 7}, 1, []int{-1}},
	}
	for _, c := range cases {
		if got := topKFrequent(c.nums, c.k); !reflect.DeepEqual(got, c.want) {
			t.Errorf("topKFrequent(%v, %d) = %v, want %v", c.nums, c.k, got, c.want)
		}
	}
}
