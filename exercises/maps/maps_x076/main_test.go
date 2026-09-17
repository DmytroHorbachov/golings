// maps_x076: Копия структуры в range
// Make the tests pass!
// I AM NOT DONE
//
// applyDiscount снижает цены всех товаров в map на 10%.
// Цены не меняются.
// Тренирует: значение в range по map — копия; её нужно записать обратно.
// Сложность: hard
package main_test

import "testing"

type Product struct {
	Name  string
	Price int
}

func applyDiscount(m map[string]Product) {
	for k, p := range m {
		p.Price = p.Price * 9 / 10
		_ = k
	}
}

func TestApplyDiscount(t *testing.T) {
	m := map[string]Product{"a": {"a", 100}, "b": {"b", 50}}
	applyDiscount(m)
	if m["a"].Price != 90 || m["b"].Price != 45 {
		t.Errorf("prices = %v", m)
	}
}
