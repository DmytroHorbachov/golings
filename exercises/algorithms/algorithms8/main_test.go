// algorithms8
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: алгоритм Дейкстры. Рёбра [from, to, weight] задают ориентированный
// граф на вершинах 1..n. Верните время, за которое сигнал из вершины k дойдёт
// до всех, или -1.
// Сложность: hard. Ожидаемая асимптотика: O(E·log V) по времени, O(V + E) по памяти
package main_test

import (
	"container/heap"
	"testing"
)

type qItem struct{ node, dist int }

type pq []qItem

func (p pq) Len() int            { return len(p) }
func (p pq) Less(i, j int) bool  { return p[i].dist < p[j].dist }
func (p pq) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x interface{}) { *p = append(*p, x.(qItem)) }
func (p *pq) Pop() interface{} {
	old := *p
	v := old[len(old)-1]
	*p = old[:len(old)-1]
	return v
}

func networkDelayTime(times [][3]int, n, k int) int {
	_ = heap.Init
	return 0
}

func TestNetworkDelayTime(t *testing.T) {
	cases := []struct {
		times [][3]int
		n, k  int
		want  int
	}{
		{[][3]int{{2, 1, 1}, {2, 3, 1}, {3, 4, 1}}, 4, 2, 2},
		{[][3]int{{1, 2, 1}}, 2, 1, 1},
		{[][3]int{{1, 2, 1}}, 2, 2, -1},
		{nil, 1, 1, 0},
		{[][3]int{{1, 2, 5}, {1, 3, 2}, {3, 2, 1}}, 3, 1, 3},
	}
	for _, c := range cases {
		if got := networkDelayTime(c.times, c.n, c.k); got != c.want {
			t.Errorf("networkDelayTime(%v, %d, %d) = %d, want %d", c.times, c.n, c.k, got, c.want)
		}
	}
}
