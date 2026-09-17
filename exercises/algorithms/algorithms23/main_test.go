// algorithms23
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: двоичный поиск по ответу. Есть кучки бананов и h часов. За час можно
// съесть до k бананов из одной кучки. Найдите минимальную скорость k.
// Сложность: medium. Ожидаемая асимптотика: O(n·log m) по времени, O(1) по памяти
package main_test

import "testing"

func minEatingSpeed(piles []int, h int) int {
	return 0
}

func TestMinEatingSpeed(t *testing.T) {
	cases := []struct {
		piles   []int
		h, want int
	}{
		{[]int{3, 6, 7, 11}, 8, 4},
		{[]int{30, 11, 23, 4, 20}, 5, 30},
		{[]int{30, 11, 23, 4, 20}, 6, 23},
		{[]int{1}, 1, 1},
		{[]int{1000000000}, 2, 500000000},
	}
	for _, c := range cases {
		if got := minEatingSpeed(c.piles, c.h); got != c.want {
			t.Errorf("minEatingSpeed(%v, %d) = %d, want %d", c.piles, c.h, got, c.want)
		}
	}
}
