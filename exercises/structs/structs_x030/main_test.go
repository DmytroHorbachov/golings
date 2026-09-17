// structs_x030: Сумма полей
// Make the tests pass!
// I AM NOT DONE
//
// total складывает количество товаров в заказе.
// Тренирует: накопление значения поля.
// Сложность: easy
package main_test

import "testing"

type Line struct {
	Price, Qty int
}

type Order struct{ Lines []Line }

func (o Order) Items() int {
	n := 0
	for _, l := range o.Lines {
		n += l.Price
	}
	return n
}

func TestItems(t *testing.T) {
	o := Order{[]Line{{100, 2}, {50, 3}}}
	if o.Items() != 5 {
		t.Errorf("Items = %d", o.Items())
	}
}
