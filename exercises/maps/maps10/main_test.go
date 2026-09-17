// maps10
// Make the tests pass!

// I AM NOT DONE
//
// Цены записываются в map по ключу, собранному через Sprintf. При чтении
// ключ форматируется иначе, и цена не находится.
// Тренирует: ключи должны строиться одной функцией.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func priceKey(item string, size float64) string {
	return fmt.Sprintf("%s/%.1f", item, size)
}

type Prices map[string]int

func (p Prices) Set(item string, size float64, price int) {
	p[priceKey(item, size)] = price
}

func (p Prices) Get(item string, size float64) int {
	return p[fmt.Sprintf("%s/%v", item, size)]
}

func TestPrices(t *testing.T) {
	p := Prices{}
	p.Set("milk", 1, 90)
	p.Set("milk", 0.5, 50)
	if p.Get("milk", 1) != 90 || p.Get("milk", 0.5) != 50 {
		t.Errorf("prices = %v", p)
	}
}
