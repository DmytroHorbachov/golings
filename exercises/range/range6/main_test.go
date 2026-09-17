// range6
// Make the tests pass!

// I AM NOT DONE
//
// totalPrice считает стоимость корзины: цена * количество для каждой позиции.
// Тренирует: доступ к полям в range по срезу структур.
// Сложность: easy
package main_test

import "testing"

type Line struct {
	Price, Qty int
}

func totalPrice(lines []Line) int {
	sum := 0
	for _, l := range lines {
		sum += l.Price + l.Qty
	}
	return sum
}

func TestTotalPrice(t *testing.T) {
	if got := totalPrice([]Line{{100, 2}, {50, 3}}); got != 350 {
		t.Errorf("totalPrice = %d, want 350", got)
	}
}
