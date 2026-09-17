// structs_x065: Сумма заказа
// Make the tests pass!
// I AM NOT DONE
//
// Order.Total считает сумму позиций и применяет скидку в процентах (0–100).
// Тренирует: методы, использующие вложенные структуры.
// Сложность: medium
package main_test

import "testing"

type Line struct{ Price, Qty int }

type Order struct {
	Lines    []Line
	Discount int
}

func (o Order) Total() int {
	total := 0
	for _, l := range o.Lines {
		total += l.Price
	}
	return total - o.Discount
}

func TestOrderTotal(t *testing.T) {
	o := Order{[]Line{{100, 2}, {50, 4}}, 10}
	if got := o.Total(); got != 360 {
		t.Errorf("Total = %d, want 360", got)
	}
}
