// maps_x083: Кэш в структуре по значению
// Make the tests pass!
// I AM NOT DONE
//
// Cache.Put сохраняет значение, но метод объявлен со значимым получателем,
// а map создаётся внутри метода. Данные теряются.
// Тренирует: присваивание полю-map в методе со значимым получателем не видно снаружи.
// Сложность: hard
package main_test

import "testing"

type Cache struct {
	items map[string]string
}

func (c Cache) Put(k, v string) {
	if c.items == nil {
		c.items = map[string]string{}
	}
	c.items[k] = v
}

func (c Cache) Get(k string) string { return c.items[k] }

func TestCache(t *testing.T) {
	var c Cache
	c.Put("a", "1")
	c.Put("b", "2")
	if c.Get("a") != "1" || c.Get("b") != "2" {
		t.Errorf("cache lost values: %v", c.items)
	}
}
