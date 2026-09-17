// algorithms_x099: Swim in Rising Water (плавание при подъёме воды)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: мин-куча (или двоичный поиск с обходом). На карте высот вода
// поднимается со временем; в момент t можно плавать по клеткам с высотой не выше t.
// Верните минимальное время, чтобы добраться из (0,0) в правый нижний угол.
// Сложность: hard. Ожидаемая асимптотика: O(r·c·log(r·c)) по времени, O(r·c) по памяти
package main_test

import (
	"container/heap"
	"testing"
)

type wcell struct{ h, r, c int }

type wheap []wcell

func (x wheap) Len() int            { return len(x) }
func (x wheap) Less(i, j int) bool  { return x[i].h < x[j].h }
func (x wheap) Swap(i, j int)       { x[i], x[j] = x[j], x[i] }
func (x *wheap) Push(v interface{}) { *x = append(*x, v.(wcell)) }
func (x *wheap) Pop() interface{} {
	old := *x
	v := old[len(old)-1]
	*x = old[:len(old)-1]
	return v
}

func swimInWater(grid [][]int) int {
	_ = heap.Init
	return 0
}

func TestSwimInWater(t *testing.T) {
	cases := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{0, 2}, {1, 3}}, 3},
		{[][]int{{0, 1, 2, 3, 4}, {24, 23, 22, 21, 5}, {12, 13, 14, 15, 16}, {11, 17, 18, 19, 20}, {10, 9, 8, 7, 6}}, 16},
		{[][]int{{0}}, 0},
		{nil, 0},
	}
	for _, c := range cases {
		if got := swimInWater(c.grid); got != c.want {
			t.Errorf("swimInWater = %d, want %d", got, c.want)
		}
	}
}
