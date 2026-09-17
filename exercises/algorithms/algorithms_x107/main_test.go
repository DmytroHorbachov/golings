// algorithms_x107: Min Cost Climbing Stairs (минимальная стоимость подъёма)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: одномерное ДП. Каждая ступень стоит cost[i]; начинать можно
// с нулевой или первой ступени, шагать на 1 или 2. Верните минимальную
// стоимость подъёма на вершину (за последней ступенью).
// Сложность: easy. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func minCostClimbingStairs(cost []int) int {
	return 0
}

func TestMinCostClimbingStairs(t *testing.T) {
	cases := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
		{[]int{5, 5}, 5},
		{nil, 0},
		{[]int{7}, 0},
	}
	for _, c := range cases {
		if got := minCostClimbingStairs(c.cost); got != c.want {
			t.Errorf("minCostClimbingStairs(%v) = %d, want %d", c.cost, got, c.want)
		}
	}
}
