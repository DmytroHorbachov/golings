// structs75
// Make the tests pass!

// I AM NOT DONE
//
// Копия структуры независима от оригинала.
// withDiscount возвращает товар со скидкой, не меняя исходный.
// Тренирует: присваивание структуры копирует её.
// Сложность: easy
package main_test

import "testing"

type Item struct {
	Name  string
	Price int
}

func withDiscount(it *Item) Item {
	c := *it
	it.Price /= 2
	return c
}

func TestWithDiscount(t *testing.T) {
	it := Item{"tea", 100}
	d := withDiscount(&it)
	if d.Price != 50 || it.Price != 100 {
		t.Errorf("discounted = %d, original = %d", d.Price, it.Price)
	}
}
