// algorithms_x028: Daily Temperatures (ожидание потепления)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: монотонный стек. Для каждого дня найдите, через сколько дней
// станет теплее; если такого дня нет — 0.
// Сложность: medium. Ожидаемая асимптотика: O(n) по времени, O(n) по памяти
package main_test

import (
	"reflect"
	"testing"
)

func dailyTemperatures(temps []int) []int {
	return nil
}

func TestDailyTemperatures(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 30, 30}, []int{0, 0, 0}},
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := dailyTemperatures(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("dailyTemperatures(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
