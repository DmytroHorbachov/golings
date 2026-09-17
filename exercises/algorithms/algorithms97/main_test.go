// algorithms97
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: одномерное ДП. Найдите минимальное количество монет для суммы
// amount (монет каждого номинала неограниченно) или -1.
// Сложность: medium. Ожидаемая асимптотика: O(n·amount) по времени, O(amount) по памяти
package main_test

import "testing"

func coinChange(coins []int, amount int) int {
	return 0
}

func TestCoinChange(t *testing.T) {
	cases := []struct {
		coins        []int
		amount, want int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
		{nil, 0, 0},
		{nil, 7, -1},
		{[]int{186, 419, 83, 408}, 6249, 20},
	}
	for _, c := range cases {
		if got := coinChange(c.coins, c.amount); got != c.want {
			t.Errorf("coinChange(%v, %d) = %d, want %d", c.coins, c.amount, got, c.want)
		}
	}
}
