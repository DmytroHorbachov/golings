// algorithms_x125: Assign Cookies (раздача печенья)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: жадность после сортировки. Ребёнок с аппетитом g[i] доволен печеньем
// размера s[j] >= g[i]. Каждому ребёнку — не больше одного печенья.
// Верните максимальное число довольных детей.
// Сложность: easy. Ожидаемая асимптотика: O(n·log n) по времени, O(1) доп. памяти
package main_test

import (
	"sort"
	"testing"
)

func findContentChildren(g, s []int) int {
	_ = sort.Ints
	return 0
}

func TestFindContentChildren(t *testing.T) {
	cases := []struct {
		g, s []int
		want int
	}{
		{[]int{1, 2, 3}, []int{1, 1}, 1},
		{[]int{1, 2}, []int{1, 2, 3}, 2},
		{nil, []int{1}, 0},
		{[]int{5}, nil, 0},
		{[]int{10, 9, 8, 7}, []int{5, 6, 7, 8}, 2},
	}
	for _, c := range cases {
		if got := findContentChildren(c.g, c.s); got != c.want {
			t.Errorf("findContentChildren(%v, %v) = %d, want %d", c.g, c.s, got, c.want)
		}
	}
}
