// algorithms53
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: k указателей и мин-куча. Из каждого отсортированного списка нужно
// взять хотя бы одно число. Найдите наименьший по длине отрезок [a, b],
// содержащий такие числа; при равной длине — с меньшим началом.
// Сложность: hard. Ожидаемая асимптотика: O(N·log k) по времени, O(k) по памяти
package main_test

import (
	"container/heap"
	"testing"
)

type item struct{ val, list, idx int }

type itemHeap []item

func (h itemHeap) Len() int            { return len(h) }
func (h itemHeap) Less(i, j int) bool  { return h[i].val < h[j].val }
func (h itemHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *itemHeap) Push(x interface{}) { *h = append(*h, x.(item)) }
func (h *itemHeap) Pop() interface{} {
	old := *h
	v := old[len(old)-1]
	*h = old[:len(old)-1]
	return v
}

func smallestRange(lists [][]int) [2]int {
	_ = heap.Init
	return [2]int{}
}

func TestSmallestRange(t *testing.T) {
	cases := []struct {
		lists [][]int
		want  [2]int
	}{
		{[][]int{{4, 10, 15, 24, 26}, {0, 9, 12, 20}, {5, 18, 22, 30}}, [2]int{20, 24}},
		{[][]int{{1, 2, 3}, {1, 2, 3}, {1, 2, 3}}, [2]int{1, 1}},
		{[][]int{{1}}, [2]int{1, 1}},
		{[][]int{{1, 5}, {6, 7}}, [2]int{5, 6}},
	}
	for _, c := range cases {
		if got := smallestRange(c.lists); got != c.want {
			t.Errorf("smallestRange(%v) = %v, want %v", c.lists, got, c.want)
		}
	}
}
