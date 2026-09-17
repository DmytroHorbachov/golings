// anonymous_functions_x016: Сортировка по цене
// Make the tests pass!
// I AM NOT DONE
//
// byPrice сортирует товары по цене, сохраняя порядок одинаковых.
// Тренирует: литерал в sort.SliceStable.
// Сложность: easy
package main_test

import (
	"sort"
	"testing"
)

type Item struct {
	Name  string
	Price int
}

func byPrice(items []Item) {
	sort.SliceStable(items, func(i, j int) bool {
		return items[i].Price > items[j].Price
	})
}

func TestByPrice(t *testing.T) {
	items := []Item{{"b", 5}, {"a", 1}, {"c", 5}}
	byPrice(items)
	if items[0].Name != "a" || items[1].Name != "b" || items[2].Name != "c" {
		t.Errorf("byPrice = %v", items)
	}
}
