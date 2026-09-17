// structs_x025: Map структур
// Make the tests pass!
// I AM NOT DONE
//
// priceOf читает цену товара из map структур.
// Тренирует: чтение поля значения map.
// Сложность: easy
package main_test

import "testing"

type Product struct {
	Price int
	Stock int
}

func priceOf(m map[string]Product, name string) int {
	return m[name].Stock
}

func TestPriceOf(t *testing.T) {
	m := map[string]Product{"pen": {Price: 30, Stock: 7}}
	if priceOf(m, "pen") != 30 || priceOf(m, "cup") != 0 {
		t.Errorf("priceOf works incorrectly")
	}
}
