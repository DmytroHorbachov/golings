// arrays93
// Make the tests pass!

// I AM NOT DONE
//
// visits считает посещения клеток, используя координаты [2]int как ключ.
// Тренирует: массивы как ключи map.
// Сложность: easy
package main_test

import "testing"

func visits(path [][2]int) map[[2]int]int {
	m := map[[2]int]int{}
	for _, p := range path {
		m[p] = 1
	}
	return m
}

func TestVisits(t *testing.T) {
	m := visits([][2]int{{0, 0}, {1, 0}, {0, 0}})
	if m[[2]int{0, 0}] != 2 || m[[2]int{1, 0}] != 1 {
		t.Errorf("visits = %v", m)
	}
}
