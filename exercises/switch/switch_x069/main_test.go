// switch_x069: fallthrough не проверяет условие
// Make the tests pass!
// I AM NOT DONE
//
// shipping: для заказов от 5000 доставка бесплатна; для заказов от 1000 — 100;
// иначе 300. fallthrough приводит к неверной цене.
// Тренирует: fallthrough передаёт управление в следующую ветку без проверки её условия.
// Сложность: hard
package main_test

import "testing"

func shipping(total int) int {
	cost := 300
	switch {
	case total >= 5000:
		cost = 0
		fallthrough
	case total >= 1000:
		cost = 100
	}
	return cost
}

func TestShipping(t *testing.T) {
	cases := map[int]int{6000: 0, 5000: 0, 1500: 100, 999: 300}
	for in, want := range cases {
		if got := shipping(in); got != want {
			t.Errorf("shipping(%d) = %d, want %d", in, got, want)
		}
	}
}
