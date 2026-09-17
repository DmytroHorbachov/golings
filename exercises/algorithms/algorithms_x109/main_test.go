// algorithms_x109: Frog Jump (прыжки лягушки)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: ДП по состояниям (позиция, длина прыжка). Камни заданы возрастающими
// координатами. Первый прыжок — длины 1; после прыжка длины k следующий может
// быть k-1, k или k+1. Можно ли добраться до последнего камня?
// Сложность: hard. Ожидаемая асимптотика: O(n²) по времени, O(n²) по памяти
package main_test

import "testing"

func canCross(stones []int) bool {
	return false
}

func TestCanCross(t *testing.T) {
	cases := []struct {
		stones []int
		want   bool
	}{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
		{[]int{0, 1}, true},
		{[]int{0, 2}, false},
		{[]int{0}, true},
		{nil, false},
	}
	for _, c := range cases {
		if got := canCross(c.stones); got != c.want {
			t.Errorf("canCross(%v) = %v, want %v", c.stones, got, c.want)
		}
	}
}
