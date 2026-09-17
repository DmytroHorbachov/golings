// maps42
// Make the tests pass!

// I AM NOT DONE
//
// Кэш цен индексирован строкой, склеенной из категории и артикула.
// Разные пары дают одинаковый ключ и перетирают друг друга.
// Тренирует: ключи из конкатенации строк неоднозначны.
// Сложность: hard
package main_test

import "testing"

type PriceCache map[string]int

func (c PriceCache) Set(cat, sku string, p int) { c[cat+sku] = p }
func (c PriceCache) Get(cat, sku string) int    { return c[cat+sku] }

func TestPriceCache(t *testing.T) {
	c := PriceCache{}
	c.Set("tv", "100", 500)
	c.Set("tv1", "00", 900)
	if c.Get("tv", "100") != 500 || c.Get("tv1", "00") != 900 {
		t.Errorf("prices collided: %v", c)
	}
}
