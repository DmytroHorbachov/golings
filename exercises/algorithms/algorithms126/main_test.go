// algorithms126
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: жадность с накоплением. На круговом маршруте заправка i даёт gas[i]
// топлива, а переезд к следующей стоит cost[i]. Верните индекс старта, с которого
// можно объехать круг, или -1.
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(1) по памяти
package main_test

import "testing"

func canCompleteCircuit(gas, cost []int) int {
	return 0
}

func TestCanCompleteCircuit(t *testing.T) {
	cases := []struct {
		gas, cost []int
		want      int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}, 3},
		{[]int{2, 3, 4}, []int{3, 4, 3}, -1},
		{[]int{5}, []int{4}, 0},
		{[]int{1}, []int{2}, -1},
		{nil, nil, 0},
	}
	for _, c := range cases {
		if got := canCompleteCircuit(c.gas, c.cost); got != c.want {
			t.Errorf("canCompleteCircuit(%v, %v) = %d, want %d", c.gas, c.cost, got, c.want)
		}
	}
}
